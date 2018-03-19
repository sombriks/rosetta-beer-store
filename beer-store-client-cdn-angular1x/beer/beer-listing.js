// beer/beer-listing.js

class BeerListing {
  constructor(BeerService) {
    console.log("there!");
    this.BeerService = BeerService;
    this.query = { search: "", page: 1, pageSize: 10 };
    this.list();
  }

  list(p) {
    console.log(this.query);
    if (p) this.query.page += p;
    this.BeerService.list(this.query).then(ret => {
      this.beers = ret.data;
    });
  }
}

angular.module("beer-store").controller("BeerListing", BeerListing);
