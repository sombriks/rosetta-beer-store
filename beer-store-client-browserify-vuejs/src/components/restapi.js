// 
const axios = require("axios")

const api = axios.create({
  headers: { "x-api-key": "my-api-key", "x-secret-key": "" },
  baseURL: "http://127.0.0.1:3000",
})

const beerservice = {
  list: (params, p = 1, s = 10) => api.get(`/beer/list/${p}/${s}`,{ params })
}

module.exports = {
  beerservice
}