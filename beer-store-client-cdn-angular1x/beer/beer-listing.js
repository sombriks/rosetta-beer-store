// beer/beer-listing.js

class BeerListing {

  constructor(BeerService) {
    console.log("there!")
    this.BeerService = BeerService
    this.beers = BeerService.list()
  }
}

angular.module("beer-store").controller("BeerListing", BeerListing)
