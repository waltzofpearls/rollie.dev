# rolli3.net

[![Build Status](https://travis-ci.org/waltzofpearls/rolli3.net.svg)](https://travis-ci.org/waltzofpearls/rolli3.net)
[![Go Report Card](https://goreportcard.com/badge/github.com/waltzofpearls/rolli3.net)](https://goreportcard.com/report/github.com/waltzofpearls/rolli3.net)

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
docker pull waltzofpearls/rolli3.net

docker run -d --name rolli3.net -p 3000:3000 --env-file .env waltzofpearls/rolli3.net
```

Build and run from the source:

```
git clone git@github.com:waltzofpearls/rolli3.net.git
cd rolli3.net
make
./rolli3.net
```

To build with docker, replace `make` with `make docker`.

To get Projects page properly working, a valid github personal access token
is needed, and it needs to be placed in the config file `config.json` at
`github.token`.

## Unit testing

```
make test
```
