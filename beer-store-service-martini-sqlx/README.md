# beer-store-service-martini-dotsql

This project shows how the beerservice can be implemented using the go 
ecosystem

The [martini](https://github.com/go-martini/martini) web framework provides 
the REST API.

~~The [sqlx](github.com/jmoiron/sqlx) library extends the go's core sql support.~~
~~It has a nice guide [here](http://jmoiron.github.io/sqlx/).~~ 

The [Gorm](http://gorm.io/) library does the database thing for us.

Also, we use [bra](https://github.com/Unknwon/bra) to watch file changes. 
It needs a config file called **.bra.toml** and once correctly placed works 
like a charm.

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
