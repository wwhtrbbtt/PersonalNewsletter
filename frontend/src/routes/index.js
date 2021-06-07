import Vue from 'vue';
import Router from 'vue-router';
import Dashboard from '../components/Dashboard'
import Home from '../components/Home'
import Settings from '../components/Settings'
import Login from '../components/Login'
import Register from '../components/Register'
import Letters from '../components/Letters'

// import Dashboard from '../components/Dashboard'

Vue.use(Router)

const router = new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: "/dashboard",
            name: "dashboard",
            component: Dashboard,
            meta: {
                authRequired: false,
              },
        },
        {
            path: "/",
            name: "home",
            component: Home
        },
        {
            path: "/settings/",
            name: "settings",
            component: Settings
        },
        {
            path: "/login/",
            name: "login",
            component: Login
        },
        {
            path: "/register/",
            name: "register",
            component: Register
        },
        {
            path: "/letters/",
            name: "letters",
            component: Letters
        },
]
});

import firebase from "firebase";

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.authRequired)) {
      if (firebase.auth().currentUser) {
        next();
      } else {
        // alert('You must be logged in to see this page');
        next({
          path: '/',
        });
      }
    } else {
      next();
    }
  });

export default router