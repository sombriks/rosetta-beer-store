<template>
  <md-toolbar>
    <md-layout md-gutter>
      <md-layout md-gutter md-align="start">
        <slot name="left"></slot>
      </md-layout>
      <md-layout md-gutter md-align="end">
        <slot name="right"></slot>
        <cart-resume></cart-resume>
        <user-resume></user-resume>
      </md-layout>
    </md-layout>
  </md-toolbar>
</template>

<script>
module.exports = {
  name: "Topbar"
};
</script>

<style>
</style>
