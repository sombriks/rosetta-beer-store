# beer-store-service-martini-gorm

This project shows how the beerservice can be implemented using the go
ecosystem

The [martini](https://github.com/go-martini/martini) web framework provides
the REST API.

The [Gorm](http://gorm.io/) library does the database thing for us.

The [gomigrate](https://github.com/DavidHuie/gomigrate) handles database
migrations.

## How to run this

Open a terminal on the folder containing this readme and

```bash
make build
make dev
```

## Trivia

The [go workspace](https://golang.org/doc/code.html) once set will manage well
how libraries, source code and binaries will stay and the project folder will
also be the execution point.

Static and queries folder are there because of execution point.

bin, pkg and src are there because this is also our \$GOPATH, as seen in the
Makefile.

This project is incompatible with go modules so we have `export GO111MODULE=off`
in the Makefile.

Since newer vscode go extension plugin prefers go modules instead of workspaces,
some tweaks where made into `.vscode/settings.json` in order to solve lots of
false positive error messages. A few remain
