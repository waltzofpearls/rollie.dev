# rollie.dev

[![Build Status](https://travis-ci.org/waltzofpearls/rollie.dev.svg)](https://travis-ci.org/waltzofpearls/rollie.dev)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/waltzofpearls/rollie.dev)](https://hub.docker.com/r/waltzofpearls/rollie.dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/waltzofpearls/rollie.dev)](https://goreportcard.com/report/github.com/waltzofpearls/rollie.dev)

My personal website. Built with Golang, Gorilla and Backbone.js.

Libs used:

* Golang
  * Gorilla mux
  * goquery
* JavaScript
  * Backbone.js
  * Bootstrap
  * D3.js
  * jQuery
  * Underscore
  * RequireJS
  * Cal-HeatMap
  * Karma
  * Jasmine

## Getting started

Pull and run from docker hub image:

```
docker pull waltzofpearls/rollie.dev

docker run -d --name rollie.dev -p 3000:3000 --env-file .env waltzofpearls/rollie.dev
```

Build and run from the source:

```
git clone git@github.com:waltzofpearls/rollie.dev.git
cd rollie.dev
make
./rollie.dev
```

To build with docker, replace `make` with `make docker`.

To get Projects page properly working, a valid github personal access token
is needed, and it needs to be placed in the config file `config.json` at
`github.token`.

## Unit testing

```
make test
```
