package beer.store

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView

class MainActivity : AppCompatActivity() {

    var beerAdapter :BeerAdapter = BeerAdapter()

    override fun onCreate(savedInstanceState: Bundle?) {
        title = "Beer Store"
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        var rview = findViewById<RecyclerView>(R.id.recycler_view)
        rview.layoutManager = LinearLayoutManager(rview.context)
        rview.adapter = beerAdapter
        Log.i("beer.store","aaa")
    }
}
