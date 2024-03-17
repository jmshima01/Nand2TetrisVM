.PHONY: build run clean submit
BIN_NAME=Virtual

build:
	@go build -o $(BIN_NAME) main.go
	@chmod +x $(BIN_NAME)

clean:
	rm -rf $(BIN_NAME)

submit: clean
	mkdir project07-jamesshima
	cp -r main.go Makefile LANGINFO test.sh README.md project07-jamesshima
	zip -r project07-jamesshima.zip project07-jamesshima
	echo "ready for canvas"
