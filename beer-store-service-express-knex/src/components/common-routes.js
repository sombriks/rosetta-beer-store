
const errfn = require("./config").errfn

exports.apply = (router, BsModel, idAttribute, withRelated) => {

  const listPage = (query, page = 1, pageSize = 10) =>
    BsModel.where(query).fetchPage({ page, pageSize, withRelated })

  const find = id => BsModel.query("where", idAttribute, id).fetch({ withRelated })

  const insert = data => new BsModel(data).save()

  const update = data => {
    let up = {}
    up[idAttribute] = data[idAttribute]
    return new BsModel(up).save(data)
  }

  const del = id => BsModel.query("where", idAttribute, id).destroy()

  router.get("/list(/:page)?(/:pageSize)?", (req, res) =>
    listPage(req.query, req.params.page, req.params.pageSize).then(ret => res.send(ret)).catch(errfn(res)))

  router.get("/:id", (req, res) =>
    find(req.params.id).then(ret => res.send(ret)).catch(errfn(res)))

  router.post("/save", (req, res) =>
    insert(req.body).then(ret => res.send(ret)).catch(errfn(res)))

  router.put("/save", (req, res) =>
    update(req.body).then(ret => res.send(ret)).catch(errfn(res)))

  router["delete"]("/:id", (req, res) => del(req.params.id).then(ret => res.send(ret)).catch(errfn(res)))

}