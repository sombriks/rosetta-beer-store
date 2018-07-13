<template>
  <md-layout md-gutter="8">
    <md-layout md-align="start" md-gutter md-flex-xsmall="100" md-flex-small="70">
      <md-input-container class="ml">
        <label>Search</label>
        <md-input v-model="s" @input="onsearch"></md-input>
      </md-input-container>
    </md-layout>
    <md-layout md-align="end" md-gutter md-flex-xsmall="100" md-flex-small="30">
      <md-button class="md-raised md-primary" @click="onsearch(s,-1)" :disabled="p == 1">
        <md-icon>chevron_left</md-icon>
      </md-button>
      <md-button class="md-raised md-primary" @click="onsearch(s,1)" :disabled="!resultlist || resultlist.length < pageSize">
        <md-icon>chevron_right</md-icon>
      </md-button>
    </md-layout>
  </md-layout>
</template>
<script>
module.exports = {
  name: "SearchBar",
  props: {
    resultlist: {
      type: Array
    },
    pageSize: {
      type: Number,
      default: 10
    }
  },
  data() {
    return {
      s: "",
      p: 1
    };
  },
  methods: {
    onsearch(s, p) {
      this.s = s;
      if (!p) this.p = 1;
      else {
        this.p += p;
        if (this.resultlist && this.resultlist.length < this.ps) this.p -= p;
        if (this.p < 1) this.p = 1;
      }
      this.$emit("onsearch", {
        pageSize: this.pageSize,
        page: this.p,
        search: s,
      });
    }
  }
};
</script>

<style scoped>
.ml {
  margin-left: 1em;
}
</style>
