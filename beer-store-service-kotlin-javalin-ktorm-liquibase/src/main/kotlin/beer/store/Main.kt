package beer.store

import beer.store.config.Db
import beer.store.controllers.BeerController
import io.javalin.Javalin

fun main() {
    Db.migrate()
    val app = Javalin.create().start(3000)
    app.get("/beer/list") { BeerController.list(it) }
    app.get("/beer/{idbeer}") { BeerController.find(it) }
}