package beer.store.models

import org.ktorm.dsl.QueryRowSet

import org.ktorm.schema.*

import java.time.LocalDateTime
import javax.swing.text.html.parser.Entity

interface Beer : Entity <Beer> (
    companion object : Entity.Factory<Beer>()

    val idBeer: Int
    val creationDateBeer: LocalDateTime
    val titleBeer: String
    val descriptionBeer: String
    val idMedia: Int
)

object Beers : Table<Beer>("beer") {
    var idBeer = int("idbeer").primaryKey().bindTo { it.idBeer }
    var creationDateBeer = datetime("creationdatebeer").bindTo { it.creationDateBeer }
    var titleBeer = varchar("titlebeer").bindTo { it.titleBeer }
    var descriptionBeer = varchar("descriptionbeer").bindTo { it.titleBeer }
    var idMedia = int("idmedia").bindTo { it.idMedia }

    override fun doCreateEntity(
        row: QueryRowSet,
        withReferences: Boolean
    ) = Beer (
        idBeer = row[idBeer],
        creationDateBeer = row[creationDateBeer],
        titleBeer = row[titleBeer],
        descriptionBeer = row[descriptionBeer],
        idMedia = row[idMedia]
    )
}