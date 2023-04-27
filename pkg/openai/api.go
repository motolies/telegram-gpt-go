package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/motolies/telegram-gpt-go/pkg/customLog"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

func Call(apiKey string, prompt string) string {
	// API URL
	apiUrl := "https://api.openai.com/v1/chat/completions"
	//apiUrl := "https://api.openai.com/v1/engines/davinci-codex/completions"

	// body 생성
	reqBody := RequestBody{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.5,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	// HTTP 요청 생성
	requestBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// HTTP 요청 보내기
	client := &http.Client{}
	resp, _ := client.Do(req)

	// HTTP 응답 읽기
	defer resp.Body.Close()
	var responseMap map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseMap)

	if resp.StatusCode != 200 {
		customLog.ColorLog(fmt.Sprintf("OpenAI API Error: %s", responseMap["error"].(map[string]interface{})["message"].(string)), customLog.ERROR)
		return responseMap["error"].(map[string]interface{})["message"].(string)
	} else {
		return responseMap["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	}
}
