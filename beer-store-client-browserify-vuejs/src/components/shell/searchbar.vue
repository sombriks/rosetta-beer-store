<template>
  <md-layout md-gutter class="mh">
   <md-layout md-align="start" md-gutter class="ml" md-flex="70">
     <md-input-container>
       <label>Search</label>
       <md-input v-model="s" @input="onsearch"></md-input>
     </md-input-container>
   </md-layout>
   <md-layout md-align="end" md-gutter md-flex="25">
     <md-button class="md-raised md-primary" @click="onsearch(s,-1)" :disabled="p == 1">
       <md-icon>chevron_left</md-icon>
     </md-button>
     <md-button class="md-raised md-primary" @click="onsearch(s,1)" :disabled="!resultlist || resultlist.length < ps">
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
    page: {
      type: Number,
      default: 1
    },
    pagesize: {
      type: Number,
      default: 10
    },
    search: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      s: "",
      ps: 10,
      p: 1
    };
  },
  created() {
    this.ps = this.pagesize;
    this.s = this.search;
    this.p = this.page;
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
      const page = this.p;
      const search = this.s;
      const pagesize = this.ps;
      this.$emit("onsearch", {
        pagesize,
        search,
        page
      });
    }
  }
};
</script>

<style scoped>
.mh {
  max-height: 4em;
}
.ml {
  margin-left: 1em;
}
</style>
