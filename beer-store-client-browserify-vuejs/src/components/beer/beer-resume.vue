<template>
  <md-card>
    <md-card-header>
      <md-avatar>
        <img :src="mediaservice.url(beer.idmedia)"/>
      </md-avatar>
    </md-card-header>
  </md-card>
</template>

<script>
const { mediaservice } = require("../restapi")
module.exports = {
  name: "BeerResume",
  props: ["beer"],
  data: _ => ({ mediaservice })
}
</script>

<style>

</style>
