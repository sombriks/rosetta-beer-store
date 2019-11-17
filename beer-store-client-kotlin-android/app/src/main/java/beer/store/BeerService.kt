package beer.store

import com.google.gson.GsonBuilder
import retrofit2.Call
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET
import retrofit2.http.Path

object BeerService {

    interface Service {

        @GET("/beer/list")
        fun list(): Call<List<Beer>>

        @GET("/beer/{id}")
        fun find(@Path("id") id: Int): Call<Beer>
    }

    val gson = GsonBuilder().setDateFormat("yyyy-MM-dd HH:mm:ss").create()

    fun client(url: String): Retrofit = Retrofit.Builder().baseUrl(url)//
        .addConverterFactory(GsonConverterFactory.create(gson)).build()

    fun instance(url: String = "http://192.168.0.111:3000"): Service =
        client(url).create(Service::class.java)

}