.PHONY: test clean build

build: tetris.go

run: build
	./tetris.go

test:
	go vet ./...
	go test ./...

clean:
	go clean ./...

tetris.go: *.go
	go build

