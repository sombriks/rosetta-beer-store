const roles = [
    { idrole: 1, descriptionrole: "manager" },
    { idrole: 2, descriptionrole: "customer" }
]

const user = {
    iduser: 1,
    loginuser: "admin@admin",
    hashuser: "e10adc3949ba59abbe56e057f20f883e" // 123456
}

const userrole = { iduser: 1, idrole: 1 }

const orderstatuses = [
    { idorderstatus: 1, descriptionorderstatus: "CART" },
    { idorderstatus: 2, descriptionorderstatus: "CHECKOUT" },
    { idorderstatus: 3, descriptionorderstatus: "DELIVERED" }
]

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
export const up = async (knex) => {
  await knex("role").insert(roles);
  await knex("user").insert(user);
  await knex("user_role").insert(userrole);
  return knex("orderstatus").insert(orderstatuses);
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
export const down = async (knex) => {
    await knex("orderstatus").del().whereIn("idorderstatus", orderstatuses.map(e => e.idorderstatus));
    await knex("user_role").del().where(userrole);
    await knex("user").del().where({ iduser: 1 });
    return knex("role").del().whereIn("idrole", roles.map(e => e.idrole));
};
