# AI CLI - ChatGPT terminal interface

![](https://img.shields.io/github/issues/michalswi/aicli)
![](https://img.shields.io/github/forks/michalswi/aicli)
![](https://img.shields.io/github/stars/michalswi/aicli)
![](https://img.shields.io/github/last-commit/michalswi/aicli)

go-based terminal application for interacting with OpenAI's ChatGPT API

## Features

- 💬 Interactive chat with ChatGPT in your terminal
- 🚀 One-shot mode: ask questions directly from command line in your terminal

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

### Help

```bash
aicli --help
```

### One-shot Mode (command-line)

Ask a question directly from the command line:

```bash
$ aicli what day is today
⏳ Waiting for ChatGPT...
--------------------------------------------------
Today is Tuesday, February 10, 2026.
--------------------------------------------------
⏱️  Response time: 12.87s | 🪙 Tokens: 97 in, 533 out, 630 total

$ aicli "what's the capital of Poland? short"
⏳ Waiting for ChatGPT...
--------------------------------------------------
Warsaw.
--------------------------------------------------
⏱️  Response time: 2.91s | 🪙 Tokens: 101 in, 76 out, 177 total
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
Welcome to AI CLI - chatGPT terminal interface
Type 'h' for help, 'q' to quit
--------------------------------------------------

[09:01:48] aicli> what's the capital of Poland? short
⏳ Waiting for ChatGPT...

Warsaw.
--------------------------------------------------
⏱️  Response time: 1.90s | 🪙 Tokens: 101 in, 12 out, 113 total

[09:02:10] aicli> q
👋 Exiting. Goodbye!
```
