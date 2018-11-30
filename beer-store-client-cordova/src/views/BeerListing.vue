<template>
  <div class="beer-listing">
    <topbar>
      <h1 slot="left">Beer Listing</h1>
    </topbar>
    <md-layout md-gutter md-column>
      <md-card v-for="beer in beers" :key="beer.idbeer">
        <md-card-header-text>
          <div class="md-title">{{beer.titlebeer}}</div>
          <div class="md-caption">{{beer.descriptionbeer}}</div>
        </md-card-header-text>
      </md-card>        
    </md-layout>
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
