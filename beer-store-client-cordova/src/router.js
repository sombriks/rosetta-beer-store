import Vue from "vue";
import Router from "vue-router";
import BeerListing from "./views/BeerListing.vue";

Vue.use(Router);

export default new Router({
  // mode: "history",
  base: process.env.BASE_URL,
  routes: [{
      path: "/",
      redirect: "/beer-listing"
    },
    {
      path: "/beer-listing",
      name: "beer-listing",
      component: BeerListing
    }
  ]
});
