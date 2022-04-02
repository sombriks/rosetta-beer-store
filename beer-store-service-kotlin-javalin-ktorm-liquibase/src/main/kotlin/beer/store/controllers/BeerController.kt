package beer.store.controllers

import beer.store.config.Db
import io.javalin.http.Context
import org.ktorm.database.asIterable

object BeerController {
    fun list(ctx: Context) {
        val rows = Db.database.useConnection { connection ->
            val q = """
                select * from beer
            """.trimIndent()
            connection.prepareStatement(q).executeQuery().asIterable()
                .map { row ->
                    row.getString(1)
                }
        }
        ctx.result(rows.joinToString(""))
    }

    fun find(ctx: Context) {
        ctx.result("bbb")
    }
}