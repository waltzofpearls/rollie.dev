.PHONY: test clean build

JS_DIR := static/javascripts
CSS_DIR := static/stylesheets
BUILD_OBJECTS := config.json tetris-go $(JS_DIR)/dist $(CSS_DIR)/dist
NODE_BIN := node_modules/.bin

build: $(BUILD_OBJECTS)

run: build
	./tetris-go

test:
	go vet ./...
	go test ./...

maintainer-clean:
	go clean ./...
	rm -f tetris-go
	rm -rf node_modules
	rm -rf $(JS_DIR)/bower
	rm -f $(JS_DIR)/dist/*.min.js
	rm -f $(JS_DIR)/dist/*.min.js.map
	rm -f $(CSS_DIR)/dist/*.css
	rm -f $(CSS_DIR)/dist/*.css.map

config.json:
	cp -f config.json-dist config.json

tetris-go:
	go get ./...
	go build

$(JS_DIR)/dist: $(JS_DIR)/bower
	$(NODE_BIN)/uglifyjs \
		$(JS_DIR)/bower/requirejs/require.js \
		-o $(JS_DIR)/dist/require.min.js
	$(NODE_BIN)/r.js -o \
		name=main \
		baseUrl=$(JS_DIR)/src/ \
		mainConfigFile=$(JS_DIR)/src/config.js \
		out=$(JS_DIR)/dist/main.min.js \
		preserveLicenseComments=false \
		findNestedDependencies=true \
		optimize=uglify2 \
		generateSourceMaps=true \
		paths.ga=empty:

$(CSS_DIR)/dist: node_modules
	$(NODE_BIN)/lessc \
		--clean-css \
		--source-map=$(CSS_DIR)/dist/style.css.map \
		$(CSS_DIR)/src/style.less \
		$(CSS_DIR)/dist/style.css

$(JS_DIR)/bower: node_modules
	$(NODE_BIN)/bower install

node_modules:
	npm install
