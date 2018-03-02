module.exports = [
  { path: "/beer-listing", component: require("../features/beer-listing.vue"), alias:"/" },
  { path: "/beer-details/:idbeer", component: require("../features/beer-details.vue") },
  
]