<template>
  <div class="beer-listing">
    <mu-appbar class="appbar" title="Beer Listing" color="indigo500" z-depth="0"></mu-appbar>
    <mu-container>
      <mu-form :model="form">
        <mu-flex justify-content="start">
          <mu-form-item class="search-bar" prop="input" label="Search">
            <mu-text-field v-model="form.input" @input="onSearch(form.input, undefined)"></mu-text-field>
          </mu-form-item>
        </mu-flex>
        <mu-flex justify-content="end">
          <div class="left-button">
            <mu-button @click="onSearch(form.input, -1)" :disabled="page === 1" color="indigo500">&#60;</mu-button>
          </div>
          <mu-button @click="onSearch(form.input, 1)" color="indigo500" :disabled="!beers || beers.length < pageSize">&#62;</mu-button>
        </mu-flex>
      </mu-form>
    </mu-container>
    <mu-list textline="two-line">
      <mu-list-item v-for="beer in beers" :key="beer.idbeer" :avatar="true">
        <mu-avatar class="icon">
          <img :src="mediaservice.url(beer.idmedia)">
        </mu-avatar>
        <mu-list-item-content>
          <mu-list-item-title>{{beer.titlebeer}}</mu-list-item-title>
          <mu-list-item-sub-title>{{beer.descriptionbeer}}</mu-list-item-sub-title>
        </mu-list-item-content>
      </mu-list-item>
    </mu-list>
  </div>
</template>

<script>
import { beerservice, mediaservice } from "@/restapi.js";

export default {
  name: "beer-listing",
  components: {},
  created() {
    this.listAllBeers();
  },
  data: () => ({
    beers: [],
    form: {
      input: ""
    },
    mediaservice,
    page: 1,
    pageSize: 10
  }),
  methods: {
    listAllBeers() {
      beerservice.list().then(ret => (this.beers = ret.data));
    },
    onSearch(s, p) {
      this.form.input = s;
      if (!p) this.page = 1;
      else {
        this.page += p;
        if (this.beers && this.beers.length < this.ps) this.page -= p;
        if (this.page < 1) this.page = 1;
      }
      const params = {
        pageSize: this.pageSize,
        page: this.page,
        search: s
      };
      beerservice.list(params).then(ret => (this.beers = ret.data));
    }
  }
};
</script>

<style>
.appbar {
  width: 100%;
  text-align: left;
}
.search-bar {
  width: 100%;
}
.left-button {
  margin-right: 10px;
}
.icon {
  margin-right: 15px;
}
</style>
