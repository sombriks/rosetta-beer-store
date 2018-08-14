#
# https://www.youtube.com/watch?v=IrYaW3zc9ds
#
export GOPATH := $(PWD)
export GO := $(shell which go)
export GOINSTALL := $(GO) install
export GOCLEAN := $(GO) clean 

bin/beerservice: src/beerservice.go
	go build src/beerservice.go -o bin/beerservice

all: bin/beerservice

clean: 
	rm -rf beerservice
	rm -rf gin-bin

dev:  
	go get -v github.com/codegangsta/gin
	go install github.com/codegangsta/gin
	go get -v
	gin beerservice.go

