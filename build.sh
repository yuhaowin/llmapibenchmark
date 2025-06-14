#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Create bin directory if it doesn't exist
echo "Creating bin directory..."
mkdir -p bin

# Build for macOS (amd64 - Intel-based Macs)
echo "Building for macOS (amd64 - Intel)..."
GOOS=darwin GOARCH=amd64 go build -o bin/llmapibenchmark_darwin_amd64 ./cmd/main.go

# Build for macOS (arm64 - Apple Silicon Macs)
echo "Building for macOS (arm64 - Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o bin/llmapibenchmark_darwin_arm64 ./cmd/main.go

# Build for Linux (amd64 - Common for x86-64 servers and desktops)
echo "Building for Linux (amd64 - x86-64)..."
GOOS=linux GOARCH=amd64 go build -o bin/llmapibenchmark_linux_amd64 ./cmd/main.go

# Build for Linux (arm64 - Common for ARM64 servers like AWS Graviton and desktops)
echo "Building for Linux (arm64 - ARMv8)..."
GOOS=linux GOARCH=arm64 go build -o bin/llmapibenchmark_linux_arm64 ./cmd/main.go

echo "Build complete! Binaries are in the bin directory."