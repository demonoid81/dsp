import Vue from 'vue';
import ViewUI from 'view-design';
import 'view-design/dist/styles/iview.css';
import App from './App.vue';
import store from "store";
import router from "router";
import i18n from './locale'
import config from "config";



Vue.use(ViewUI);

Vue.prototype.$config = config

new Vue({
    el: '#app',
    router,
    store,
    i18n,
    render: h => h(App)
})