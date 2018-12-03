// 
const axios = require("axios")

const env = process.env.NODE_ENV || "development"

console.log("we are on [%s] environment")

const addr = {
  production: "http://192.168.0.4:3000",
  development: "http://127.0.0.1:3000"
}

const api = axios.create({
  headers: { "x-api-key": "my-api-key", "x-secret-key": "" },
  baseURL: addr[env],
})

const beerservice = {
  list: params => api.get("/beer/list", { params }),
  find: id => api.get(`/beer/${id}`)
}

const mediaservice = {
  url: id => id ? `${addr[env]}/media/${id}` : `${addr[env]}/icon.svg`,
}

module.exports = {
  beerservice,
  mediaservice
}