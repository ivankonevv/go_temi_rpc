import Vuex from 'vuex'
import authModule from "./modules/auth"
import Vue from "vue";

Vue.use(Vuex)


const store = new Vuex.Store ({
    modules: {
        auth: authModule
    }
});

export default store;