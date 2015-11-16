# tetris-go

[![Build Status](https://travis-ci.org/waltzofpearls/tetris-go.svg)](https://travis-ci.org/waltzofpearls/tetris-go)

My personal website [tetris] built with the following technologies:

Languages:

* Golang
* JavaScript
* Less

Frameworks:

* Backbone.js
* Bootstrap

Libraries:

* Gorilla mux
* goquery
* D3.js
* jQuery
* Underscore
* RequireJS
* Cal-HeatMap

Testing:

* Go Test
* Karma
* Jasmine

Building:

* make
* Docker

Package managers:

* npm
* Bower

## Getting started

To clone, build and run the website, do as follows:

```
git clone git@github.com:waltzofpearls/tetris-go.git \
    $GOPATH/src/github.com/waltzofpearls/tetris-go
cd $GOPATH/src/github.com/waltzofpearls/tetris-go
make
./tetris-go
```

To get Projects page properly working, a valid github personal access token
is needed, and it needs to be placed in the config file `config.json` at
`github.token`.

## Testing

```
make test
```
