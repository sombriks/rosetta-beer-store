
const beers = [
  { idbeer: 1, titlebeer: "Brahma", descriptionbeer: "A número 1!" },
  { idbeer: 2, titlebeer: "Skol", descriptionbeer: "A que desce redondo!" },
  { idbeer: 2, titlebeer: "Itaipava", descriptionbeer: "A cerveja do verão!" },
]

exports.up = knex =>
  knex("beer").insert(beers)

exports.down = (knex, Promise) =>
  Promise.each(beers, b => knex("beer").del().where(b))
