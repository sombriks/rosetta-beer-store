import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import MuseUI from 'muse-ui';
import 'muse-ui/dist/muse-ui.css';

Vue.config.productionTip = false;

Vue.use(MuseUI)

const init = _ => {
  new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount("#app");
}

if (window.cordova) document.addEventListener("deviceready", init)
else init()