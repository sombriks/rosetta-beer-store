package beer.store

import com.google.gson.annotations.SerializedName
import java.util.*

data class Beer(
    @SerializedName("idbeer")
    val idBeer: Int = 0,
    @SerializedName("titlebeer")
    val titleBeer: String = "Sample beer",
    @SerializedName("creationdatebeer")
    val creationDateBeer: Date = Date(),
    @SerializedName("descriptionbeer")
    val descriptionBeer: String = "Sample empty bottle beer",
    @SerializedName("idmedia")
    val idMedia: Int = 0
)