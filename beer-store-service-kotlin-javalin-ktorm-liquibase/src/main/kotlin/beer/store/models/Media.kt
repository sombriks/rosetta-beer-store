package beer.store.models

import org.ktorm.entity.Entity
import org.ktorm.schema.*
import java.time.LocalDateTime

interface Media : Entity<Media> {
    companion object : Entity.Factory<Media>()

    val idMedia: Int
    val creationDateMedia: LocalDateTime
    val dataMedia: ByteArray
    val nomeMedia: String
    val mimeMedia: String
}

object Medias : Table<Media>("media") {
    val idMedia = int("idmedia").primaryKey().bindTo { it.idMedia }
    val creationDateMedia =
        datetime("creationdatemedia").bindTo { it.creationDateMedia }
    val dataMedia = blob("datamedia").bindTo { it.dataMedia }
    val nomeMedia = varchar("nomemedia").bindTo { it.nomeMedia }
    val mimeMedia = varchar("mimemedia").bindTo { it.mimeMedia }
}