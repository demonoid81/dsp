import axios from "axios";

const state = {
    ssp: [],
    sspItem: {
        dsp: []
    }
}

const getters = {
    ssp: state => state.ssp,
    sspItem: state => state.sspItem,
}

const mutations = {
    setSSP (state, ssp) {
        state.ssp = ssp
    },
    setSSPItem (state, ssp) {
        state.sspItem = ssp
    },
    clearCurSSP (state, ssp) {
        state.sspItem = {
            dsp: []
        }
    },
    setCurSSPItem(state, {value, name}) {
        console.log(value, name)
        state.sspItem[name] = value
    },
    addDPSInCurSSP(state) {
        state.sspItem.dsp.push({})
    },
    removeDSPinCurSSP(state, index) {
        state.sspItem.dsp = state.sspItem.dsp.filter((_, i) => {
            return i !== index
        })
    },
    enableDSP(state, {value, index}) {
        state.sspItem.dsp = state.sspItem.dsp.map((item, i) => {
            if (i !== index) {
                return {
                    ...item,
                    enabled: value
                }
            } else {
                return item
            }
        })
    },
}

const actions = {
    async getSSP ({commit}) {
        try {
            let { data } = await axios.get('/api/ssp');
            data = data.map((item) => {
                return {
                    ...item,

                }
            })
            commit('setSSP', data)
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
