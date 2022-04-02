package beer.store.models

import org.ktorm.entity.Entity
import java.util.Date

interface Media : Entity<Media> {
    companion object : Entity.Factory<Media>()

    val idMedia: Int //    idmedia           integer primary key autoincrement,
    val creationDateMedia: Date //    creationdatemedia timestamp    not null default CURRENT_TIMESTAMP,
    val dataMedia: ByteArray //    datamedia         blob         not null,
    val nomeMedia: String //    nomemedia         varchar(255) not null,
    val mimeMedia: String //    mimemedia         varchar(255) not null
}