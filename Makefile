.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: dev
dev: ## Run the app in dev mode
	docker compose up --remove-orphans --force-recreate --build

.PHONY: deps
deps: ## Install and update dependencies
	docker compose exec frontend npm install
	docker compose cp backend/Cargo.toml backend:/backend/Cargo.toml
	docker compose exec backend cargo build

.PHONY: build
build: | build-frontend build-backend ## Build for production frontend and backend code

env ?= dev

.PHONY: build-frontend
build-frontend: ## Build for production frontend bundle
	( \
		cd frontend; \
		npm install; \
		npm run build -- --mode $(env) \
	)

.PHONY: build-backend
build-backend: ## Build for production backend binary
	( \
		cd backend; \
		cargo build --release \
	)

IMAGE := rollie.dev
CONTAINER := rollie.dev

.PHONY: docker
docker: ## Build and run the app in docker
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
