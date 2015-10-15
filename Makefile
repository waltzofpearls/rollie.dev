.PHONY: test clean build

build: tetris.go node_modules

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

node_modules:
	npm install bower
	node_modules/.bin/bower install

