/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
export const up = async (knex) => {
    return knex.schema
        .createTable("media", tb => {
            tb.increments("idmedia")
            tb.timestamp("creationdatemedia").notNullable().defaultTo(knex.fn.now())
            tb.binary("datamedia").notNullable()
            tb.string("nomemedia").notNullable()
            tb.string("mimemedia").notNullable()
        })
        .createTable("user", tb => {
            tb.increments("iduser")
            tb.timestamp("creationdateuser").notNullable().defaultTo(knex.fn.now())
            tb.string("loginuser").notNullable().unique()
            tb.string("hashuser").notNullable()
            tb.integer("idmedia").references("media.idmedia")
        })
        .createTable("role", tb => {
            tb.increments("idrole")
            tb.string("descriptionrole").notNullable()
        })
        .createTable("user_role", tb => {
            tb.integer("iduser").notNullable().references("user.iduser").onDelete("cascade")
            tb.integer("idrole").notNullable().references("role.idrole").onDelete("cascade")
            tb.primary(["iduser", "idrole"])
        })
        .createTable("beer", tb => {
            tb.increments("idbeer")
            tb.timestamp("creationdatebeer").notNullable().defaultTo(knex.fn.now())
            tb.string("titlebeer")
            tb.text("descriptionbeer")
            tb.integer("idmedia").references("media.idmedia")
        })
        .createTable("beerstock", tb => {
            tb.increments("idbeerstock")
            tb.timestamp("creationdatebeerstock").notNullable().defaultTo(knex.fn.now())
            tb.integer("idbeer").unique().notNullable().references("beer.idbeer").onDelete("cascade")
            tb.integer("countstock").notNullable().defaultTo(0)
        })
        .createTable("orderstatus", tb => {
            tb.increments("idorderstatus")
            tb.string("descriptionorderstatus")
        })
        .createTable("order", tb => {
            tb.increments("idorder")
            tb.timestamp("creationdateorder").notNullable().defaultTo(knex.fn.now())
            tb.integer("idorderstatus").notNullable().references("orderstatus.idorderstatus").defaultTo(1)
            tb.integer("iduser").notNullable().references("user.iduser").onDelete("cascade")
            tb.timestamp("lastupdateorder").notNullable().defaultTo(knex.fn.now())
        })
        .createTable("itemorder", tb => {
            tb.increments("iditemorder")
            tb.integer("idorder").notNullable().references("order.idorder").onDelete("cascade")
            tb.integer("idbeer").unique().notNullable().references("beer.idbeer").onDelete("cascade")
            tb.integer("amountitemorder").notNullable().defaultTo(1)
        })
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
export const down = async (knex) => {
    return knex.schema
        .dropTable("itemorder")
        .dropTable("order")
        .dropTable("orderstatus")
        .dropTable("beerstock")
        .dropTable("beer")
        .dropTable("user_role")
        .dropTable("role")
        .dropTable("user")
        .dropTable("media")
};
