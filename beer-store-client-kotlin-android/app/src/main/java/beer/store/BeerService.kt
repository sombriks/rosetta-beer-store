package beer.store

import com.google.gson.GsonBuilder
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET
import retrofit2.http.Path
import retrofit2.http.Query

object BeerService {

    interface Service {

        @Throws(Exception::class)
        @GET("/beer/list")
        suspend fun list(@Query("search") search: String = "", @Query("page") page: Int = 1, @Query("pageSize") pageSize: Int = 10): MutableList<Beer>

        @Throws(Exception::class)
        @GET("/beer/{id}")
        suspend fun find(@Path("id") id: Int): Beer
    }

    val gson = GsonBuilder().setDateFormat("yyyy-MM-dd HH:mm:ss").create()

    fun client(url: String): Retrofit = Retrofit.Builder().baseUrl(url)//
        .addConverterFactory(GsonConverterFactory.create(gson)).build()

    fun instance(url: String = "http://192.168.0.111:3000"): Service =
        client(url).create(Service::class.java)

}