// beers 
const router = require("express").Router()
const Bookshelf = require("../components/config").Bookshelf
const knex = require("../components/config").knex
const commonRoutes = require("../components/common-routes")

const Beer = Bookshelf.Model.extend({
  idAttribute: "idbeer",
  tableName: "beer",
})

commonRoutes.apply(router, Beer, "idbeer")

module.exports = {
  router,
  Beer,
}
