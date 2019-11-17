package beer.store

import android.view.View
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView

class BeerViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {

    val beerLabel = itemView.findViewById<TextView>(R.id.beer_label)

}