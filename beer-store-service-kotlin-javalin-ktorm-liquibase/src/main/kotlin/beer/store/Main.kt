package beer.store

import beer.store.config.Db
import beer.store.controllers.BeerController
import beer.store.controllers.MediaContoller
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.datatype.jsr310.JSR310Module
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule
import com.fasterxml.jackson.module.kotlin.KotlinModule
import io.javalin.Javalin
import io.javalin.plugin.json.JavalinJackson
import org.ktorm.jackson.KtormModule

// service entry point
fun main() {
    // ensure database is sane
    Db.migrate()

    // custom config to make ktor and jackson behave
    val mapper = ObjectMapper()
    mapper.registerModule(JavaTimeModule())
    mapper.registerModule(KotlinModule.Builder().build())
    mapper.registerModule(KtormModule())

    // spin up app
    val app = Javalin.create {
        it.jsonMapper(JavalinJackson(mapper))
    }.start(3000)

    // routes
    app.get(BeerController.BEER_LIST) { BeerController.list(it) }
    app.get(BeerController.BEER_FIND) { BeerController.find(it) }

    app.get(MediaContoller.MEDIA_LIST) { MediaContoller.list(it) }
    app.get(MediaContoller.MEDIA_FIND) { MediaContoller.find(it) }

}