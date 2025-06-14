package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// PrintBenchmarkHeader prints the benchmark header with details about the test.
func PrintBenchmarkHeader(modelName string, inputTokens int, maxTokens int, latency float64) {
	banner :=
		`
################################################################################################################
				          LLM API Throughput Benchmark
				    https://github.com/yuhaowin/llmapibenchmark
					 Time：%s
################################################################################################################`

	fmt.Printf(banner+"\n", time.Now().UTC().Format("2006-01-02 15:04:05 UTC+0"))
	fmt.Printf("Input Tokens: %d\n", inputTokens)
	fmt.Printf("Output Tokens: %d\n", maxTokens)
	fmt.Printf("Test Model: %s\n", modelName)
	fmt.Printf("Latency: %.2f ms\n\n", latency)
}

// SaveResultsToMD saves the benchmark results to a Markdown file.
func SaveResultsToMD(results [][]interface{}, modelName string, inputTokens int, maxTokens int, latency float64) {
	filename := fmt.Sprintf("API_Throughput_%s.md", modelName)
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("```\nInput Tokens: %d\n", inputTokens))
	file.WriteString(fmt.Sprintf("Output Tokens: %d\n", maxTokens))
	file.WriteString(fmt.Sprintf("Test Model: %s\n", modelName))
	file.WriteString(fmt.Sprintf("Latency: %.2f ms\n```\n\n", latency))
	file.WriteString("| Concurrency | Generation Throughput (tokens/s) |  Prompt Throughput (tokens/s) | Min TTFT (s) | Max TTFT (s) |\n")
	file.WriteString("|-------------|----------------------------------|-------------------------------|--------------|--------------|\n")

	for _, result := range results {
		concurrency := result[0].(int)
		generationSpeed := result[1].(float64)
		promptThroughput := result[2].(float64)
		minTTFT := result[3].(float64)
		maxTTFT := result[4].(float64)
		file.WriteString(fmt.Sprintf("| %11d | %32.2f | %29.2f | %12.2f | %12.2f |\n",
			concurrency,
			generationSpeed,
			promptThroughput,
			minTTFT,
			maxTTFT))
	}

	fmt.Printf("Results saved to: %s\n\n", filename)
}
