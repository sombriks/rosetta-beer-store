package beer.store

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        title = "Beer Store"
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }
}
