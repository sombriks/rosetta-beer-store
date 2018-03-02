// multimedia resources
const router = require("express").Router()
const { Bookshelf, knex } = require("../components/config")

const commonRoutes = require("../components/common-routes")

const Media = Bookshelf.Model.extend({
  idAttribute: "idmedia",
  tableName: "media"
})

commonRoutes.apply(router, Media)

module.exports = {
  router,
  Media,
}
