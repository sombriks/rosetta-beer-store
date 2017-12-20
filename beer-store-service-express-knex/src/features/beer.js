// beers and beer stock
const router = require("express").Router()
const Bookshelf = require("../components/config").Bookshelf
const knex = require("../components/config").knex

const Beer = Bookshelf.Model.extend({
  idAttribute: "idbeer",
  tableName: "beer",
})

const BeerStock = Bookshelf.Model.extend({
  idAttribute: "idbeerstock",
  tableName: "beerstock",
  beer() {
    return this.belongsTo(Beer, "idbeer")
  }
})

module.exports = {
  router,
  Beer,
  BeerStock
}
