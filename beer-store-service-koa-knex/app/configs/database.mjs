import Knex from "knex";
import cfg from "./knexfile.cjs"

export const database = Knex(cfg[process.env.NODE_ENV]);

export const doMigrate = () => database.migrate.latest(cfg[process.env.NODE_ENV]);
