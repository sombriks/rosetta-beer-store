//
const environment = process.env.NODE_ENV || "development"

console.log(`we are on [${environment}] mode!`)

const knex = require("knex")(require("../../knexfile")[environment])

const Bookshelf = require("bookshelf")(knex)

Bookshelf.plugin('pagination')

const errfn = res => err => {
  res.status(500).send(err)
  console.log(err)
}

module.exports = {
  environment,
  Bookshelf,
  errfn,
  knex
}
