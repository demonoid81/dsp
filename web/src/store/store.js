import Vuex from 'vuex';
import Vue from "vue";


Vue.use(Vuex);
const store = new Vuex.Store({
    state: {
        ssp: [],
        dsp: []
    }
})

export default store