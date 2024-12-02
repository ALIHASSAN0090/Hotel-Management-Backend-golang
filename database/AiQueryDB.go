package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AiQueryDB(query string, userID int, c *gin.Context) (string, error) {
	// Implement this function as needed
	return "", nil
}

func GetAiResponceDB(order_details, hotel_details, question string) (string, error) {
	apiKey := os.Getenv("OPENAI_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_KEY environment variable not set")
	}

	requestBody := map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "system", "content": "Behave as a food hotel chat assistant."},
			{"role": "user", "content": fmt.Sprintf("You are a hotel's online chat assistent you have hotel data : $1  . and orders data of the current user : $2 . And Question of User is : $3 . Properly give responce as online chat assistant. Only provide responce on basis of provided data.", order_details, hotel_details, question)},
		},
	}

	reqBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if message, ok := choices[0].(map[string]interface{}); ok {
			if content, ok := message["message"].(map[string]interface{})["content"].(string); ok {
				return content, nil
			}
		}
	}

	return "", fmt.Errorf("response not found in API call")
}
