// app.js
angular
  .module("beer-store", ["ngMaterial", "ngRoute"])
//https://docs.angularjs.org/api/ng/function/angular.bootstrap
angular
  .bootstrap(document.getElementById("app"), ["beer-store"])
