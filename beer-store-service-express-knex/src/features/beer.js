// beers 
const router = require("express").Router()
const { Bookshelf, knex } = require("../components/config")

const commonRoutes = require("../components/common-routes")

const Beer = Bookshelf.Model.extend({
  idAttribute: "idbeer",
  tableName: "beer",
})

commonRoutes.apply(router, Beer, [], (qb, q) => {
  if (q.search) {
    let s = q.search
    qb.where("titlebeer", "like", `%${s}%`)
  }
  delete q.search
})

module.exports = {
  router,
  Beer,
}
