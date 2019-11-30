<template>
  <Page>
    <ActionBar>
      <Label text="Beer Store"></Label>
    </ActionBar>
    <StackLayout>
      <!-- https://nativescript-vue.org/en/docs/elements/layouts/stack-layout/ -->
      <TextField hint="Search beer" v-model="textFieldValue" />
      <Button text="Search" @tap="doSearch" />
    </StackLayout>
  </Page>
</template>

<script>
export default {
  name: "home",
  data() {
    return {
      textFieldValue: ""
    };
  },
  methods: {
    doSearch() {
      console.log(this.textFieldValue);
    }
  }
};
</script>

<style scoped lang="scss">
@import "~@nativescript/theme/scss/variables/blue";

// Custom styles
.fas {
  @include colorize($color: accent);
}

.info {
  font-size: 20;
  horizontal-align: center;
  vertical-align: center;
}
</style>
