import Vue from "nativescript-vue";
import BeerList from "./components/BeerList";
import { store } from "./store";

Vue.config.silent = false;

new Vue({
    template: `
        <Frame>
            <BeerList/>
        </Frame>`,
    store,
    components: {
        BeerList
    }
}).$start();
