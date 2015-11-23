.PHONY: test clean clean-go clean-js clean-css distclean build
.PHONY: docker docker-build docker-run docker-purge

JS_DIR := static/javascripts
CSS_DIR := static/stylesheets
BUILD_OBJECTS := config.json tetris-go $(JS_DIR)/dist/require.min.js \
	$(JS_DIR)/dist/main.min.js $(CSS_DIR)/dist/style.css
NODE_BIN := node_modules/.bin
DOCKER_IMG := tetris-go-image
DOCKER_CON := tetris-go-container
DOCKER_PRT := 49002

build: $(BUILD_OBJECTS)

run: build
	./tetris-go

test:
	go vet ./...
	go test ./...

clean: clean-go clean-js clean-css

clean-go:
	go clean ./...
	rm -f tetris-go

clean-js:
	rm -f $(JS_DIR)/dist/*.min.js
	rm -f $(JS_DIR)/dist/*.min.js.map

clean-jspkg:
	rm -rf node_modules
	rm -rf $(JS_DIR)/bower

clean-css:
	rm -f $(CSS_DIR)/dist/*.css
	rm -f $(CSS_DIR)/dist/*.css.map

distclean: clean-go clean-js clean-jspkg clean-css

config.json:
	cp -f config.json-dist config.json

tetris-go:
	go get ./...
	go build

$(JS_DIR)/dist/require.min.js: $(JS_DIR)/bower
	$(NODE_BIN)/uglifyjs \
		$(JS_DIR)/bower/requirejs/require.js \
		-o $(JS_DIR)/dist/require.min.js

$(JS_DIR)/dist/main.min.js: $(JS_DIR)/bower
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

$(CSS_DIR)/dist/style.css: node_modules
	$(NODE_BIN)/lessc \
		--clean-css \
		--source-map=$(CSS_DIR)/dist/style.css.map \
		$(CSS_DIR)/src/style.less \
		$(CSS_DIR)/dist/style.css

$(JS_DIR)/bower: node_modules
	yes y | $(NODE_BIN)/bower install --allow-root

node_modules:
	npm install

docker: docker-purge docker-build docker-run

docker-build: build
	docker build \
		-t $(DOCKER_IMG):latest \
		-f Dockerfile .

docker-run:
	docker run -d \
		--name $(DOCKER_CON) \
		-p $(DOCKER_PRT):3000 \
		--env-file ./env.list \
		$(DOCKER_IMG):latest

docker-purge:
	docker ps -a | grep $(DOCKER_CON) > /dev/null \
		&& docker kill $(DOCKER_CON) \
		&& docker rm -v $(DOCKER_CON)
	docker images | grep $(DOCKER_IMG) > /dev/null \
		&& docker rmi $(DOCKER_IMG)
