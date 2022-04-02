package beer.store.models

import org.ktorm.dsl.QueryRowSet

import org.ktorm.schema.*

import java.time.LocalDateTime

// https://www.ktorm.org/en/entities-and-column-binding.html
// https://www.ktorm.org/en/define-entities-as-any-kind-of-classes.html
data class Beer(
    var idBeer: Int?,
    var creationDateBeer: LocalDateTime?,
    var titleBeer: String?,
    var descriptionBeer: String?,
    var idMedia: Int?,
)

object Beers : BaseTable<Beer>("beer") {
    var idBeer = int("idbeer").primaryKey()
    var creationDateBeer = datetime("creationdatebeer")
    var titleBeer = varchar("titlebeer")
    var descriptionBeer = varchar("descriptionbeer")
    var idMedia = int("idmedia")

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