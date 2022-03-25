# beer-store-gin-ent

This is a modern go project implementing the beerservice.

It differs from [the other project](../beer-store-service-go-martini-gorm/README.md)
mainly because by the time it was made, the go language relied on
[go workspace](https://go.dev/doc/gopath_code) and now the official go
recommendation for project setup is [go modules](https://go.dev/doc/tutorial/create-module).

## libraries

- [gin](https://github.com/gin-gonic/gin) as our middleware
- [ent](https://github.com/ent/ent) as our ORM
- [sql-migrate](https://github.com/rubenv/sql-migrate) as our migration library

## how to run this project

Since it's a modern go project organized with go modules, just do:

```bash
go run main.go
```

### trivia

Since rosetta-beerstore is a monorepo/multiproject, module name got very funky.
Therefore, unlike workspace-based go projects, internal modules must be imported
using they full names.

So this:

```go
package main

import (
  "fmt"

  "./service"
)

func main() {
  service.Start()
}
```

Becomes this:

```go
package main

import (
  "fmt"

  "github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service"
)

func main() {
  service.Start()
}
```

Also, dependency resolution got pretty simple. Just add the dependency you want:

```go
package models

import "entgo.io/ent"

type Beer struct {
  ent.Schema
}
```

And then issue this command on terminal:

```bash
$ go mod tidy
go: finding module for package entgo.io/ent
go: downloading entgo.io/ent v0.10.1
go: found entgo.io/ent in entgo.io/ent v0.10.1
# ...
```

It may take a while but it works.

It will download and pin the dependency version for you in `go.mod`.
