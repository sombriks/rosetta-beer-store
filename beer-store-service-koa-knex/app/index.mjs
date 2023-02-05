import Koa from "koa";
import cors from "@koa/cors"
import Router from "@koa/router";
import bodyParser from "koa-bodyparser";

import {beerRouter} from "./routes/beers.mjs";

export const app = new Koa();

app.use(cors());
app.use(bodyParser());

app.use(beerRouter.routes()).use(beerRouter.allowedMethods());

const status = new Router();
status.get("/status", async ctx => ctx.body = "online")
app.use(status.routes()).use(status.allowedMethods());
