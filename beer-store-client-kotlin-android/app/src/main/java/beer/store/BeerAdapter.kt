package beer.store

import android.util.Log
import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView.Adapter

class BeerAdapter : Adapter<BeerViewHolder>() {

    val beers = mutableListOf<Beer>()

    init {
        for (i in 1..10) beers.add(Beer(i))
    }

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
}