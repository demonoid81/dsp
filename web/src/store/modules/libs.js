import axios from "axios";


const state = {
    countries: [],
    os: [],
    browsers:[]
}

const getters = {
    countries: state => state.countries,
    os: state => state.os,
    browsers: state => state.browsers
}

const mutations = {
    setCountries (state, countries) {
        state.countries = countries
    },
    setOS (state, os) {
        state.countries = countries
    },
    setBrowsers (state, countries) {
        state.countries = countries
    },
}

const actions = {
    async getCountries ({commit}) {
        try {
            const { data } = await axios.get('/api/countries');
            commit('setCountries', data)
            return true
        } catch(err) {
            console.log(err)
            return false;
        }
    },
    async getOS ({commit}) {
        try {
            const { data } = await axios.get('/api/os');
            commit('setCountries', data)
            return true
        } catch(err) {
            console.log(err)
            return false;
        }
    },
    async getBrowsers ({commit}) {
        try {
            const { data } = await axios.get('/api/browsers');
            commit('setCountries', data)
            return true
        } catch(err) {
            console.log(err)
            return false;
        }
    }
}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
