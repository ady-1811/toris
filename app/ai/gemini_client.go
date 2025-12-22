package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/toris/utils"
	"google.golang.org/api/option"
)

type GeminiCommandClient struct {
	Client *genai.Client
	Model  *genai.GenerativeModel
	OSName string
}

func NewGeminiCommandClient(modelID string) (*GeminiCommandClient, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("No .env file found or error loading it: %v", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing GEMINI_API_KEY environment variable")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	c := &GeminiCommandClient{
		Client: client,
	}
	c.OSName = utils.GetOS()
	c.Model = initializeModel(c, modelID)
	return c, nil
}

func (c *GeminiCommandClient) GetCommand(ctx context.Context, userInput string) (*CommandResponse, error) {
	return c.getResponse(ctx, c.Model, userInput)
}

func (c *GeminiCommandClient) ScanForErrors(ctx context.Context) (*CommandResponse, error) {
	logs, err := utils.GetConsoleOutput() //placeholder until I figure out how to extract logs
	if err != nil {
		return nil, fmt.Errorf("Your OS doesn't support this command: %v", err)
	}
	fmt.Print(logs)
	error_instruction := "Give me instructions to fix the following error: "
	return c.getResponse(ctx, c.Model, error_instruction+logs)
}

func (c *GeminiCommandClient) getResponse(ctx context.Context, model *genai.GenerativeModel, userInput string) (*CommandResponse, error) {
	resp, err := model.GenerateContent(ctx, genai.Text(userInput))
	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("no candidates returned from Gemini")
	}

	var cmdResp CommandResponse
	part := resp.Candidates[0].Content.Parts[0]
	if text, ok := part.(genai.Text); ok {
		err := json.Unmarshal([]byte(text), &cmdResp)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %v", err)
		}
		return &cmdResp, nil
	}

	return nil, fmt.Errorf("unexpected response format")
}

func initializeModel(c *GeminiCommandClient, modelID string) *genai.GenerativeModel {
	model := c.Client.GenerativeModel(modelID)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(fmt.Sprintf(`
				You are a command-line assistant. Your task is to translate natural language
				into a JSON command object for %s, and give instructions on how to run it. 
				If the user asks you to fix an error, provide solution in the same format, but you may 
				leave the "command" empty only when not required.

				Always return ONLY valid JSON in this format:
				{
					"command": string,
					"confidence": float
					"instructions": [string]
				}`, c.OSName)),
		},
	}

	model.ResponseMIMEType = "application/json"
	temp := float32(0.1)
	model.Temperature = &temp

	return model
}
