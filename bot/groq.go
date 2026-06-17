package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GroqRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type GroqResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}
func (c *ChatBot) GenerateResponse(prompt string) (string, error) {

	c.AddMessage("user", prompt)

	reqBody := GroqRequest{
		Model:    "llama-3.3-70b-versatile",
		Messages: c.History,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.groq.com/openai/v1/chat/completions",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var groqResp GroqResponse

	if err := json.NewDecoder(resp.Body).Decode(&groqResp); err != nil {
		return "", err
	}

	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("no response from model")
	}

	answer := groqResp.Choices[0].Message.Content

	c.AddMessage("assistant", answer)

	return answer, nil
}