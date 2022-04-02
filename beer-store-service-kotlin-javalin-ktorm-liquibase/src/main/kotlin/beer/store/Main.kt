package beer.store

import beer.store.config.Db
import beer.store.controllers.BeerController
import beer.store.controllers.MediaContoller
import io.javalin.Javalin

// service entry point
fun main() {
    Db.migrate()

    val app = Javalin.create().start(3000)

    app.get(BeerController.BEER_LIST) { BeerController.list(it) }
    app.get(BeerController.BEER_FIND) { BeerController.find(it) }

    app.get(MediaContoller.MEDIA_LIST) { MediaContoller.list(it) }
    app.get(MediaContoller.MEDIA_FIND) { MediaContoller.find(it) }

}