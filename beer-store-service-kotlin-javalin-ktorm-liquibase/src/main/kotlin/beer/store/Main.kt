package beer.store

import beer.store.controllers.Beers
import io.javalin.Javalin
import liquibase.Liquibase
import liquibase.database.DatabaseConnection

fun main() {
    val app = Javalin.create().start(3000)
    app.get("/beer/list") { Beers.list(it) }
    app.get("/beer/{idbeer}") { Beers.find(it) }
}