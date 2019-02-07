<template>
  <div>
    <topbar>
      <h1 slot="left">Beer Details</h1>
    </topbar>
    <md-button @click="$router.go(-1)" class="md-primary md-raised">
      <md-icon>arrow_back</md-icon>
    </md-button>
    <beer-item v-if="beer" :beer="beer"></beer-item>
  </div>
</template>

<script>
const { beerservice } = require("../restapi")
module.exports = {
  name: "BeerDetails",
  data: _ => ({ beer: null }),
  created() {
    beerservice
      .find(this.$route.params.idbeer)
      .then(ret => this.beer = ret.data)
  }
}
</script>

<style>
</style>
