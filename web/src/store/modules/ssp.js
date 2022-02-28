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
    dspItem: (state) => (id, name) => {
        return state.sspItem.dsp[id][name]
    }
}

const mutations = {
    setSSP(state, ssp) {
        state.ssp = ssp
    },
    setSSPItem(state, ssp) {
        state.sspItem = ssp
    },
    clearCurSSP(state, ssp) {
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
    deleteSSP(state, id) {
        state.ssp = state.ssp.filter(item => item.ssp_id !== id)
    },
    removeDSPinCurSSP(state, index) {
        state.sspItem.dsp = state.sspItem.dsp.filter((_, i) => {
            return i !== index
        })
    },
    dspItemUpdate(state, {index, name, value}) {
        state.sspItem.dsp[index][name] = value
    },
    addSSP (state) {
        let find = false
        state.ssp = state.ssp.map(item => {
            if (item.ssp_id === state.sspItem.ssp_id) {
                find = true
                return state.sspItem
            }
            return item
        })
        if (!find) {
            state.ssp.push(state.sspItem)
        }
    },
}

const actions = {
    async getSSP({commit}) {
        try {
            let {data} = await axios.get('/api/ssp');
            data = data.map((item) => {
                return {
                    ...item,

                }
            })
            commit('setSSP', data)
            return true
        } catch (err) {
            console.log(err)
            return false;
        }
    },
    async setSSP ({commit}) {
        const err = await axios.post('/api/ssp', state.sspItem).then(
            () => {
                commit('addSSP', state.dspItem)
                return true
            }).catch( err => {
            console.log(err)
            return false;
        })
    },
    async deleteSSP({commit}, id) {
        const err = await axios.delete('/api/ssp', {params: {"id": id}})
            .then(() => {
                commit('deleteSSP', id)
                return true
            })
            .catch(err => {
                console.log(err)
                return false;
            })
    },
}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
