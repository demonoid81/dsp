import Vue from 'vue'
import Router from 'vue-router'
import ViewUI from 'view-design';
import routes from './routers'
import { setTitle } from 'libs/utils'
import store from "store/index";

import config from 'config/index'
const {homeName} = config

Vue.use(Router)
const router = new Router({
    routes,
    mode: 'history'
})

const LOGIN_PAGE_NAME = 'login'

router.beforeEach((to, from, next) => {
    ViewUI.LoadingBar.start();
    const token = store.getters['user/token']
    if (!token && to.name !== LOGIN_PAGE_NAME) {
        console.log("login")
        next({
            name: LOGIN_PAGE_NAME
        })
    } else if (!token && to.name === LOGIN_PAGE_NAME) {
        next()
    } else if (token && to.name === LOGIN_PAGE_NAME) {
        next({
            name: homeName
        })
    } else {
        next()
    }
})

router.afterEach(to => {
    setTitle(to, router.app)
    ViewUI.LoadingBar.finish()
    window.scrollTo(0, 0)
})

export default router
