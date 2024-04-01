.PHONY: run clean test build
BIN="VMTranslator"

build:
	go build -o $(BIN) main.go
	chmod +x $(BIN)
clean:
	rm -rf $(BIN)

submit: clean
	mkdir project08-jamesshima
	cp -r main.go Makefile LANGINFO README.md project08-jamesshima
	zip -r project08-jamesshima.zip project08-jamesshima
	echo "ready for canvas"


