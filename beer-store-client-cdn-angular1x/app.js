// app.js
angular
  .module("beer-store", ["ngMaterial", "ngRoute"])

const baseURL = "http://127.0.0.1:3000" // service endpoint

https://docs.angularjs.org/guide/module#registration-in-the-config-block
angular.module("beer-store").config($httpProvider => {
  $httpProvider.interceptors.push($q => ({
    request(config) {
      // the $http must be instructed to ignore baseurl for local resources
      if (!config.url.match(/\.js|\.html/))
        config.url = baseURL + config.url
      console.log(config.url)
      return config
    }
  }))
})