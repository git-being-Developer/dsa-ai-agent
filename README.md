DSA AI Agent
DSA AI Agent is a powerful command-line tool designed to help you practice and improve your Data Structures and Algorithms skills. Powered by an AI, this tool can generate boilerplate code for LeetCode-style problems, check your solutions for correctness, provide hints, fix your code, and even suggest optimizations.

Features
Problem Initialization: Automatically generate a boilerplate file for a given problem and language (dsa-init).

Solution Analysis: Submit your code for in-depth analysis from an AI agent.

Multiple Analysis Modes:

check: Verifies if your solution is correct and optimal.

help: Adds comments to your code to help you find and fix logical errors.

fix: Provides a corrected version of your code with explanations of the changes.

Session Management: Start and end coding sessions to focus on one problem at a time.

Language Support: Built to support multiple programming languages (currently configured for Go, Python, and C++).

Prerequisites
Before you begin, ensure you have the following installed:

Go (version 1.18 or later)

Git

Installation
Clone the repository:

git clone <your-repository-url>
cd dsa-ai-agent

Install dependencies:
This project uses Go modules to manage dependencies. Run the following command to download them:

go mod tidy

Build the project:
Compile the application into a single executable binary:

go build -o dsa-agent ./cmd/main.go

Configuration
The agent uses the OpenAI API to perform its analysis. You need to configure your API key to use the tool.

⚠️ Important Security Warning!

Your current code has a hardcoded API key in scrapperAI/scrapperAi.go and submitAI/submit.go. You must remove this key and use an environment variable to keep it secure before pushing your code to a public repository.

Remove the hardcoded key:
Delete the const OPEN_AI_API = "sk-..." line from both scrapperAI/scrapperAi.go and submitAI/submit.go.

Modify the code to use an environment variable:
Update the connectToAgent function in both files to read the key from the environment.

// In scrapperAI/scrapperAi.go and submitAI/submit.go

import (
    "os"
    "errors"
    // ... other imports
)

func (o *OpenAIAgent) connectToAgent() error { // or OpenAISubmit for the other file
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        return errors.New("the OPENAI_API_KEY environment variable is not set")
    }
    client := openai.NewClient(
        option.WithAPIKey(apiKey),
    )
    o.client = client
    fmt.Println("✅ AI Agent connected.")
    return nil
}

Set the Environment Variable:
Set the variable in your terminal session before running the agent.

On macOS/Linux:

export OPENAI_API_KEY="your-secret-api-key"

On Windows (Command Prompt):

set OPENAI_API_KEY="your-secret-api-key"

Usage
You can run the agent using the compiled binary from the root directory.

./dsa-agent

Commands
The agent provides a simple, interactive command-line interface.

Start a new session:
dsa-init <problem-slug> <language>
This command fetches boilerplate code for a problem and creates a new file in the problems/ directory (e.g., problems/two-sum.go).

Example:

> dsa-init combination-sum go

Submit your solution for analysis:
dsa-submit <mode>
Once you have edited the problem file with your solution, use this command to have the AI analyze it.

Available Modes:

check: Validates if your solution is correct.

help: Provides hints by adding comments to your code.

fix: Corrects your code and explains the changes.

Example:

> dsa-submit fix

End the current session:
dsa-exit
This command terminates the active session, allowing you to start a new one.

Get help:
dsa-help
Displays the list of available commands and their usage.

Project Structure
dsa-ai-agent/
├── cmd/
│   └── main.go         # Main application entry point and CLI handler
├── scrapperAI/
│   └── scrapperAi.go   # AI agent for generating boilerplate code
├── submitAI/
│   └── submit.go       # AI agent for analyzing user solutions
├── problems/
│   └── ...             # Problem files are generated and edited here
├── go.mod
├── go.sum
└── README.md
