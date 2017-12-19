
const roles = [
  { idrole: 1, descriptionrole: "manager" },
  { idrole: 2, descriptionrole: "customer" }
]

const user = {
  iduser: 1,
  loginuser: "admin@admin",

}

const orderstatuses = [
  { idorderstatus: 1, descriptionorderstatus: "CART" },
  { idorderstatus: 2, descriptionorderstatus: "CHECKOUT" },
  { idorderstatus: 3, descriptionorderstatus: "DELIVERED" },

]

exports.up = knex => knex("role").insert(roles).then(_ => {
  return knex("user").insert(user)
}).then(_ => {
  return knex("orderstatus").insert(orderstatuses)
})

exports.down = knex =>
  knex("orderstatus").del()
    .whereIn("idorderstatus", orderstatuses.map(e => e.idorderstatus)).then(_ => {
      return knex("user").del().where({ iduser: 1 })
    }).then(_ => {
      return knex("role").del().whereIn("idrole", roles.map(e => e.idrole))
    })