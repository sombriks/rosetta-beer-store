package beer.store

import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView.Adapter
import kotlinx.coroutines.*


class BeerAdapter : Adapter<BeerViewHolder>() {

    val beers = mutableListOf<Beer>()

    override fun getItemCount(): Int {
        return beers.size
    }

    override fun onBindViewHolder(holder: BeerViewHolder, position: Int) {
        val b = beers[position]
        holder.beerLabel.text = b.titleBeer + " " + b.idBeer
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): BeerViewHolder {
        var view = LayoutInflater.from(parent.context).inflate(R.layout.beer_item, parent, false)
        return BeerViewHolder(view)
    }

    fun search(search: String = "", page: Int = 1, pageSize: Int = 10) {
        GlobalScope.launch(Dispatchers.Main) {
            beers.removeAll(beers)
            beers.addAll(BeerService.instance().list(search,page,pageSize))
            notifyDataSetChanged()
        }
    }
}