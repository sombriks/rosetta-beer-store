package beer.store.controllers

import io.javalin.http.Context

object Beers {
    fun list(ctx: Context) {
        ctx.result("aaa")
    }
    fun find(it: Context) {

    }
}