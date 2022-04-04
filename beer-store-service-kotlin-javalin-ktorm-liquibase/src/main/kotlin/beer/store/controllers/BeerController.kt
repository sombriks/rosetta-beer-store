package beer.store.controllers

import beer.store.config.Db
import beer.store.models.Beers
import io.javalin.http.Context
import org.ktorm.dsl.*
import org.ktorm.entity.find
import org.ktorm.entity.sequenceOf

object BeerController {

    const val BEER_LIST = "/beer/list"
    const val BEER_FIND = "/beer/{idbeer}"

    fun list(ctx: Context) {

        val search = ctx.queryParam("search") ?: ""
        val page = (ctx.queryParam("page") ?: "1").toInt()
        val pageSize = (ctx.queryParam("pageSize") ?: "10").toInt()

        val rows = Db.database.from(Beers).select()
            .whereWithOrConditions {
                it += Beers.descriptionBeer like "%$search%"
                it += Beers.titleBeer like "%$search%"
            }
            .orderBy(Beers.idBeer.asc())
            .limit(pageSize)
            .offset((page - 1) * pageSize)
            .map { Beers.createEntity(it) }

        ctx.json(rows)
    }

    fun find(ctx: Context) {
        val id = (ctx.pathParam("idbeer") ?: "0").toInt()
        val beer = Db.database.sequenceOf(Beers).find { it.idBeer eq id }
        if (beer != null) {
            ctx.json(beer)
        } else {
            ctx.status(404).result("NOT_FOUND")
        }
    }
}