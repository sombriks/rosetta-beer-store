import Vue from "nativescript-vue";
import Home from "./components/Home";
import { store } from "./store";

Vue.config.silent = false;

new Vue({
    template: `
        <Frame>
            <Home />
        </Frame>`,
    store,
    components: {
        Home
    }
}).$start();
