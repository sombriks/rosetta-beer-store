<template>
  <div class="beer-listing">
    <h1 slot="left">Beer Listing</h1>
    <ul v-for="beer in beers" :key="beer.idbeer">
      <li class="md-title">{{beer.titlebeer}}</li>
      <li class="md-caption">{{beer.descriptionbeer}}</li>
    </ul>
  </div>
</template>

<script>
import { beerservice } from "@/restapi.js";

export default {
  name: "beer-listing",
  components: {},
  created() {
    this.listAllBeers();
  },
  data: _ => ({
    beers: []
  }),
  methods: {
    listAllBeers() {
      beerservice.list().then(ret => (this.beers = ret.data));
    }
  }
};
</script>
