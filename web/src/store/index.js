import Vue from 'vue'
import Vuex from 'vuex'

import user from './modules/user'
import app from './modules/app'
import libs from "./modules/libs";
import campaigns from "./modules/campaigns"

Vue.use(Vuex)

export default new Vuex.Store({
    namespaced: true,
    modules: {
        user: user,
        app: app,
        libs: libs,
        campaigns: campaigns
    }
})