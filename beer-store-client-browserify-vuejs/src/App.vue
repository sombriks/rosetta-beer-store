<template>
  <router-view></router-view>  
</template>

<script>
const VueRouter = require("vue-router");
const router = new VueRouter({
  routes: [...require("./routes"), { path: "/", redirect: "/beer-listing" }]
});
module.exports = {
  name: "MountPoint",
  router
};
</script>
