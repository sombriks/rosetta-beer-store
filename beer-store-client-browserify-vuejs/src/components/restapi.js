// 
const axios = require("axios")

const api = axios.create({
  headers: { "x-api-key": "my-api-key", "x-secret-key": "" },
  baseURL: "http://127.0.0.1:3000",
})

