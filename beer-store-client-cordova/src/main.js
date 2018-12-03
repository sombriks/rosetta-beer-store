import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import VueMaterial from "vue-material";
import 'vue-material/dist/vue-material.min.css'

Vue.config.productionTip = false;

Vue.use(VueMaterial)

const init = _ => {
  new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount("#app");
}

if (window.cordova) document.addEventListener("deviceready", init)
else init()