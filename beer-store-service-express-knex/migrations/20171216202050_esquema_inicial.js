
exports.up = knex => knex.schema
.createTable("media",tb => {
  tb.increments("idmedia")
  tb.timestamp("dtcriacaomedia").notNullable().defaultTo(knex.fn.now())
  tb.binary("datamedia").notNullable()
  tb.string("nomemedia").notNullable()
})
.createTable("beer", tb => {
  tb.increments("idberr")
  tb.timestamp("dtcriacaobeer").notNullable().defaultTo(knex.fn.now())
  tb.string("titlebeer")
  tb.text("descriptionbeer")
  tb.integer("idmedia")
})

exports.down = knex.schema.
  dropTable("beer")
  dropTable("media")
  
