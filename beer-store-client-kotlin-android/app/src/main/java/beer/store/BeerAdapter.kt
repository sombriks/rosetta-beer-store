package beer.store

import android.util.Log
import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView.Adapter

class BeerAdapter : Adapter<BeerViewHolder>() {

    override fun getItemCount(): Int {
        Log.i("beer.store","getItemCount")
        return 10
    }

    override fun onBindViewHolder(holder: BeerViewHolder, position: Int) {
        Log.i("beer.store","onBindViewHolder")
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): BeerViewHolder {
        Log.i("beer.store","onCreateViewHolder")
        var view = LayoutInflater.from(parent.context).inflate(R.layout.beer_item,parent,false)
        return BeerViewHolder(view)
    }
}