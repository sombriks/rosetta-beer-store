// beers and beer stock
const router = require("express").Router()
const Bookshelf = require("../components/config").Bookshelf
const knex = require("../components/config").knex
const commonRoutes = require("../components/common-routes")

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

commonRoutes.apply(router, Beer)

module.exports = {
  router,
  Beer,
  BeerStock
}
