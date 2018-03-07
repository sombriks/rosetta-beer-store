
const beers = [
  { idbeer: 1, titlebeer: "Brahma", descriptionbeer: "A número 1!" },
  { idbeer: 2, titlebeer: "Antartica Original", descriptionbeer: "Pilsen" },
  { idbeer: 3, titlebeer: "Itaipava", descriptionbeer: "A cerveja do verão!" },
  { idbeer: 4, titlebeer: "Devassa", descriptionbeer: "Tropical Lager" },
  { idbeer: 5, titlebeer: "Corona", descriptionbeer: "Extra" },
  { idbeer: 6, titlebeer: "Therezópolis", descriptionbeer: "Cerveja especial" },
  { idbeer: 7, titlebeer: "Budweiser", descriptionbeer: "King of Beers" },
  { idbeer: 8, titlebeer: "Heineken", descriptionbeer: "Premium quality" },
  { idbeer: 9, titlebeer: "Skol", descriptionbeer: "A que desce redondo!" },
  { idbeer: 10, titlebeer: "Kaiser", descriptionbeer: "Cerveja bem cervejada" },
  { idbeer: 11, titlebeer: "Eisenbahn", descriptionbeer: "Cerveja puro malte" },
  { idbeer: 12, titlebeer: "Liefmans", descriptionbeer: "Fruitesse" },
  { idbeer: 13, titlebeer: "Bohemia", descriptionbeer: "Cerveja pilsen" },
]

exports.up = knex =>
  knex("beer").insert(beers)

exports.down = (knex, Promise) =>
  Promise.each(beers, b => knex("beer").del().where(b))
