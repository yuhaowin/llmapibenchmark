@echo off

echo Creating bin directory...
if not exist bin mkdir bin

echo Building for Windows (amd64 - Standard x86-64 systems)...
set GOOS=windows
set GOARCH=amd64
go build -o bin\llmapibenchmark_windows_amd64.exe .\cmd\main.go

echo Building for Windows (arm64 - Windows on ARM devices)...
set GOOS=windows
set GOARCH=arm64
go build -o bin\llmapibenchmark_windows_arm64.exe .\cmd\main.go

echo Build complete! Binaries are in the bin directory.
pause
