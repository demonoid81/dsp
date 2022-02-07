import axios from "axios";

const state = {
    dsp: [],
    dspItem: {},
}

const getters = {
    dsp: state => state.dsp,
    dspItem: state => state.dspItem
}

const mutations = {
    setDSP (state, dsp) {
        state.dsp = dsp
    },
    setDSPItem (state, dsp) {
        state.dspItem = dsp
    },
    clearDSPItem (state) {
        state.dspItem = {}
    },
    setCurSSPItem(state, {value, name}) {
        console.log(value, name)
        state.curSSP[name] = value
    },
    addDPSInCurSSP(state) {
        state.curSSP.dsp.push({})
    },
    removeDSPinCurSSP(state, index) {
        state.curSSP.dsp = state.curSSP.dsp.filter((_, i) => {
            return i !== index
        })
    },
    enableDSP(state, {value, index}) {
        state.curSSP.dsp = state.curSSP.dsp.map((item, i) => {
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
    async getDSP ({commit}) {
        try {
            const { data } = await axios.get('/api/dsp');
            commit('setDSP', data)
            return true
        } catch(err) {
            console.log(err)
            return false;
        }
    },
}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
