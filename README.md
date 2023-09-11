# rollie.dev

My personal website. Built with Rust + Svelte.

* Rust:
  * Actix Web
  * Reqwest
  * GraphQL Client
* Svelte:
  * Svelte Material UI
* And with:
  * Vite
  * TypeScript
  * Bootstrap
  * Cal-HeatMap

## Getting started

Build and run for development with file watcher:

```shell
make dev
```

Build the production container and run it locally:

```shell
make docker
```

NOTE: before running the command, create .env file with env vars

```shell
GITHUB_TOKEN=_github_personal_access_token_
LISTEN_HTTP=0.0.0.0:3000
```

## Deploy

Deploy to [fly.io](https://fly.io/) via `Dockerfile`

```shell
make deploy
```
