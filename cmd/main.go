package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"dsa-ai-agent/config"
	"dsa-ai-agent/scrapperAI"
	"dsa-ai-agent/submitAI"
)

var (
	scrapperAgent *scrapperAI.OpenAIAgent
	sessionActive bool
)

func main() {

	config.NewAppConfig()
	config.AppConfig.LoadConfig()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("✅ DSA AI AGENT is Ready.")
	printHelp()

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(input, "dsa") {
			if input != "" {
				fmt.Println("Invalid command. Please start with 'dsa'.")
			}
			fmt.Print("> ")
			continue
		}

		args := strings.Fields(input)
		cmd := args[0]

		switch cmd {
		case "dsa-init":
			handleInit(args)
		case "dsa-submit":
			handleSubmit(args)
		case "dsa-exit":
			handleExit()
		case "dsa-help":
			printHelp()
		default:
			fmt.Println("Unknown command. Type 'dsa-help' for a list of commands.")
		}
		fmt.Print("> ")
	}
}

func handleInit(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: dsa-init <problem-slug> <language>")
		fmt.Println("Example: dsa-init two-sum python")
		fmt.Println("To submit a custom code, use dsa-submit --force <mode> <title> <language>. Note: title should match")
		return
	}
	if sessionActive {
		fmt.Println("A session is already active. Please run 'dsa-exit' first.")
		return
	}

	slug := args[1]
	lang := args[2]

	fmt.Println("--- Starting new session ---")
	scrapperAgent = scrapperAI.NewAIAgent(lang, slug, "")
	err := scrapperAgent.Run()
	if err != nil {
		log.Printf("❌ Error starting session: %v", err)
		scrapperAgent = nil
		return
	}

	sessionActive = true
	fmt.Printf("\n✅ Session started for '%s'. You can now edit the file in the 'problems' directory.\n", scrapperAgent.GetQuestionTitle())
	fmt.Println("When you are ready, use 'dsa-submit <mode>' to analyze your code.")
}

// handleSubmit sends the user's code to the fixer/checker agent.
func handleSubmit(args []string) {

	if !sessionActive {
		fmt.Println("No active session. Please start one with 'dsa-init'.")
		return
	}
	if len(args) != 2 {
		fmt.Println("Usage: dsa-submit <mode>")
		fmt.Println("Available modes: check, help, fix")
		return
	}

	mode := args[1]
	// Validate mode
	validModes := map[string]bool{"check": true, "help": true, "fix": true}
	if !validModes[mode] {
		fmt.Printf("Invalid mode '%s'. Available modes: check, help, fix\n", mode)
		return
	}

	fmt.Println("--- Submitting code for analysis ---")
	fixer := submitAI.NewSubmitAgent()

	// Pass all the necessary details from the active scrapper agent.
	err := fixer.ProcessSolution(
		mode,
		scrapperAgent.QuestionSlug,
		scrapperAgent.Lang,
		scrapperAgent.QuestionTitle,
	)

	if err != nil {
		log.Printf("❌ Error processing solution: %v", err)
	}
}

// handleExit resets the current session.
func handleExit() {
	if !sessionActive {
		fmt.Println("No active session to exit.")
		return
	}
	fmt.Println("--- Ending session ---")
	sessionActive = false
	scrapperAgent = nil
	fmt.Println("✅ Session ended. You can now start a new one.")
}

// printHelp displays the available commands.
func printHelp() {
	fmt.Println("\n--- Available Commands ---")
	fmt.Println("  dsa-init <slug> <lang>   - Start a new session (e.g., dsa-init two-sum python)")
	fmt.Println("  dsa-submit <mode>        - Analyze your code (modes: check, help, fix)")
	fmt.Println("  dsa-exit                 - End the current session")
	fmt.Println("  dsa-help                 - Show this help message")
	fmt.Println("--------------------------")
}
