# beer-store-service-kotlin-javalin-ktorm-liquibase

Sample service using kotlin on backend and other fancy things.

- [Javalin](https://javalin.io/) as service controller
- [Ktorm](https://www.ktorm.org/) as ORM and query builder
- [HikariCP](https://github.com/brettwooldridge/HikariCP) for database pooling
  and connection
- [Liquibase](https://liquibase.org/) for database migrations

Project requires Java11 or newer to run

## How do I run this

It uses [gradle](https://gradle.org/) as build tool, so:

```bash
gradle run
```

will build the project and run it

## trivia

Liquibase and HikariCP aren't kotlin-native libraries. However, one of biggest
kotlin advantages is to be able to use everything java ecosystem has to offer.

Unlike [other](http://knexjs.org/#Migrations)
[fancy](https://alembic.sqlalchemy.org/en/latest/tutorial.html#create-a-migration-script)
[migration](https://alembic.sqlalchemy.org/en/latest/tutorial.html#create-a-migration-script)
tools, Liquibase does not offer a tool to create migrate templates. I covered it
in [this article](https://sombriks.com.br/#/blog/0025-migrations-with-liquibase-and-sql.md)
about what liquibase can do.

