<template>
  <div class="main" id="main">
      <h1>PersonalNewsletter</h1>
      <h2>Create your own customized newsletter in about 3 minutes.</h2>
      <p>
        The idea is simple: get all the information you need, every day, ad-free.
        Create your own newsletter, add a few modules, configure it a bit and thats it!
        
        
      </p>
      <br><br>
      <div style="text-align: center;">
        <button v-if="!loggedIn" class="save-button" v-on:click="$router.push('/register')">Get started</button>
        <button v-if="loggedIn" class="save-button" v-on:click="$router.push('/letters')">Your letters</button>
      </div>
  </div>
</template>

<script>
import firebase from "firebase/app";
import "firebase/auth";

export default {
  name: 'Home',

  created() {
    firebase.auth().onAuthStateChanged( user => {
    if (user) {
        this.loggedIn = true
    } else {
        this.loggedIn = false
    }
    });
  },
  data: function() {
    return {
      loggedIn: false,
    }
  },
}
</script>

<style scoped>
h1, h2, button {
  text-align: center;
}
p {
  text-align: justify;
  font-size: 17px;
}

</style>
