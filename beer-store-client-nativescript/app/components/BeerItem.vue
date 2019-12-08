<template>
  <WrapLayout v-if="beer">
    <Label>{{ beer.titlebeer }}</Label>
  </WrapLayout>
</template>

<script>
export default {
  name: "beer-item",
  props: ["beer"]
};
</script>

<style></style>
