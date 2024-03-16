.PHONY: build run clean
BIN_NAME=vm_translator

build:
	@go build -o $(BIN_NAME) main.go
	@chmod +x $(BIN_NAME)
