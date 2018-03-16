// beer/beer-listing.js

class BeerListing {
  constructor(BeerService) {
    console.log("there!");
    this.BeerService = BeerService;
    BeerService.list().then(ret => {
      this.beers = ret.data;
      console.log(this.beers);
    });
  }
}

angular.module("beer-store").controller("BeerListing", BeerListing);
