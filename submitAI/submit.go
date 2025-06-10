package submitAI

import (
	"context"
	"dsa-ai-agent/config"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
)

type CodeSubmit interface {
	ProcessSolution(mode, title, lang string) error
}

type OpenAISubmit struct {
	client        openai.Client
	userCode      string
	aiResponse    string
	prompt        string
	questionDesc  string
	questionTitle string
	questionLang  string
}

func NewSubmitAgent() *OpenAISubmit {
	return &OpenAISubmit{}
}

func (c *OpenAISubmit) connectToAgent() error {
	client := openai.NewClient(
		option.WithAPIKey(config.AppConfig.OPEN_AI_API),
	)
	c.client = client
	fmt.Println("✅ AI Agent connected.")
	return nil
}

// readCode reads the user's solution from the file created by the scrapper.
func (c *OpenAISubmit) readCode(title, lang string) error {
	fileName := fmt.Sprintf("%s.%s", title, lang)
	filePath := filepath.Join("problems", fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read solution file '%s': %w", filePath, err)
	}
	c.userCode = string(content)
	fmt.Printf("✅ Successfully read code from %s\n", filePath)
	return nil
}

// initializePrompt creates the correct prompt based on the processing mode.
func (c *OpenAISubmit) initializePrompt(mode string) {
	var promptTemplate string

	switch mode {
	case "check":
		promptTemplate = `You are a LeetCode problem validator. Your only job is to determine if the provided user's code is a correct and optimal solution for the given problem. Do not provide explanations or code improvements. Analyze the user's code based on the problem description and common edge cases. Respond with only "OK" if the solution is correct, or "Not OK" if it is incorrect.

---
**Problem Description:**
%s
---
**User's Code (%s):**
%s
---
**Your Response:**`

	case "help":
		promptTemplate = `You are a helpful coding tutor. Your task is to analyze a user's incorrect LeetCode solution and provide hints and comments to guide them. Add comments directly into the code to explain the logical errors and suggest where to look for the fix. Do not write the final correct code. Return the user's original code, but with your helpful comments added.

---
**Problem Description:**
%s
---
**User's Code (%s):**
%s
---
**Your Commented Code:**`

	case "fix":
		promptTemplate = `You are an expert programmer. Your task is to fix the provided incorrect LeetCode solution. Your response should be the complete, corrected code that solves the problem efficiently. Add comments as a code explaining what you changed and why.no extra explanation

---
**Problem Description:**
%s
---
**User's Incorrect Code (%s):**
%s
---
**Your Corrected Code:**`

	case "optimize":
		promptTemplate = `You are a principal engineer specializing in performance. The user has provided a correct but potentially suboptimal solution to a LeetCode problem. Your task is to rewrite the code to be more optimal in terms of time or space complexity. Your response must be only the optimized code, with comments explaining the performance improvements.

---
**Problem Description:**
%s
---
**User's Correct Code (%s):**
%s
---
**Your Optimized Code:**`
	}

	c.prompt = fmt.Sprintf(promptTemplate, c.questionDesc, c.questionLang, c.userCode)
}

// executeAIPrompt sends the generated prompt to the API and stores the response.
func (c *OpenAISubmit) executeAIPrompt() error {
	fmt.Println("🤖 Sending request to AI for analysis... (this may take a moment)")
	resp, err := c.client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(c.prompt),
		},
		Model:       openai.ChatModelGPT4,
		Temperature: param.NewOpt(0.2),
	})
	if err != nil {
		return fmt.Errorf("chat completion failed: %w", err)
	}

	if len(resp.Choices) == 0 {
		return errors.New("no choices returned from OpenAI")
	}

	c.aiResponse = resp.Choices[0].Message.Content
	fmt.Println("✅ AI analysis complete.")
	return nil
}

// saveCodeToFile saves the AI's response to a new file, e.g., "two-sum-fixed.py"
func (c *OpenAISubmit) saveCodeToFile(mode string) error {
	var extension string
	switch c.questionLang {
	case "python":
		extension = ".py"
	case "go":
		extension = ".go"
	default:
		return fmt.Errorf("unsupported language: %s", c.questionLang)
	}

	dir := "problems"
	// Create a new filename based on the mode, e.g., two-sum-fixed.py
	fileName := fmt.Sprintf("%s-%s%s", c.questionTitle, mode, extension)
	filePath := filepath.Join(dir, fileName)
	code := strings.SplitN(c.aiResponse, "\n", 2)
	c.aiResponse = code[1]
	c.aiResponse = strings.Trim(c.aiResponse, "`")
	err := os.WriteFile(filePath, []byte(c.aiResponse), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Printf("✅ AI response saved successfully to %s\n", filePath)
	return nil
}

// ProcessSolution is the main public method that orchestrates the entire workflow.
func (c *OpenAISubmit) ProcessSolution(mode, title, lang, problemDesc string) error {
	c.questionTitle = title
	c.questionLang = lang
	c.questionDesc = problemDesc
	var err error
	if err = c.connectToAgent(); err != nil {
		return err
	}
	if err := c.readCode(title, getExtention(lang)); err != nil {
		return err
	}
	c.initializePrompt(mode)
	err = c.executeAIPrompt()

	if err != nil {
		return err
	}

	if mode == "check" {
		fmt.Printf("\n--- AI Verification Result ---\n%s\n--------------------------\n", c.aiResponse)

	} else {
		if err := c.saveCodeToFile(mode); err != nil {
			return err
		}
	}

	return nil
}

//util function

func getExtention(lang string) string {
	switch lang {
	case "python":
		return "py"
	case "go":
		return "go"
	case "c++":
		return "cpp"
	}
	return ""
}
