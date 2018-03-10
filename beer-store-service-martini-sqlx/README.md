# beer-store-service-martini-dotsql

This project shows how the beerservice can be implemented using the go 
ecosystem

The [martini](https://github.com/go-martini/martini) web framework provides 
the REST API.

The [sqlx](github.com/jmoiron/sqlx) library extends the go's core sql support.
It has a nice guide [here](http://jmoiron.github.io/sqlx/). 

We also use the [gin](github.com/codegangsta/gin) watcher to save time while 
developing this service

## How do i run this?

Open a terminal on the folder coltaining this readme and 

```bash
make dev
```

After a few errors pop out, use the *go get -v * thing to install every 
missing dependency

Make sure you have the go installed correctly on your system and provide 
reasonable values for **GOROOT** and **GOPATH** variables.
