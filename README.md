# LLM API Benchmark Tool

## Overview

The LLM API Benchmark Tool is a flexible Go-based utility designed to measure and analyze the performance of OpenAI-compatible API endpoints across different concurrency levels. This tool provides in-depth insights into API throughput, generation speed, and token processing capabilities.

## Download and Installation
```bash
# Download and grant execute permissions
wget https://github.com/yuhaowin/llmapibenchmark/releases/download/v1.0.1/llmapibenchmark_darwin_arm64 -O ./llmapibenchmark_darwin_arm64
chmod +x ./llmapibenchmark_darwin_arm64
```

## Build Instructions

### Prerequisites
- Go 1.20 or higher
- Git

### Building from Source

**macOS/Linux:**
```bash
chmod +x build.sh
./build.sh
```

**Windows:**
```cmd
build.bat
```

The built binaries will be placed in the `bin` directory. The Linux builds (`amd64` and `arm64`) are suitable for most common server environments.
- macOS (Intel): `bin/llmapibenchmark_darwin_amd64`
- macOS (Apple Silicon): `bin/llmapibenchmark_darwin_arm64`
- Linux (amd64): `bin/llmapibenchmark_linux_amd64`
- Linux (arm64): `bin/llmapibenchmark_linux_arm64`
- Windows (amd64): `bin/llmapibenchmark_windows_amd64.exe`
- Windows (arm64): `bin/llmapibenchmark_windows_arm64.exe`

## Key Features

- 🚀 Dynamic Concurrency Testing
- 📊 Comprehensive Performance Metrics
- 🔍 Flexible Configuration
- 📝 Markdown Result Reporting
- 🌐 Compatible with Any OpenAI-Like API
- 📏 Arbitrary Length Dynamic Input Prompt

## Performance Metrics Measured

1. **Generation Throughput**
   - Measures tokens generated per second
   - Calculates across multiple concurrency levels

2. **Prompt Throughput**
   - Analyzes input token processing speed
   - Helps understand API's prompt handling efficiency

3. **Time to First Token (TTFT)**
   - Measures initial response latency
   - Provides both minimum and maximum TTFT
   - Critical for understanding real-time responsiveness

## Example Output
```
Input Tokens: 45
Output Tokens: 512
Test Model: Qwen2.5-7B-Instruct-AWQ
Latency: 2.20 ms
```

| Concurrency | Generation Throughput (tokens/s) |  Prompt Throughput (tokens/s) | Min TTFT (s) | Max TTFT (s) |
|-------------|----------------------------------|-------------------------------|--------------|--------------|
|           1 |                            58.49 |                        846.81 |         0.05 |         0.05 |
|           2 |                           114.09 |                        989.94 |         0.08 |         0.09 |
|           4 |                           222.62 |                       1193.99 |         0.11 |         0.15 |
|           8 |                           414.35 |                       1479.76 |         0.11 |         0.24 |
|          16 |                           752.26 |                       1543.29 |         0.13 |         0.47 |
|          32 |                           653.94 |                       1625.07 |         0.14 |         0.89 |


### Minimal Configuration

**macOS:**
```bash
# For Intel Macs
./bin/llmapibenchmark_darwin_amd64 -base_url=https://your-api-endpoint.com/v1
# For Apple Silicon Macs
./bin/llmapibenchmark_darwin_arm64 -base_url=https://your-api-endpoint.com/v1
```

**Linux:**
```bash
# For amd64 servers/desktops
./bin/llmapibenchmark_linux_amd64 -base_url=https://your-api-endpoint.com/v1
# For arm64 servers/desktops
./bin/llmapibenchmark_linux_arm64 -base_url=https://your-api-endpoint.com/v1
```

**Windows:**
```cmd
REM For amd64 systems
bin\llmapibenchmark_windows_amd64.exe -base_url=https://your-api-endpoint.com/v1
REM For arm64 systems
bin\llmapibenchmark_windows_arm64.exe -base_url=https://your-api-endpoint.com/v1
```

### Full Configuration

Replace `./llmapibenchmark_os_arch` or `llmapibenchmark_os_arch.exe` with the appropriate binary from the `bin` directory for your system.

**Example for macOS (Apple Silicon):**
```bash
./bin/llmapibenchmark_darwin_arm64 \
  -base_url=https://your-api-endpoint.com/v1 \
  -apikey=YOUR_API_KEY \
  -model=gpt-3.5-turbo \
  -concurrency=1,2,4,8,16 \
  -max_tokens=512 \
  -numWords=513 \
  -prompt="Your custom prompt here"
```

**Example for Windows (amd64):**
```cmd
bin\llmapibenchmark_windows_amd64.exe ^
  -base_url=https://your-api-endpoint.com/v1 ^
  -apikey=YOUR_API_KEY ^
  -model=gpt-3.5-turbo ^
  -concurrency=1,2,4,8,16 ^
  -max_tokens=512 ^
  -numWords=513 ^
  -prompt="Your custom prompt here"
```

## Command-Line Parameters

| Parameter         | Description                                | Default                                                                                 | Required |
|-------------------|--------------------------------------------|-----------------------------------------------------------------------------------------|----------|
| `-base_url`       | Base URL for LLM API endpoint              | Empty (MUST be specified)                                                               | Yes      |
| `-apikey`         | API authentication key                     | None                                                                                    | No       |
| `-model`          | Specific AI model to test                  | Automatically discovers first available model                                           | No       |
| `-concurrency`    | Comma-separated concurrency levels to test | `1,2,4,8,16,32,64,128`                                                                  | No       |
| `-max_tokens`     | Maximum tokens to generate per request     | `512`                                                                                   | No       |
| `-numWords`       | Number of words for input prompt           | Not set (optional)                                                                      | No       |
| `-prompt`         | Text prompt for generating responses       | `"Write a long story, no less than 10,000 words, starting from a long, long time ago."` | No       |

## Output

The tool generates:
1. Console-based real-time results
2. Markdown file (`API_Throughput_{ModelName}.md`) with detailed results

### Result File Columns

- **Concurrency**: Number of concurrent requests
- **Generation Throughput**: Tokens generated per second
- **Prompt Throughput**: Input token processing speed
- **Min TTFT**: Minimum time to first token
- **Max TTFT**: Maximum time to first token

## Best Practices

- Test with various prompt lengths and complexities
- Compare different models
- Monitor for consistent performance
- Be mindful of API rate limits
- Use `-numWords` to control input length

## Limitations

- Requires active API connection
- Results may vary based on network conditions
- Does not simulate real-world complex scenarios

## Disclaimer

This tool is for performance analysis and should be used responsibly in compliance with API provider's usage policies.