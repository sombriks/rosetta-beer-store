# beer-store-service-martini-dotsql

This project shows how the beerservice can be implemented using the go 
ecosystem

The [martini](https://github.com/go-martini/martini) web framework provides 
the REST API.

The [sqlx](github.com/jmoiron/sqlx) library extends the go's core sql support.
It has a nice guide [here](http://jmoiron.github.io/sqlx/). 

So far, there is no decent tool to act as a development watcher. We tried 
gin (from codegangsta) and go-watcher (from canthefason) but they lack respect
for $(GOPATH) and execution points.

## How do i run this?

Open a terminal on the folder coltaining this readme and 

```bash
make dev
```

The [go workspace](https://golang.org/doc/code.html) once set will manage well 
how libraries, source code and binaries will stay and the project folder will
also be the execution point.

Static and queries folder are there because of execution point.

bin, pkg and src are there because this is also our $GOPATH, as seen in the 
Makefile.
