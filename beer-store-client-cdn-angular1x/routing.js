// bootstrap.js

angular.module("beer-store").config($routeProvider => {
  $routeProvider.when("/beer-listing", {
    templateUrl: "beer/beer-listing.html",
    controller: "BeerListing",
    controllerAs: "ctl"
  });

  $routeProvider.when("/beer-details/:idbeer", {
    templateUrl: "beer/beer-details.html",
    controller: "BeerDetails",
    controllerAs: "ctl"
  });

  $routeProvider.otherwise({
    redirectTo: "/beer-listing"
  });
});

//https://docs.angularjs.org/api/ng/function/angular.bootstrap
angular.bootstrap(document.getElementById("app"), ["beer-store"]);
