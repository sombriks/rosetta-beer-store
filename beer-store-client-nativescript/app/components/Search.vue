<template>
  <WrapLayout>
    <!-- https://nativescript-vue.org/en/docs/elements/layouts/stack-layout/ -->
    <TextField width="60%" hint="Search beer" v-model="filter.q" />
    <Button width="auto" text="Search" @tap="$emit('onSearch', filter)" />
    <DockLayout width="100%">
      <Button width="20%" dock="left" text="Prev" @tap="doPrev(filter)" />
      <Button width="20%" dock="right" text="Next" @tap="doNext(filter)" />
    </DockLayout>
  </WrapLayout>
</template>

<script>
export default {
  name: "search",
  props: ["filter", "results"],
  methods: {
    doPrev(filter) {
      if (filter.page < 2) return;
      filter.page--;
      this.$emit("onSearch", filter);
    },
    doNext(filter) {
      if (this.results.length == filter.pageSize) return;
      filter.page++;
      this.$emit("onSearch", filter);
    }
  }
};
</script>

<style></style>
