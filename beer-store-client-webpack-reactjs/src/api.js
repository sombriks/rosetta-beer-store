import axios from "axios";

const env = process.env.NODE_ENV || "development";

console.log(`we are on [${env}] environment`);

const addr = {
  production: "https://rosetta-beer-store.io",
  development: "http://127.0.0.1:3000",
};

const api = axios.create({
  headers: {"x-api-key": "my-api-key", "x-secret-key": ""},
  baseURL: addr[env],
});

export const beerservice = {
  list: params => api.get("/beer/list", {params}),
  find: id => api.get(`/beer/${id}`),
};

export const mediaservice = {
  url: id => (id ? `${addr[env]}/media/${id}` : `${addr[env]}/icon.svg`),
};
