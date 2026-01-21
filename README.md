# AI CLI - ChatGPT Terminal Interface

![](https://img.shields.io/github/issues/michalswi/aicli)
![](https://img.shields.io/github/forks/michalswi/aicli)
![](https://img.shields.io/github/stars/michalswi/aicli)
![](https://img.shields.io/github/last-commit/michalswi/aicli)

A simple Go-based terminal application for interacting with OpenAI's ChatGPT API.

## Features

- 💬 Interactive chat with ChatGPT in your terminal
- 🚀 One-shot mode: ask questions directly from command line using `:`
- 🔄 Maintains conversation context in interactive mode
- ⚡ Built with `github.com/sashabaranov/go-openai`

## Prerequisites

- Go 1.25 or higher
- OpenAI API key

## Installation

1. Clone or download this repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Setup

Set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY="your-api-key-here"
```

To make this permanent, add it to your `~/.zshrc` or `~/.bashrc`:

```bash
echo 'export OPENAI_API_KEY="your-api-key-here"' >> ~/.zshrc
source ~/.zshrc
```

## Usage

### One-shot Mode (Command-line)

Ask a question directly from the command line using the colon (`:`) separator:

```bash
./aicli : tell me something about kubernetes
./aicli : what is the difference between TCP and UDP?
./aicli : explain aws s3 buckets
```

You can also run without building first:

```bash
go run main.go : how does DNS work?
```

### Interactive Mode

Run the application without arguments to enter interactive mode:

```bash
go run main.go
```

Or build and run:

```bash
make build-mac
or
make build-linux

./aicli
```

### Commands

- Type any message to chat with ChatGPT
- `h` or `help` - Display help message
- `q`, `quit`, or `exit` - Exit the application

## Example Session

### One-shot Mode
```bash
$ ./aicli : what is the capital of France?
⏳ Waiting for ChatGPT...
The capital of France is Paris.

$ ./aicli : explain docker in simple terms
⏳ Waiting for ChatGPT...
Docker is a platform that allows you to package applications and their dependencies into containers...
```

### Interactive Mode
```
Welcome to AI CLI - ChatGPT Terminal Interface
Type 'h' for help, 'q' to quit
--------------------------------------------------

[15:30:45] aicli> What is the capital of France?
⏳ Waiting for ChatGPT...

The capital of France is Paris.

[15:30:52] aicli> Tell me more about it
⏳ Waiting for ChatGPT...

Paris is not only the capital but also the largest city of France...
```

## Configuration

The application uses `GPT-5-mini` by default. To change the model, edit the `openaiModel` constant in `main.go`:

```go
const openaiModel = openai.GPT5Mini
```

## License

See LICENSE file for details.
