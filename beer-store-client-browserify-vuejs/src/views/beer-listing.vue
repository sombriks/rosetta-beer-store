<template>
  <div>
    <topbar>
      <h1 slot="left">Beer Listing</h1>
    </topbar>
    <md-layout md-gutter md-column>
      <searchbar @onsearch="dosearch" :resultlist="beerlist"></searchbar>
      <beer-item v-for="beer in beerlist" :key="beer.idbeer" :beer="beer">
        <md-button
          slot="heading-options"
          class="md-icon-button"
          @click="$router.push(`/beer-details/${beer.idbeer}`)"
        >
          <md-icon>visibility</md-icon>
        </md-button>
      </beer-item>
    </md-layout>
  </div>
</template>

<script>
const { beerservice } = require("../restapi");
module.exports = {
  name: "BeerListing",
  data: _ => ({
    page: 1,
    beerlist: []
  }),
  created() {
    this.dosearch();
  },
  methods: {
    async dosearch(s) {
      const ret = await beerservice.list(s);
      this.beerlist = ret.data;
    }
  }
};
</script>

