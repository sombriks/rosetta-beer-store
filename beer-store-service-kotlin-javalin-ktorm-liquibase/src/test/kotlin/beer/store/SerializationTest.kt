package beer.store

import beer.store.config.Db
import beer.store.models.Beers

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.databind.SerializationFeature
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule
import com.google.gson.Gson

import org.ktorm.entity.sequenceOf
import org.ktorm.entity.toList

import kotlin.test.Test

class SerializationTest {

    @Test
    fun shouldSerializeResult() {
        Db.migrate()
        val rows = Db.database.sequenceOf(Beers).toList()

        // yes toString works
        println(rows)

        //
        val gson = Gson()
        println(gson.toJson(rows))

        val mapper = ObjectMapper()
        mapper.registerModule(JavaTimeModule())
        mapper.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);
        println(mapper.writeValueAsString(rows))
    }

}