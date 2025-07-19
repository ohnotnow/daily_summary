# Daily Summary

A command-line tool that inspects your recent file activity and uses the OpenAI API to generate a concise, insightful summary of what you’ve been working on over a specified time window.

## Features

- Scans files changed within the last _N_ hours (default: 12) using [`fd`](https://github.com/sharkdp/fd).
- Gathers file metadata via macOS’s `mdls` command.
- Builds a structured prompt and sends it to OpenAI’s Chat Completions API.
- Prints a human-readable summary highlighting projects, tools, and notable activities.

## Binaries

Pre-built binaries for macOS, Linux, and Windows are available on the [Releases](https://github.com/ohnotnow/daily_summary/releases) page.

## Prerequisites

- Go 1.18 or later (if self-compiling) 
- OpenAI API key (set via `OPENAI_API_KEY`)  
- `fd` (a simple, fast and user-friendly alternative to `find`)  
- **macOS only:** `mdls` (built into macOS; used to read file metadata)

### Install Prerequisites

macOS (Homebrew)
```
brew install go fd
# mdls is pre-installed on macOS
```

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/ohnotnow/daily_summary.git
   cd daily_summary
   ```

2. (Optional) Build the binary:
   ```
   go build -o daily_summary dailysummary.go
   ```

## Usage

Set your OpenAI API key in the environment:
```
export OPENAI_API_KEY="sk-..."
# Windows PowerShell:
#  $Env:OPENAI_API_KEY = "sk-..."
```

### Run without building
```
go run dailysummary.go [flags]
```

### Run the built binary
```
./daily_summary [flags]
```

### Flags

- `--since-hours`  
  Number of hours back to include file changes (default: `12`).

- `--model`  
  OpenAI model to use (default: `4o-mini`).

### Example

Scan changes over the last 6 hours with GPT-4:
```
export OPENAI_API_KEY="sk-..."
go run dailysummary.go --since-hours 6 --model gpt-4
```

If you have the [glow](https://github.com/charmbracelet/glow) tool available then you can get markdown rendering in your terminal too :

```
go run dailysummary.go --since-hours 6 | glow
```

## Environment Variables

- `OPENAI_API_KEY`  
  Your OpenAI API key. Must be set before running.

## Contributing

Contributions are welcome! Please open issues or submit pull requests to improve functionality, add cross-platform metadata support, or refine the prompt templates.

## License

This project is licensed under the MIT License.  
See `LICENSE` for details.

## Example Output

# 1. Overall Summary
Over the past several hours, you focused on building and refining a suite of command-line and automation tools in both Go and Python, while also updating your personal notes and capturing visual summaries of a chat or channel.

# 2. Project or Theme Breakdown
- **Code Review Tool (Python)**
  - Created `main.py` and `system_prompt.md` under `code-reviewer/`
  - Drafted a `README.md` to document usage
- **Prompt Evaluator (Python)**
  - Initialized a Python project (`pyproject.toml`, `uv.lock`)
  - Authored `main.py`, README and LICENSE files
- **Gepetto CLI (Go)**
  - Wrote `gepetto.go` in `gepetto_cli_go/`
  - Updated the `gepetto` executable in your local `bin/` directory
- **Daily Summary Generator (Go)**
  - Added `dailysummary.go` under `daily_summary/`
  - Created `README.md` and LICENSE for project scaffolding
- **Shell Utility**
  - Added `ff.sh`, presumably a quick shell script for media or data processing
- **Disk & Channel Summaries**
  - Generated `disk_activity_summary.json`
  - Captured a screenshot (`.png`) and exported a channel summary (`.webp`)
- **Personal Notes (Obsidian Vault)**
  - Updated `Shopping.md`
  - Edited `Work/AI Working Group.md`

# 3. Technology and Tools Used
- **Languages**: Go, Python, Shell scripting
- **Documentation**: Markdown for READMEs, LICENSEs and personal notes
- **Packaging**: Python packaging via `pyproject.toml` and lockfile
- **CLI Tools**: Custom executables and scripts (e.g., `gepetto`, `ff.sh`)
- **Note-taking**: Obsidian Vault Markdown files
- **Visual Capture**: Screenshot and WebP export

# 4. Notable Activities
- Laid the groundwork for two new Go-based tools: one for daily summaries, another as a generic CLI (`gepetto`).
- Kick-started two Python utilities: a code-review assistant and a prompt evaluation tool, complete with project metadata and documentation.
- Wrote or updated multiple README and LICENSE files to formalize new projects.
- Ran a disk-activity summary job and captured a chat/channel summary via image exports.
- Maintained personal productivity notes, including a shopping list and AI Working Group meeting notes.

