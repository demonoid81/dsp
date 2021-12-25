import Vue from 'vue';
import VueRouter from 'vue-router';

import ViewUI from 'view-design';
import 'view-design/dist/styles/iview.css';

// Import Vue App, routes, store
import App from './App.vue';
import routes from './routes';

Vue.use(VueRouter);
Vue.use(ViewUI);

// Configure router
const router = new VueRouter({
    routes,
    linkActiveClass: 'active',
    mode: 'history'
});

new Vue({
    el: '#app',
    router,
    render: h => h(App)
})