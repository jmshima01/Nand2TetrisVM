.PHONY: build run clean
BIN_NAME=Virtual

build:
	@go build -o $(BIN_NAME) main.go
	@chmod +x $(BIN_NAME)

clean:
	rm -rf $(BIN_NAME)
