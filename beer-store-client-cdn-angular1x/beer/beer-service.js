// beer/beer-service.js
class BeerService {

  constructor($http) {
    this.$http = $http
    console.log("a-ok!")
  }

  list(params) {
    return this.$http.get("/beer/list", { params })
  }

  find(idbeer) {
    return this.$http.get(`/beer/${idbeer}`)
  }
}

angular // yes yes those people used to love CamelCase
  .module("beer-store").service("BeerService", BeerService)