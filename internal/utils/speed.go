package utils

import (
	"sync"
	"time"

	"github.com/Yoosu-L/llmapibenchmark/internal/api"

	"github.com/sashabaranov/go-openai"
)

// MeasureSpeed measures API generation throughput and TTFT.
func MeasureSpeed(baseURL, apiKey, model, prompt string, concurrency, maxTokens int, latency float64) (float64, float64, float64, float64) {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseURL
	client := openai.NewClientWithConfig(config)

	var wg sync.WaitGroup
	var responseTokens sync.Map
	var promptTokens sync.Map

	// Measure TTFT
	maxTTFT, minTTFT := api.MeasureTTFT(client, model, prompt, concurrency)

	start := time.Now()

	// Send requests concurrently
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resp, err := api.AskOpenAI(client, model, prompt, maxTokens)
			if err != nil {
				return
			}
			responseTokens.Store(index, resp.Usage.CompletionTokens)
			promptTokens.Store(index, resp.Usage.PromptTokens)
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	// Calculate total tokens
	totalResponseTokens := 0
	responseTokens.Range(func(_, value interface{}) bool {
		totalResponseTokens += value.(int)
		return true
	})

	totalPromptTokens := 0
	promptTokens.Range(func(_, value interface{}) bool {
		totalPromptTokens += value.(int)
		return true
	})

	// Calculate speed (tokens/second)
	generationSpeed := float64(totalResponseTokens) / (duration.Seconds() - latency/1000)

	// Calculate Prompt Throughput
	promptThroughput := float64(totalPromptTokens) / (maxTTFT - latency/1000)

	return generationSpeed, promptThroughput, maxTTFT, minTTFT
}

func MeasureSpeedwithRandomInput(baseURL, apiKey, model string, numWords int, concurrency, maxTokens int, latency float64) (float64, float64, float64, float64) {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseURL
	client := openai.NewClientWithConfig(config)

	var wg sync.WaitGroup
	var responseTokens sync.Map
	var promptTokens sync.Map

	// Measure TTFT
	maxTTFT, minTTFT := api.MeasureTTFTwithRandomInput(client, model, numWords, concurrency)

	start := time.Now()

	// Send requests concurrently
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resp, err := api.AskOpenAIwithRandomInput(client, model, numWords, maxTokens)
			if err != nil {
				return
			}
			responseTokens.Store(index, resp.Usage.CompletionTokens)
			promptTokens.Store(index, resp.Usage.PromptTokens)
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	// Calculate total tokens
	totalResponseTokens := 0
	responseTokens.Range(func(_, value interface{}) bool {
		totalResponseTokens += value.(int)
		return true
	})

	totalPromptTokens := 0
	promptTokens.Range(func(_, value interface{}) bool {
		totalPromptTokens += value.(int)
		return true
	})

	// Calculate speed (tokens/second)
	generationSpeed := float64(totalResponseTokens) / (duration.Seconds() - latency/1000)

	// Calculate Prompt Throughput
	promptThroughput := float64(totalPromptTokens) / (maxTTFT - latency/1000)

	return generationSpeed, promptThroughput, maxTTFT, minTTFT
}
