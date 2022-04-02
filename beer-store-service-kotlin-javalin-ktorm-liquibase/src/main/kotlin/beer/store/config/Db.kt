package beer.store.config

import com.zaxxer.hikari.HikariConfig
import com.zaxxer.hikari.HikariDataSource

import liquibase.Liquibase
import liquibase.database.DatabaseFactory
import liquibase.database.jvm.JdbcConnection
import liquibase.resource.ClassLoaderResourceAccessor

import org.ktorm.database.Database


object Db {

    private var ds = HikariDataSource(HikariConfig("/application.properties"))

    var database = Database.Companion.connect(ds)
        private set

    fun migrate(){
        val database = DatabaseFactory.getInstance()
            .findCorrectDatabaseImplementation(JdbcConnection(ds.connection))
        Liquibase("changelog.xml",ClassLoaderResourceAccessor(),database).update("")
    }
}