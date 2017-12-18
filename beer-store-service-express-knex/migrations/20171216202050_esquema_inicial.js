
exports.up = knex => knex.schema
  .createTable("media", tb => {
    tb.increments("idmedia")
    tb.timestamp("creationdatemedia").notNullable().defaultTo(knex.fn.now())
    tb.binary("datamedia").notNullable()
    tb.string("nomemedia").notNullable()
  })
  .createTable("beer", tb => {
    tb.increments("idberr")
    tb.timestamp("creationdatebeer").notNullable().defaultTo(knex.fn.now())
    tb.string("titlebeer")
    tb.text("descriptionbeer")
    tb.integer("idmedia").references("media.idmedia")
  })
  .createTable("user", tb => {
    tb.increments("iduser")
    tb.timestamp("creationdateuser").notNullable().defaultTo(knex.fn.now())
    tb.string("loginuser").notNullable().unique()
    tb.string("hashuser").notNullable()
    tb.integer("idmedia").references("media.idmedia")
  })
  .createTable("beerstock", tb => {
    tb.increments("idbeerstock")
    tb.timestamp("creationdatebeerstock").notNullable().defaultTo(knex.fn.now())
    tb.integer("idbeer").unique().notNullable().references("beer.idbeer").onDelete("cascade")
    tb.integer("idstock").notNullable().defaultTo(0)
  })

exports.down = knex.schema.
  dropTable("beerstock")
  .dropTable("user")
  .dropTable("beer")
  .dropTable("media")

