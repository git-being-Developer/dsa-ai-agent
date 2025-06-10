# DSA AI Agent

**DSA AI Agent** is a powerful command-line tool designed to help you practice and improve your Data Structures and Algorithms skills. Powered by an AI, this tool can generate boilerplate code for LeetCode-style problems, check your solutions for correctness, provide hints, fix your code, and even suggest optimizations.

## Features

* **Problem Initialization**: Automatically generate a boilerplate file for a given problem and language (`dsa-init`).
* **Solution Analysis**: Submit your code for in-depth analysis from an AI agent.
* **Multiple Analysis Modes**:
    * `check`: Verifies if your solution is correct and optimal.
    * `help`: Adds comments to your code to help you find and fix logical errors.
    * `fix`: Provides a corrected version of your code with explanations of the changes.
* **Session Management**: Start and end coding sessions to focus on one problem at a time.
* **Language Support**: Built to support multiple programming languages (currently configured for Go, Python, and C++).

## Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://go.dev/doc/install) (version 1.18 or later)
- [Git](https://git-scm.com/downloads)

## Installation

1.  **Clone the repository:**
    ```sh
    git clone <your-repository-url>
    cd dsa-ai-agent
    ```

2.  **Install dependencies:**
    This project uses Go modules to manage dependencies. Run the following command to download them:
    ```sh
    go mod tidy
    ```

3.  **Build the project:**
    Compile the application into a single executable binary:
    ```sh
    go build -o dsa-agent ./cmd/main.go
    ```

## Configuration

The agent uses the OpenAI API to perform its analysis. You need to configure your API key to use the tool.

1.  **Set the Environment Variable:**
    Set the variable in your terminal session before running the agent.
    In the conf.env file
    set OPEN_AI_KEY="your-secret-api-key"
    ```

## Usage

You can run the agent using the compiled binary from the root directory.

```sh
./dsa-agent
```

### Commands

The agent provides a simple, interactive command-line interface.

-   **Start a new session:**
    `dsa-init <problem-slug> <language>`
    This command fetches boilerplate code for a problem and creates a new file in the `problems/` directory (e.g., `problems/two-sum.go`).

    *Example:*
    ```
    > dsa-init combination-sum go
    ```

-   **Submit your solution for analysis:**
    `dsa-submit <mode>`
    Once you have edited the problem file with your solution, use this command to have the AI analyze it.

    *Available Modes:*
    -   `check`: Validates if your solution is correct.
    -   `help`: Provides hints by adding comments to your code.
    -   `fix`: Corrects your code and explains the changes.

    *Example:*
    ```
    > dsa-submit fix
    ```

-   **End the current session:**
    `dsa-exit`
    This command terminates the active session, allowing you to start a new one.

-   **Get help:**
    `dsa-help`
    Displays the list of available commands and their usage.

## Project Structure

```
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
