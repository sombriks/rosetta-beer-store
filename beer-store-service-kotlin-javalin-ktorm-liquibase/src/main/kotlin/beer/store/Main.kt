package beer.store

import io.javalin.Javalin
import liquibase.Liquibase

fun main() {
    val liquibase = Liquibase()
    println("aaaa")
    val app = Javalin.create().start(3000)
}