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
export OPENAI_API_KEY="your-api-key-here"
alias aicli='PATH_TO_BINARY/aicli'
```

## Usage

### Build

```bash
$ make

Usage:
  make <target>

Targets:
  build-mac        Build for mac
  build-linux      Build for linux
```

### One-shot Mode (Command-line)

Ask a question directly from the command line:

```bash
$ aicli what day is today
⏳ Waiting for ChatGPT...
--------------------------------------------------
Today is Thursday, January 22, 2026.
--------------------------------------------------
⏱️  Response time: 2.35s | 🪙 Tokens: 83 in, 85 out, 168 total

$ aicli "what is the capital of France?"
⏳ Waiting for ChatGPT...
--------------------------------------------------
Paris.
--------------------------------------------------
⏱️  Response time: 1.97s | 🪙 Tokens: 86 in, 11 out, 97 total
```

**Important:** For queries with special characters or apostrophes, use quotes:
```bash
aicli "what's the capital of poland?"
aicli "how does DNS work?"
```

### Interactive Mode

Run the application without arguments to enter interactive mode:

```bash
$ aicli
Welcome to AI CLI - ChatGPT Terminal Interface
Type 'h' for help, 'q' to quit
--------------------------------------------------

[19:49:49] aicli> what is the capital of France?
⏳ Waiting for ChatGPT...

Paris — the capital and largest city of France, located on the River Seine in northern France.
--------------------------------------------------
⏱️  Response time: 2.60s | 🪙 Tokens: 86 in, 92 out, 178 total

[19:49:50] aicli> q
👋 Exiting. Goodbye!
```
