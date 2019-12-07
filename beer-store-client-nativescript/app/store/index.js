import Vue from "nativescript-vue";
import Vuex from "vuex";
import * as http from "tns-core-modules/http";

Vue.use(Vuex);

const state = {
    results: []
};
const mutations = {
    setResults(state, results) {
        console.log(results);
        state.results = results || [];
    }
};
const actions = {
    doSearch({ commit, state }, { q, page, pageSize }) {
        const url = `http://192.168.0.111:3000/beer/list?search=${q}&page=${page}&pageSize=${pageSize}`;
        // TODO get baseUrl from env
        // TODO axios smells better
        http.getJSON(url).then(
            ret => commit("setResults", ret),
            err => console.log(err)
        );
    }
};

export const store = new Vuex.Store({ state, mutations, actions });
