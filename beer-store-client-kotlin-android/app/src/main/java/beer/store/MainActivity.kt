package beer.store

import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import androidx.appcompat.app.AppCompatActivity
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView

class MainActivity : AppCompatActivity() {

    var page = 1
    var pageSize = 10

    val beerAdapter: BeerAdapter = BeerAdapter()

    override fun onCreate(savedInstanceState: Bundle?) {
        title = "Beer Store"
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val rview = findViewById<RecyclerView>(R.id.beer_recycler_view)
        rview.layoutManager = LinearLayoutManager(rview.context)
        rview.adapter = beerAdapter

        val campoBusca = findViewById<EditText>(R.id.campo_busca)

        val doSearch = findViewById<Button>(R.id.do_search)

        doSearch.setOnClickListener {
            page = 1
            beerAdapter.search(campoBusca.text.toString(), page, pageSize)
        }

        val prev = findViewById<Button>(R.id.prev)

        prev.setOnClickListener {
            if (page > 1) page--
            beerAdapter.search(campoBusca.text.toString(), page, pageSize)
        }

        val next = findViewById<Button>(R.id.next)

        next.setOnClickListener {
            if (beerAdapter.beers.size == pageSize) page++
            beerAdapter.search(campoBusca.text.toString(), page, pageSize)
        }

        beerAdapter.search()
    }
}

