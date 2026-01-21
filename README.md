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

### Build

```bash
go mod tidy

make build-mac
or
make build-linux
```

### One-shot Mode (Command-line)

Ask a question directly from the command line:

```bash
$ ./aicli "what is the capital of France?"
⏳ Waiting for ChatGPT...
The capital of France is Paris.
```

**Important:** For queries with special characters or apostrophes, use quotes:
```bash
./aicli "what's the capital of japan?"
./aicli "how does DNS work?"
```

### Interactive Mode

Run the application without arguments to enter interactive mode:

```bash
$ ./aicli
Welcome to AI CLI - ChatGPT Terminal Interface
Type 'h' for help, 'q' to quit
--------------------------------------------------

[19:49:49] aicli> what is the capital of France?
⏳ Waiting for ChatGPT...

The capital of France is Paris.

[19:49:50] aicli> q
👋 Exiting. Goodbye!
```
