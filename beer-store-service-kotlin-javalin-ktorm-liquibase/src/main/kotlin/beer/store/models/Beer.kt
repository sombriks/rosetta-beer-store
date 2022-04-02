package beer.store.models

import org.ktorm.entity.Entity
import java.util.Date

// https://www.ktorm.org/en/entities-and-column-binding.html
interface Beer : Entity<Beer> {
    companion object : Entity.Factory<Beer>()

    var idBeer: Int//    idbeer           integer primary key autoincrement,
    var creationDateBeer: Date//    creationdatebeer timestamp not null default CURRENT_TIMESTAMP,
    var titleBeer: String//    titlebeer        varchar(255),
    var descriptionBeer: String //    descriptionbeer  text,
    var idMedia: Int//    idmedia          integer,
}