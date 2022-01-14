import axios from 'axios';
import Cookies from 'js-cookie'
import config from 'config/index'
const {cookieExpires} = config


const getToken = () => {
    const token = Cookies.get('token')
    if (token) return token
    else return false
}

const getters = {
    user: state => state.user,
    token: state => state.token,
}

const mutations = {
    setToken (state, token) {
        console.log(token)
        state.token = token
        Cookies.set("token", token, { expires: cookieExpires || 1 })
    },
}

const actions = {
    async handleLogin({commit}, {username, password}) {
        username = username.trim()
        try {
            const { data } = await axios.get('/auth', {
                params: {
                    username: username,
                    password: password
                }
            });
            console.log(data)
            commit('setToken', data)
            return true
        } catch(err) {
            console.log(err)
            return false;
        }
    }
}

export default {
    namespaced: true,
    state: {
        user: null,
        token: getToken()
    },
    actions,
    mutations,
    getters,
};
