//
require("./main.css")

const Vue = require("vue")
const VueRouter = require("vue-router")
const VueMaterial = require("vue-material")

Vue.use(VueRouter)
Vue.use(VueMaterial)

Vue.component("beer-resume", require("./components/beer/beer-resume.vue"))

Vue.component("topbar", require("./components/shell/topbar.vue"))
Vue.component("searchbar", require("./components/shell/searchbar.vue"))
Vue.component("cart-resume", require("./components/shell/cart-resume.vue"))
Vue.component("user-resume", require("./components/shell/user-resume.vue"))

window.rootvm = new Vue({
  render: r => r(require("./components/mountpoint.vue")),
  el: document.getElementById("app"),
})