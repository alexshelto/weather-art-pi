# Makefile for Go project

# Default target: run the application with an optional file flag
run:
	go run cmd/weather-art-pi/main.go 

# Run with a specific file, e.g., make run-file file=testdata/sample.json
run-file:
	go run cmd/weather-art-pi/main.go -file=$(file)

