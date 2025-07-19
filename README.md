# Daily Summary

A command-line tool that inspects your recent file activity and uses the OpenAI API to generate a concise, insightful summary of what you’ve been working on over a specified time window.

## Features

- Scans files changed within the last _N_ hours (default: 12) using [`fd`](https://github.com/sharkdp/fd).
- Gathers file metadata via macOS’s `mdls` command.
- Builds a structured prompt and sends it to OpenAI’s Chat Completions API.
- Prints a human-readable summary highlighting projects, tools, and notable activities.

## Prerequisites

- Go 1.18 or later  
- OpenAI API key (set via `OPENAI_API_KEY`)  
- `fd` (a simple, fast and user-friendly alternative to `find`)  
- **macOS only:** `mdls` (built into macOS; used to read file metadata)

### Install Prerequisites

macOS (Homebrew)
```
brew install go fd
# mdls is pre-installed on macOS
```

Ubuntu / Debian
```
sudo apt update
sudo apt install -y golang fd-find
# Note: macOS’s `mdls` is not available on Linux; file metadata gathering is limited.
```

Windows (PowerShell + Scoop or Chocolatey)
```
# With Chocolatey
choco install golang fd
# mdls is not available; metadata gathering is limited.
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
  OpenAI model to use (default: `gpt-4o`).

### Example

Scan changes over the last 6 hours with GPT-4:
```
export OPENAI_API_KEY="sk-..."
go run dailysummary.go --since-hours 6 --model gpt-4
```

## Environment Variables

- `OPENAI_API_KEY`  
  Your OpenAI API key. Must be set before running.

## Contributing

Contributions are welcome! Please open issues or submit pull requests to improve functionality, add cross-platform metadata support, or refine the prompt templates.

## License

This project is licensed under the MIT License.  
See `LICENSE` for details.
