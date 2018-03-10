#
# build, dev and clean steps
# also, be sure that your $GOROOT is set
#
all: 
	go build beerservice.go

clean: 
	rm -rf beerservice
	rm -rf gin-bin

dev:  
	go get -v github.com/codegangsta/gin
	go install github.com/codegangsta/gin
	gin beerservice.go