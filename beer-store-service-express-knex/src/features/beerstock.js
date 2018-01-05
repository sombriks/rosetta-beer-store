// beer stock
const router = require("express").Router()
const Bookshelf = require("../components/config").Bookshelf
const knex = require("../components/config").knex
const commonRoutes = require("../components/common-routes")

const BeerStock = Bookshelf.Model.extend({
  idAttribute: "idbeerstock",
  tableName: "beerstock",
  beer() {
    return this.belongsTo(require("./beer").Beer, "idbeer")
  }
})

commonRoutes.apply(router, BeerStock, ["beer"])

module.exports = {
  BeerStock,
  router,
}