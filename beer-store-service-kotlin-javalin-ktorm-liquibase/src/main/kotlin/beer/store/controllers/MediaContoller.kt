package beer.store.controllers

import beer.store.config.Db
import beer.store.models.Medias
import io.javalin.http.Context
import org.ktorm.dsl.*
import org.ktorm.entity.*

object MediaContoller {

    const val MEDIA_LIST = "/media/list"
    const val MEDIA_FIND = "/media/{idmedia}"

    fun list(ctx: Context) {

        val search = ctx.queryParam("search") ?: ""
        val page = (ctx.queryParam("page") ?: "1").toInt()
        val pageSize = (ctx.queryParam("pageSize") ?: "10").toInt()

        val result = Db.database.sequenceOf(Medias)
            .filter { it.nomeMedia like "%$search%" }
            .sortedBy { it.idMedia.asc() }
            .drop((page - 1) * pageSize)
            .take(pageSize)
            .toList()

        ctx.json(result)

    }

    fun find(ctx: Context) {
        val id = (ctx.pathParam("idmedia") ?: "0").toInt()
        val result = Db.database.from(Medias).select()
            .where { (Medias.idMedia eq id) }
            .map { Medias.createEntity(it) }.firstOrNull()
        if (result != null) {
            ctx.json(result)
        } else {
            ctx.status(404).result("NOT_FOUND")
        }
    }
}