.PHONY: test clean build

build: tetris.go static/javascripts/bower/

run: build
	./tetris.go

test:
	go vet ./...
	go test ./...

clean:
	go clean ./...
	rm -f tetris.go
	rm -rf node_modules/
	rm -rf static/javascripts/bower/

tetris.go:
	go build

static/javascripts/bower/: node_modules/
	node_modules/.bin/bower install

node_modules/:
	npm install

