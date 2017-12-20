//
let environment = process.env.NODE_ENV || "development"

const knex = require("knex")(require("../../knexfile")[environment])

const Bookshelf = require("bookshelf")(knex)

module.exports = {
  environment,
  Bookshelf,
  knex
}
