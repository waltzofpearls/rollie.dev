.PHONY: test clean clean-go clean-js clean-css distclean build
.PHONY: docker docker-build docker-run docker-purge

PKG = $$(go list ./... | grep -v -e '/aggregated')
JS_DIR := static/javascripts
CSS_DIR := static/stylesheets
NODE_BIN := node_modules/.bin
IMAGE := rollie.dev
CONTAINER := rollie.dev

all: build

build: build-frontend build-backend

build-frontend: $(JS_DIR)/dist/require.min.js $(JS_DIR)/dist/main.min.js $(CSS_DIR)/dist/style.css

build-backend: config.json rollie.dev

test:
	go vet $(PKG)
	go test -race -v -cover -run "$(filter)" $(PKG)

clean: clean-go clean-js clean-css

clean-all: clean clean-jspkg

clean-go:
	go clean ./...
	rm -f rollie.dev

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

rollie.dev:
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

dev:
	( \
		docker ps -a | grep $(CONTAINER) > /dev/null && ( \
			docker kill $(CONTAINER); \
			docker rm -v $(CONTAINER) \
		) \
	) || true
	docker build --platform linux/amd64 -t $(IMAGE) .
	docker run -it \
		--platform linux/amd64 \
		--name $(CONTAINER) \
		-p 3000:3000 \
		--env-file .env \
		$(IMAGE):latest

deploy:
	flyctl deploy
