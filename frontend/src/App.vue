<template>
<div id="app">
  <Settings />
  <!-- <home /> -->
</div>
</template>

<script>
// import Home from './components/Home.vue'
import Settings from './components/Settings.vue'
import { db } from './dbs'

export default {
  name: 'App',
  components: {
    // Home,
    Settings
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
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
