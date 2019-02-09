<template>
  <md-card>
    <md-card-header v-if="beer">
      <md-avatar>
        <img :src="mediaservice.url(beer.idmedia)">
      </md-avatar>
      <md-card-header-text>
        <div class="md-title">{{beer.titlebeer}}</div>
        <div class="md-caption">{{beer.descriptionbeer}}</div>
      </md-card-header-text>
      <slot name="heading-options"></slot>
    </md-card-header>
  </md-card>
</template>

<script>
const { mediaservice } = require("../../api");
module.exports = {
  name: "BeerItem",
  props: ["beer"],
  data: _ => ({ mediaservice })
};
</script>

<style>
</style>
