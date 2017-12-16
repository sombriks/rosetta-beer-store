<template>
  <router-view></router-view>  
</template>

<script>
const VueRouter = require("vue-router");
const router = new VueRouter({
  routes: require("./routes")
});
module.exports = {
  name: "MountPoint",
  router
};
</script>
