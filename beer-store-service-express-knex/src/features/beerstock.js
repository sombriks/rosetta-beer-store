// beer stock
const router = require("express").Router()
const { Bookshelf, knex } = require("../components/config")

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