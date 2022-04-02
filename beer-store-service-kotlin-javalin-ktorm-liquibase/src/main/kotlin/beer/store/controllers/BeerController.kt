package beer.store.controllers

import beer.store.config.Db
import beer.store.models.Beer
import beer.store.models.Beers
import io.javalin.http.Context
import org.ktorm.entity.map
import org.ktorm.entity.sequenceOf
import org.ktorm.entity.toList

object BeerController {
    fun list(ctx: Context) {
        val rows = Db.database.sequenceOf(Beers).toList()
        ctx.json(rows)
    }

    fun find(ctx: Context) {
        ctx.result("bbb")
    }
}