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
    addDSP (state) {
        let find = false
        state.dsp = state.dsp.map(item => {
            console.log(item)
            if (item.id === state.dspItem.id) {
                find = true
                return state.dspItem
            }
            return item
        })
        if (!find) {
            state.dsp.push(state.dspItem)
        }
    },
    deleteDSP (state, id) {
        state.dsp = state.dsp.filter(item => item.id !== id)
    },
    setDSPItem (state, dsp) {
        state.dspItem = dsp
    },
    clearDSPItem (state) {
        state.dspItem = {}
    },
    setDSPItemField(state, {value, name}) {
        state.dspItem[name] = value
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
    async setDSP ({commit}) {
            const err = await axios.post('/api/dsp', state.dspItem).then(
                () => {
                    commit('addDSP', state.dspItem)
                    return true
                }).catch( err => {
                console.log(err)
                return false;
            })
    },
    async deleteDSP ({commit}, id) {
        const err = await axios.delete('/api/dsp', {params: {"id": id}})
            .then(() => {
                commit('deleteDSP', id)
                return true
            })
            .catch( err => {
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
