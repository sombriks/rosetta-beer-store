//
const express = require("express")
const bodyParser = require("body-parser")
const morgan = require("morgan")
const cors = require("cors")

const knex = require("./components/config").knex

const app = express()

app.use(morgan("common", "immediate"))

app.use(bodyParser.json({ limit: 1024 * 1024 }))
const type = ['application/octet-stream', 'image/*', 'application/pdf', 'audio/*']
const limit = 10 * 1024 * 1024
app.use(bodyParser.raw({ type, limit }))

app.use(cors())

app.use("/beer", require("./features/beer").router)
app.use("/beerstock", require("./features/beerstock").router)
app.use("/media", require("./features/media").router)
app.use("/order", require("./features/order").router)
app.use("/user", require("./features/user").router)

exports.app = app

exports.init = _ => {
  knex.migrate.latest().then(_ => {
    app.listen(3000)
    console.log("Beer store ready")
  })
}
