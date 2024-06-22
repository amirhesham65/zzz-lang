BIN := bin/hera-lang

.DEFAULT_GOAL := run

.PHONY:fmt vet build clean run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: 
	go test ./...

build: vet
	go build -o $(BIN)

clean:
	rm $(BIN)

run: build
	./$(BIN)