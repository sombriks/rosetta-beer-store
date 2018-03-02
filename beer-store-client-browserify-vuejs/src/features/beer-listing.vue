<template>
  <md-layout md-gutter md-column>
    <topbar>
      <h1 slot="left">Beer Listing</h1>
    </topbar>
    <md-layout md-gutter md-column>
      <searchbar @onsearch="dosearch" :resultlist="beerlist"></searchbar>
      <beer-resume v-for="beer in beerlist" :key="beer.idbeer" :beer="beer"></beer-resume>
    </md-layout>
  </md-layout>
</template>

<script>
const { beerservice } = require("../components/restapi")
module.exports = {
  name: "BeerListing",
  data: _ => ({
    page: 1,
    beerlist: []
  }),
  created() {
    this.dosearch()
  },
  methods: {
    dosearch(s) {
      beerservice.list(s).then(ret => this.beerlist = ret.data)
    }
  }
}
</script>

