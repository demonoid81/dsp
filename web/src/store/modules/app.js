import config from "config";
import routers from "router/routers";

const { homeName } = config

export const getHomeRoute = (routers, homeName = 'home') => {
    let i = -1
    let len = routers.length
    let homeRoute = {}
    while (++i < len) {
        let item = routers[i]
        if (item.children && item.children.length) {
            let res = getHomeRoute(item.children, homeName)
            if (res.name) return res
        } else {
            if (item.name === homeName) homeRoute = item
        }
    }
    return homeRoute
}

export const localSave = (key, value) => {
    localStorage.setItem(key, value)
}

export const localRead = (key) => {
    return localStorage.getItem(key) || ''
}

export const hasChild = (item) => {
    return item.children && item.children.length !== 0
}

export const hasOneOf = (targetarr, arr) => {
    return targetarr.some(_ => arr.indexOf(_) > -1)
}

const showThisMenuEle = (item, access) => {
    // if (item.meta && item.meta.access && item.meta.access.length) {
    //     if (hasOneOf(item.meta.access, access)) return true
    //     else return false
    // } else
        return true
}

export const getMenuByRouter = (list, access) => {
    let res = []
    list.forEach(item => {
        if (!item.meta || (item.meta && !item.meta.hideInMenu)) {
            let obj = {
                icon: (item.meta && item.meta.icon) || '',
                name: item.name,
                meta: item.meta
            }
            if ((hasChild(item) || (item.meta && item.meta.showAlways))) {
                // && showThisMenuEle(item, access)) {
                obj.children = getMenuByRouter(item.children, access)
            }
            if (item.meta && item.meta.href) obj.href = item.meta.href
            if (showThisMenuEle(item, access)) res.push(obj)
        }
    })
    return res
}

const state = {
    local: null,
    homeRoute: {},
}

const getters = {
    local: state => state.local,
    menuList: (state, getters, rootState) => {
        console.log(routers)
       return getMenuByRouter(routers, rootState.user.access)
    }
}

const mutations = {
    setHomeRoute (state, routes) {
        state.homeRoute = getHomeRoute(routes, homeName)
    },
    setLocal (state, lang) {
        localSave('local', lang)
        state.local = lang
    },
}

const actions = {}

export default {
    namespaced: true,
    state: () => state,
    actions,
    mutations,
    getters,
};
