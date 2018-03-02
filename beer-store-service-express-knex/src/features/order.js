// beer orders, item orders
const router = require("express").Router()
const { Bookshelf, knex } = require("../components/config")

const commonRoutes = require("../components/common-routes")

const Order = Bookshelf.Model.extend({
  tableName: "order",
  idAttribute: "idorder"
})

commonRoutes.apply(router, Order)

module.exports = {
  router,
  Order,
}
