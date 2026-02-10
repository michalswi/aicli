package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Handle special flags that don't require API key
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "-h", "--help":
			displayHelp()
			return
		case "-v", "--version":
			fmt.Printf("%s version %s\n", appName, version)
			return
		case "-sp", "--sys-prompt":
			fmt.Println("System Prompt:")
			fmt.Println(strings.Repeat("-", 50))
			fmt.Println(systemPrompt)
			return
		case "-m", "--model":
			fmt.Println(openaiModel)
			return
		}
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the OPENAI_API_KEY environment variable")
	}

	client := openai.NewClient(apiKey)
	if len(os.Args) > 1 {
		// Join all arguments as the query
		query := strings.Join(os.Args[1:], " ")
		handleOneShot(client, query)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to AI CLI - ChatGPT terminal interface")
	fmt.Println("Type 'h' for help, 'q' to quit")
	fmt.Println(strings.Repeat("-", 50))

	conversation := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}

	for {
		command, shouldContinue := prompt(reader)
		if !shouldContinue {
			break
		}

		if command == "h" {
			displayHelp()
			continue
		}

		conversation = append(conversation, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: command,
		})

		response, duration, err := chatWithGPT(client, conversation)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("\n%s\n", response.Choices[0].Message.Content)
		fmt.Printf("%s\n", strings.Repeat("-", 50))
		fmt.Printf("⏱️  Response time: %.2fs | ", duration.Seconds())
		fmt.Printf("🪙 Tokens: %d in, %d out, %d total\n",
			response.Usage.PromptTokens,
			response.Usage.CompletionTokens,
			response.Usage.TotalTokens)

		conversation = append(conversation, response.Choices[0].Message)
	}
}

// handleOneShot processes a single query and exits
func handleOneShot(client *openai.Client, query string) {
	conversation := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: query,
		},
	}

	response, duration, err := chatWithGPT(client, conversation)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Println(response.Choices[0].Message.Content)
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Printf("⏱️  Response time: %.2fs | ", duration.Seconds())
	fmt.Printf("🪙 Tokens: %d in, %d out, %d total\n",
		response.Usage.PromptTokens,
		response.Usage.CompletionTokens,
		response.Usage.TotalTokens)
}

// chatWithGPT sends the conversation to OpenAI and returns the response with duration
func chatWithGPT(client *openai.Client, conversation []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, time.Duration, error) {
	fmt.Println("⏳ Waiting for ChatGPT...")

	startTime := time.Now()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openaiModel,
			Messages: conversation,
		},
	)
	duration := time.Since(startTime)

	if err != nil {
		return openai.ChatCompletionResponse{}, 0, fmt.Errorf("ChatGPT API error: %w", err)
	}

	return resp, duration, nil
}

// prompt displays a prompt to the user and reads their input
func prompt(reader *bufio.Reader) (string, bool) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("\n[%s] %s> ", timestamp, appName)

	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", true
	}

	command = strings.TrimSpace(command)

	switch command {
	case "q", "quit", "exit":
		fmt.Println("👋 Goodbye!")
		return "", false
	case "h", "help":
		return "h", true
	case "":
		return "", true
	}

	return command, true
}

// displayHelp prints available commands
func displayHelp() {
	fmt.Println("\n📖 Available Commands:")
	fmt.Println("\nInteractive mode commands:")
	fmt.Println("  h, help  - Display this help message")
	fmt.Println("  q, quit  - Exit the application")
	fmt.Println("  Any other text will be sent to ChatGPT")
	fmt.Println("\nCommand-line options:")
	fmt.Println("  -h,  --help       - Display this help message")
	fmt.Println("  -v,  --version    - Display version information")
	fmt.Println("  -sp, --sys-prompt - Display the system prompt")
	fmt.Println("  -m,  --model      - Display the current model being used")
	fmt.Println("  <query>          - Send a one-shot query to ChatGPT")
	fmt.Println()
}
