<template>
<div id="app">
  <navbar></navbar>
  <router-view></router-view>
  <span class="navbar-toggler-icon"></span>
</div>
</template>

<script>
// import Home from './components/Home.vue'
// import Dashboard from './components/Dashboard.vue'
import { db } from './dbs'
import navbar from "./components/Navbar";

export default {
  name: 'App',
  components: {
    navbar
    // Home,
    // Dashboard
  },
    data() {
    return {
      documents: [],
    }
  },

  firestore: {
    documents: db.collection('documents'),
  },
}


</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  /* margin-top: 60px; */
  margin: 0;
}
</style>
