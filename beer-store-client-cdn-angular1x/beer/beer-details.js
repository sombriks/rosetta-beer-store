// beer/beer-details.js
class BeerDetails {
  constructor(BeerService, $routeParams) {
    this.BeerService = BeerService;
    BeerService.find($routeParams.idbeer).then(ret => {
      this.beer = ret.data;
    });
  }
}

angular.module("beer-store").controller("BeerDetails", BeerDetails);
