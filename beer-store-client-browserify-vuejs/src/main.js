require("./main.css");
console.log(`this is beer store client [${require("../package.json").version}] version`);

const Vue = require("vue");
const VueRouter = require("vue-router");
const VueMaterial = require("vue-material");

Vue.use(VueRouter);
Vue.use(VueMaterial);

Vue.component("beer-item", require("./components/beer/beer-item.vue"));

Vue.component("topbar", require("./components/shell/topbar.vue"));
Vue.component("searchbar", require("./components/shell/searchbar.vue"));
Vue.component("cart-resume", require("./components/shell/cart-resume.vue"));
Vue.component("user-resume", require("./components/shell/user-resume.vue"));

window.rootvm = new Vue({
  render: r => r(require("./App.vue")),
  el: document.getElementById("app"),
});
