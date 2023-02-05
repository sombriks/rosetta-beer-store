import Router from "@koa/router"

import {database} from "../configs/database.mjs";

export const listBeers = (search, page, pageSize) =>
    database("beer").select()
        .whereLike("titlebeer", `%${search}%`)
        .orWhereLike("descriptionbeer", `%${search}%`)
        .offset((page - 1) * pageSize)
        .limit(pageSize)

export const findBeer = (idbeer) => database("beer")
    .select().where({idbeer}).first();

export const insertBeer = (beer) => database("beer").insert(beer);
export const updateBeer = (idbeer, beer) =>
    database("beer").update(beer).where({idbeer});

export const deleteBeer = (idbeer) =>
    database("beer").del().where({idbeer});

export const beerRouter = new Router();

beerRouter.get("/beers", async ctx =>
    ctx.body = await listBeers(ctx.query.search || "", ctx.query.page || 1, ctx.query.pageSize || 10));
beerRouter.get("/beer/list", async ctx =>
    ctx.body = await listBeers(ctx.query.search || "", ctx.query.page || 1, ctx.query.pageSize || 10));

beerRouter.get("/beers/:id", async ctx =>
    ctx.body = await findBeer(ctx.params.id) || ctx.throw(404, "NOT_FOUND"));

beerRouter.get("/beer/:id", async ctx =>
    ctx.body = await findBeer(ctx.params.id) || ctx.throw(404, "NOT_FOUND"));

beerRouter.post("/beers", async ctx =>
    ctx.body = await insertBeer(ctx.request.body))

beerRouter.put("/beers/:id", async ctx =>
    ctx.body = await updateBeer(ctx.params.id, ctx.request.body));

beerRouter.del("/beers/:id", async ctx =>
    ctx.body = await deleteBeer(ctx.params.id));