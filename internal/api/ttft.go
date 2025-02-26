package api

import (
	"context"
	"log"
	"math"
	"sync"
	"time"

	"github.com/sashabaranov/go-openai"
)

// MeasureTTFT calculates the maximum and minimum Time to First Token (TTFT) for API responses.
func MeasureTTFT(client *openai.Client, model, prompt string, concurrency int) (float64, float64) {
	var wg sync.WaitGroup
	ttftChan := make(chan float64, concurrency)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			stream, err := client.CreateChatCompletionStream(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: model,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleSystem,
							Content: "You are a helpful assistant.",
						},
						{
							Role:    openai.ChatMessageRoleUser,
							Content: prompt,
						},
					},
					MaxTokens:   512,
					Temperature: 1,
					Stream:      true,
				},
			)
			if err != nil {
				log.Printf("TTFT Request error: %v", err)
				return
			}
			defer stream.Close()

			// Listen for the first response
			_, err = stream.Recv()
			if err != nil {
				log.Printf("TTFT Stream error: %v", err)
				return
			}

			// Record TTFT
			ttft := time.Since(start).Seconds()
			ttftChan <- ttft
		}(i)
	}

	wg.Wait()
	close(ttftChan)

	// Calculate maximum and minimum TTFT
	maxTTFT := 0.0
	minTTFT := math.Inf(1)
	for ttft := range ttftChan {
		if ttft > maxTTFT {
			maxTTFT = ttft
		}
		if ttft < minTTFT {
			minTTFT = ttft
		}
	}

	return maxTTFT, minTTFT
}

func MeasureTTFTwithRandomInput(client *openai.Client, model string, numWords int, concurrency int) (float64, float64) {
	var wg sync.WaitGroup
	ttftChan := make(chan float64, concurrency)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			prompt := generateRandomPhrase(numWords)
			start := time.Now()
			stream, err := client.CreateChatCompletionStream(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: model,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleSystem,
							Content: "You are a helpful assistant.",
						},
						{
							Role:    openai.ChatMessageRoleUser,
							Content: prompt,
						},
					},
					MaxTokens:   512,
					Temperature: 1,
					Stream:      true,
				},
			)
			if err != nil {
				log.Printf("TTFT Request error: %v", err)
				return
			}
			defer stream.Close()

			// Listen for the first response
			_, err = stream.Recv()
			if err != nil {
				log.Printf("TTFT Stream error: %v", err)
				return
			}

			// Record TTFT
			ttft := time.Since(start).Seconds()
			ttftChan <- ttft
		}(i)
	}

	wg.Wait()
	close(ttftChan)

	// Calculate maximum and minimum TTFT
	maxTTFT := 0.0
	minTTFT := math.Inf(1)
	for ttft := range ttftChan {
		if ttft > maxTTFT {
			maxTTFT = ttft
		}
		if ttft < minTTFT {
			minTTFT = ttft
		}
	}

	return maxTTFT, minTTFT
}
