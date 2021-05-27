<!-- https://vuefire.vuejs.org/vuefire/querying.html#one-time-read -->
<!-- https://rimdev.io/the-v-for-key/ -->
<template>
  <div class="main">
      <h1>Settings</h1>
      <button v-on:click="saveChanges">save changes</button><br><br>
      Letter name<br>
      <input type="text" v-model="config.Feedname" /><br><br>
      Letter time<br>
      <input type="time" v-model="config.Time" /><br><br>
      Greeting name<br>
      <input v-model="config.Greetingname" /><br><br>
      Email address<br>
      <input v-model="config.Email" /><br><br>

      <div v-for="(module, index) in config.Modules" :key="index">
        <hr>
        <h2>{{ module.Name }}</h2>
        <div v-for="(setting, index2) in module.Settings" :key="index2">
          <h3>{{ setting.Name }}</h3>
            <input v-model="setting.Value" /><br><br>
        </div>
      </div>
  </div>
</template>

<script>

import { db } from '../dbs'

export default {
  name: 'Dashboard',

  data: () => ({ 
    config: {}
   }),

  created() {
    this.loadLetter()
  },

  methods: {
    saveChanges: function() {
      console.log("saved")
      console.log(this.config)
      db.collection("NewLetters")
      .doc("test")
      .set(this.config)
      .then(() => {
        console.log('user updated!')
        alert("saved")
      })
    },

    loadLetter: function() {
      db.collection("NewLetters")
      .doc("test")
      .get()
      .then(snapshot => {
        const data = snapshot.data()
        this.config = data
        console.log(data)
      })
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
