#
# https://www.youtube.com/watch?v=IrYaW3zc9ds
#
export GOPATH := $(PWD)
export GOBIN := $(GOPATH)/bin
export PATH := $(PATH):$(GOBIN)
export PRG := beerservice

.PHONY: all

all: build deps

dev:
	echo "you should 'make deps' at least once"
	$(GOBIN)/bra run

build:
	echo "you should 'make deps' at least once"
	go install src/main/$(PRG).go

deps: bin pkg src
	# regular dependencies
	go get -v github.com/go-martini/martini
	go get -v github.com/martini-contrib/cors
	go get -v github.com/martini-contrib/render
	# go get -v github.com/jmoiron/sqlx
	go get -v github.com/jinzhu/gorm
	go get -v github.com/mattn/go-sqlite3
	# dev dependencies
	go get -v github.com/mdempsky/gocode
	go get -v github.com/uudashr/gopkgs/cmd/gopkgs
	go get -v github.com/ramya-rao-a/go-outline
	go get -v github.com/acroca/go-symbols
	go get -v golang.org/x/tools/cmd/guru
	go get -v golang.org/x/tools/cmd/gorename
	go get -v github.com/derekparker/delve/cmd/dlv
	go get -v github.com/rogpeppe/godef
	go get -v golang.org/x/tools/cmd/godoc
	go get -v github.com/golang/lint/golint
	go get -v github.com/sqs/goreturns
	go get -v github.com/Unknwon/bra

bin:
	mkdir bin

pkg:
	mkdir pkg

src:
	mkdir src

clean: 
	rm -rf bin pkg src/github.com src/gopkg.in src/golang.org beerstore.sqlite3 
