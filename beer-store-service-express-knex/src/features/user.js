// system users and roles
const router = require("express").Router()
const Bookshelf = require("../components/config").Bookshelf
const knex = require("../components/config").knex
const commonRoutes = require("../components/common-routes")
const errfn = require("../components/config").errfn

const Role = Bookshelf.Model.extend({
  idAttribute: "idrole",
  tableName: "role"
})

const User = Bookshelf.Model.extend({
  idAttribute: "iduser",
  tableName: "user",
  roles() {
    return this.belongsToMany(Role, "user_role", "idrole", "iduser")
  }
})

router.get("/roles", (req, res) =>
  Role.fetchAll().then(ret => res.send(ret)).catch(errfn(res)))

router.post("/addrole", (req, res) =>
  knex("user_role").insert(req.body).then(ret => res.send(ret)).catch(errfn(res)))

router.post("/removerole", (req, res) =>
  knex("user_role").del().where(req.body).then(ret => res.send("OK")).catch(errfn(res)))

commonRoutes.apply(router, User, "iduser", ["roles"])

module.exports = {
  router,
  User,
  Role,
}
