package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type LLMRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type LLMResponse struct {
	Response string `json:"response"`
}

func GenerateBadJazzQuote() string {
	llmApiUrl := os.Getenv("LLM_API_URL")
	if llmApiUrl == "" {
		llmApiUrl = "http://llm:11434"
	}
	prompt := "Generate a short, intentionally bad and mocking quote about jazz or jazz musicians. The quote should not hold jazz in high regard and should sound like a joke. Make sure to always denigrate headless guitars and use the word toan instead of tone"
	llmReq := LLMRequest{
		Model:  "llama3.2",
		Prompt: prompt,
		Stream: false,
	}
	jsonData, err := json.Marshal(llmReq)
	if err != nil {
		log.Printf("[LLM error: %v]", err)
		return fmt.Sprintf("[LLM error: %v]", err)
	}
	resp, err := http.Post(llmApiUrl+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[LLM error: %v]", err)
		return fmt.Sprintf("[LLM error: %v]", err)
	}
	defer resp.Body.Close()
	var llmResp LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResp); err != nil {
		log.Printf("[LLM error: %v]", err)
		return fmt.Sprintf("[LLM error: %v]", err)
	}
	log.Printf("Generated quote: %s", llmResp.Response)
	return strings.TrimSpace(llmResp.Response)
}
