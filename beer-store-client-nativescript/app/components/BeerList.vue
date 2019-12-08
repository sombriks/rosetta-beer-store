<template>
  <Page>
    <ActionBar>
      <Label text="Beer Store"></Label>
    </ActionBar>
    <StackLayout>
      <Search :filter="filter" :results="results" @onSearch="doSearch" />
      <ListView height="*" for="item in results" @itemTap="doDetail">
        <v-template>
          <BeerItem :beer="item" />
        </v-template>
      </ListView>
    </StackLayout>
  </Page>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Search from "./Search";
import BeerItem from "./BeerItem";
export default {
  name: "beer-list",
  components: {
    Search,
    BeerItem
  },
  data() {
    return {
      filter: {
        q: "",
        page: 1,
        pageSize: 10
      }
    };
  },
  computed: {
    ...mapState(["results"])
  },
  methods: {
    ...mapActions(["doSearch"]),
    doDetail(event) {
      console.log(event.item);
    }
  }
};
</script>

<style scoped lang="scss">
@import "~@nativescript/theme/scss/variables/blue";
</style>
