# beer-store-service-kotlin-javalin-ktorm-liquibase

Sample service using kotlin on backend and other fancy things.

- [Javalin](https://javalin.io/) as service controller
- [Ktorm](https://www.ktorm.org/) as ORM and query builder
- [HikariCP](https://github.com/brettwooldridge/HikariCP) for database pooling
  and connection
- [Liquibase](https://liquibase.org/) for database migrations

Project requires Java 11 or newer to run

## How do I run this

It uses [gradle](https://gradle.org/) as build tool, so:

```bash
gradle run
```

will build the project and run it.

You also can open the project on Intellij Ultimate, everything will feel like
_home_ :-) 

## Trivia

Liquibase and HikariCP aren't kotlin-native libraries. However, one of biggest
kotlin advantages is to be able to use everything java ecosystem has to offer.

Unlike [other](http://knexjs.org/#Migrations)
[fancy](https://alembic.sqlalchemy.org/en/latest/tutorial.html#create-a-migration-script)
[migration](https://alembic.sqlalchemy.org/en/latest/tutorial.html#create-a-migration-script)
tools, Liquibase does not offer a tool to create migrate templates. I covered it
in [this article](https://sombriks.com.br/#/blog/0025-migrations-with-liquibase-and-sql.md)
about what liquibase can do.

### jackson and ktorm had a fight

jackson had a difficult time serializing ktorm proxies when I used the
_recommended_ strategy for object-relational mapping. i left `Media` mapping as
an example of what ktorm recommends.

However, javalin endpoint keeps throwing exceptions and no research produced a
solution.

Except when I used the alternative mapping method offered by ktorm:

```kotlin
package beer.store.models

import org.ktorm.dsl.QueryRowSet

import org.ktorm.schema.*

import java.time.LocalDateTime

// https://www.ktorm.org/en/entities-and-column-binding.html
// https://www.ktorm.org/en/define-entities-as-any-kind-of-classes.html
data class Beer(
    var idBeer: Int?,
    var creationDateBeer: LocalDateTime?,
    var titleBeer: String?,
    var descriptionBeer: String?,
    var idMedia: Int?,
)

object Beers : BaseTable<Beer>("beer") {
    var idBeer = int("idbeer").primaryKey()
    var creationDateBeer = datetime("creationdatebeer")
    var titleBeer = varchar("titlebeer")
    var descriptionBeer = varchar("descriptionbeer")
    var idMedia = int("idmedia")

    override fun doCreateEntity(
        row: QueryRowSet,
        withReferences: Boolean
    ) = Beer(
        idBeer = row[idBeer],
        creationDateBeer = row[creationDateBeer],
        titleBeer = row[titleBeer],
        descriptionBeer = row[descriptionBeer],
        idMedia = row[idMedia]
    )
}
```

Mapping got way more verbose, but it doesn't use dynamic proxies anymore and
then jackson started to work again.

So far the least pleasant ORM framework is ktorm.

## Would I use this in production

Kotlin/Javalin/Liquibase? Sure!
