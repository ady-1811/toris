package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"github.com/joho/godotenv"
)

type GeminiCommandClient struct {
	Client  *genai.Client
	ModelID string
	OSName  string
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
		Client:  client,
		ModelID: modelID,
	}
	c.OSName = c.getOS()

	return c, nil
}

func (c *GeminiCommandClient) getOS() string {
	return fmt.Sprintf("%s (%s)", runtime.GOOS, runtime.GOARCH)
}

func (c *GeminiCommandClient) GetCommand(ctx context.Context, userInput string) (*CommandResponse, error) {
	model := c.Client.GenerativeModel(c.ModelID)
	
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(fmt.Sprintf(`
				You are a command-line assistant. Your task is to translate natural language
				into a JSON command object for %s.

				Always return ONLY valid JSON in this format:
				{
					"command": string,
					"confidence": float
				}`, c.OSName)),
		},
	}

	model.ResponseMIMEType = "application/json"
	
	temp := float32(0.1)
	model.Temperature = &temp

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
