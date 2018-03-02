// 
// that CRUD pattern that always repeat itself.
//
// const router = require("express").Router()
// const { Bookshelf, knex } = require("../components/config")
// const commonRoutes = require("../components/common-routes")
// ...
// commonRoutes.apply(routes, NewModel)
// 
const errfn = require("./config").errfn

exports.apply = (router, BsModel, withRelated, searchClause) => {

  const idAttribute = BsModel.prototype.idAttribute;

  const listPage = query => {

    let page = query.page || 1
    delete query.page
    let pageSize = query.pageSize || 10
    delete query.pageSize

    return BsModel.where(qb => {
      if (searchClause)
        searchClause(qb, query)
      qb.where(query)
    }).fetchPage({ page, pageSize, withRelated })
  }

  const find = id => BsModel.query("where", idAttribute, id).fetch({ withRelated })

  const insert = data => new BsModel(data).save()

  const update = data => {
    let up = {}
    up[idAttribute] = data[idAttribute]
    return new BsModel(up).save(data)
  }

  const del = id => BsModel.query("where", idAttribute, id).destroy()

  router.get("/list", (req, res) =>
    listPage(req.query).then(ret => res.send(ret)).catch(errfn(res)))

  router.get("/:id", (req, res) =>
    find(req.params.id).then(ret => res.send(ret)).catch(errfn(res)))

  router.post("/save", (req, res) =>
    insert(req.body).then(ret => res.send(ret)).catch(errfn(res)))

  router.put("/save", (req, res) =>
    update(req.body).then(ret => res.send(ret)).catch(errfn(res)))

  router["delete"]("/:id", (req, res) => 
    del(req.params.id).then(ret => res.send(ret)).catch(errfn(res)))

}