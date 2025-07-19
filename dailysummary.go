package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"net/http"
	"io"
)

type FileMetadata struct {
	Path     string `json:"path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Size     int64  `json:"size_bytes"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func runFdSince(sinceHours int) ([]string, error) {
	fdCmd := exec.Command("fd", ".", os.Getenv("HOME"), "--type", "f", "--changed-within", fmt.Sprintf("%dh", sinceHours))
	fdOut, err := fdCmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(fdOut)), "\n")

	// Filter paths
	exclude := []string{".git", "node_modules", "vendor", "venv", "__pycache__", ".tox", ".mypy_cache", ".pytest_cache", ".next", ".cache", "dist", "build", ".idea", ".vscode", ".DS_Store", ".Trash", "Library", "Photos Library.photoslibrary", "Containers", "Mail", "Safari", "com.apple", ".bundle", ".framework", ".xcassets"}
	filtered := []string{}
	for _, path := range lines {
		excludeThis := false
		for _, ex := range exclude {
			if strings.Contains(path, "/"+ex+"/") {
				excludeThis = true
				break
			}
		}
		if !excludeThis {
			filtered = append(filtered, path)
		}
	}
	return filtered, nil
}

func getMdls(path string) (FileMetadata, error) {
	md := FileMetadata{Path: path}
	fields := map[string]*string{
		"kMDItemFSName": &md.Name,
		"kMDItemContentType": &md.Type,
		"kMDItemFSCreationDate": &md.Created,
		"kMDItemFSContentChangeDate": &md.Modified,
	}
	for k := range fields {
		out, _ := exec.Command("mdls", "-name", k, "-raw", path).Output()
		*fields[k] = strings.TrimSpace(string(out))
	}
	out, _ := exec.Command("mdls", "-name", "kMDItemLogicalSize", "-raw", path).Output()
	fmt.Sscanf(string(out), "%d", &md.Size)
	return md, nil
}

func buildPrompt(files []FileMetadata) string {
    prompt := `You are an assistant that reviews structured file activity logs and summarises a user's recent computer usage. The data includes files created or modified in the last 12 hours, with metadata such as name, path, type, size, and timestamps.

Your task is to analyse this data and provide a concise, insightful summary of what the user was likely doing during this period. Please include:

1. Overall Summary
2. Project or Theme Breakdown
3. Technology and Tools Used
4. Notable Activities
5. Avoid referring to the user by name, even if it appears in file paths. The user is running this themselves - so if a personal reference is needed then use "you"
6. You should not mention system-level files, app caches, media libraries, or operating system internals such as Apple TV databases, unless they clearly represent user-driven activity.

Here is the file activity data (in JSON):

`

	jsonBytes, _ := json.MarshalIndent(files, "", "  ")
	return prompt + "```\n" + string(jsonBytes) + "\n```"
}

func callOpenAI(apiKey, model, prompt string) (string, error) {
	reqBody := ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{Role: "user", Content: prompt},
		},
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, _ := io.ReadAll(resp.Body)

	var chatResp ChatResponse
	if err := json.Unmarshal(respBytes, &chatResp); err != nil {
		return string(respBytes), err
	}
	return chatResp.Choices[0].Message.Content, nil
}

func main() {
	var since int
	var model string
	flag.IntVar(&since, "since-hours", 12, "Number of hours back to include file changes")
	flag.StringVar(&model, "model", "gpt-4o", "OpenAI model to use")
	flag.Parse()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}

	paths, err := runFdSince(since)
	if err != nil {
		log.Fatal("fd failed:", err)
	}

	var files []FileMetadata
	for _, path := range paths {
		md, err := getMdls(path)
		if err == nil {
			files = append(files, md)
		}
	}

	prompt := buildPrompt(files)
	summary, err := callOpenAI(apiKey, model, prompt)
	if err != nil {
		log.Fatal("OpenAI API error:", err)
	}

	fmt.Println("\n===== Summary =====")
	fmt.Println(summary)
}


