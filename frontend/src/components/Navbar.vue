<template>
<div class="main" id="navbar">
  <nav class="topnav">
        <router-link v-bind:class="{active: currentRouteName == 'home'}" to="/">Home</router-link>
        <router-link v-if="loggedIn" v-bind:class="{active: currentRouteName == 'letters'}" to="/Letters/">Letters</router-link>
        <router-link v-if="loggedIn" v-bind:class="{active: currentRouteName == 'settings'}" to="/settings/">Settings</router-link>
        <router-link v-if="!loggedIn" v-bind:class="{active: currentRouteName == 'register'}" to="/register/">Register</router-link>
        <router-link v-if="!loggedIn" v-bind:class="{active: currentRouteName == 'login'}" to="/login/">Login</router-link>
  </nav>
</div>
</template>

<script>
import firebase from "firebase";

export default {
  data: function() {
    return {
      loggedIn: false
    }
  },
  created: function() {
        firebase.auth().onAuthStateChanged(user => {
        if (user) {
            this.loggedIn = true
            console.log("logged in")
        } else {
          this.loggedIn = false
          console.log("not logged in")
        }
        });
  },

  computed: {
      currentRouteName() {
        return this.$route.name;
      },
  }
}
</script>

<style scoped>


</style>