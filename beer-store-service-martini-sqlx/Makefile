#
# build, dev and clean steps
#
# when using sistem-wide go, all you need is to append $HOME/go/bin on path
#
export PATH := $(HOME)/go/bin:$(PATH)

all: 
	go get -v
	go build beerservice.go -o beerservice

clean: 
	rm -rf beerservice
	rm -rf gin-bin

dev:  
	go get -v github.com/codegangsta/gin
	go install github.com/codegangsta/gin
	go get -v
	gin beerservice.go