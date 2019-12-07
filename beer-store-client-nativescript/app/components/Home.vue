<template>
  <Page>
    <ActionBar>
      <Label text="Beer Store"></Label>
    </ActionBar>
    <StackLayout>
      <Search :filter="filter" :results="results" @onSearch="doSearch" />
    </StackLayout>
  </Page>
</template>

<script>
import Search from "./Search";
import { mapState, mapActions } from "vuex";
export default {
  name: "home",
  components: {
    Search
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
    ...mapActions(["doSearch"])
  }
};
</script>

<style scoped lang="scss">
@import "~@nativescript/theme/scss/variables/blue";
</style>
