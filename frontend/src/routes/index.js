import Vue from 'vue';
import Router from 'vue-router';
import Dashboard from '../components/Dashboard'
import Home from '../components/Home'
// import Login from '../components/Login'
// import Register from '../components/Register'
// import Dashboard from '../components/Dashboard'

Vue.use(Router)

const router = new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: "/dashboard",
            name: "dashboard",
            component: Dashboard
        },
        {
            path: "/",
            name: "home",
            component: Home
        }
]
});

export default router