const state = {
    lang: null,
}

const getters = {
    lang: state => state.lang,
}

const mutations = {}

const actions = {}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
