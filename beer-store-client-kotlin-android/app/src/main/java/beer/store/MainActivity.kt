package beer.store

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class MainActivity : AppCompatActivity() {

    var beerAdapter :BeerAdapter = BeerAdapter()

    override fun onCreate(savedInstanceState: Bundle?) {
        title = "Beer Store"
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        var rview = findViewById<RecyclerView>(R.id.recycler_view)
        rview.layoutManager = LinearLayoutManager(rview.context)
        rview.adapter = beerAdapter


        // TODO melhorar isso aqui e conter no service se possível
        BeerService.instance().list().enqueue(object : Callback<List<Beer>> {
            override fun onResponse(call: Call<List<Beer>>, response: Response<List<Beer>>) {
                beerAdapter.beers.removeAll(beerAdapter.beers)
                response.body()?.forEach {
                    Log.i("beer.store",it.toString())
                    beerAdapter.beers.add(it)
                }
            }

            override fun onFailure(call: Call<List<Beer>>, t: Throwable) {
                Log.wtf("beer.store",t.toString())
            }
        })
    }
}
