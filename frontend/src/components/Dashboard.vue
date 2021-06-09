<!-- https://vuefire.vuejs.org/vuefire/querying.html#one-time-read -->
<!-- https://rimdev.io/the-v-for-key/ -->

<template>
  <div class="main" id="main" v-if="show">
      <h1>Edit your newsletter</h1>

      <Input text="Newsletter name"
             type="text" 
             placeholder="eg. Sport newsletter" 
             v-model="config.Feedname" />
      <br><br>

      <Input text="Newsletter time"
             type="text" 
             placeholder="eg. 20:30" 
             v-model="config.Time" />

      <br>
      <div v-if="timeWarning != ''" style="color: red;"><br>{{ timeWarning }}</div>
      <br>

      <h1>Modules in your newsletter</h1>

      <p v-if="!config.Modules.length">
         No modules yet - add one!
      </p>


      <div v-for="(module, index) in config.Modules" :key="index">

        <i><h2>{{ modulesInfo[module.Name].ShowName }}</h2></i>
        <p>{{ modulesInfo[module.Name].Description }}</p>
        
        <div v-for="(setting, index2) in module.Settings" :key="index2">

        <Input :text="setting.Name"
              :type="parseType(modulesInfo[module.Name].Settings[index2].Type)" 
              v-model="setting.Value" />
        <br><br>

        </div>
        <button class="save-button s-b-small" v-on:click="removeModule(index)">Remove</button>
      </div>

    <br>
    <Dropdown title="Add a module" :options="possibleModulesDropdown()" v-model="chosenModule"/>
    <br><br>

    <div style="text-align: center;">
      <button type="button" class="save-button" v-on:click="saveChanges">Done</button><br><br>
      <button type="button" class="save-button" v-on:click="deleteLetter">Delete</button><br><br>
    </div>
  </div>
</template>

<script>
import db from "../main"

import firebase from 'firebase/app';
import 'firebase/auth';
import Input from './ui/Input.vue';
import Dropdown from './ui/Dropdown.vue';



export default {
  components: { Input, Dropdown },
  name: 'Dashboard',

  data: () => ({ 
    config: {
      Modules: [],
    },
    possibleModules: [],
    modulesInfo: {},
    show: false,
    chosenModule: "",
    test: "",
   }),

   created() {
    let id = this.$route.query.letter
    firebase.auth().onAuthStateChanged( user => {
        if (user) {
            this.loadPossibleModules()
            this.loadLetter(id)
        } else {
            console.log("not logged in")
        }
    })
  },

  computed: {

    timeWarning() {
      try {
        let hours = parseInt(this.config.Time.split(":")[0])
        let minutes = parseInt(this.config.Time.split(":")[1])

        if (this.config.Time.match(/\d{2}:\d{2}/)[0] == this.config.Time && hours < 24 && minutes < 60) {
          return ""
        }
      } catch {
        return "Invalid time. Use the 24h format (eg.  20:00)"
      }
      return "Invalid time. Use the 24h format (eg.  20:00)"
    },
  },

  watch: {
    chosenModule: function() {
      this.addModule()
      this.chosenModule = ""
    }
  },

  methods: {
    possibleModulesDropdown: function() {
      let options = []
      this.possibleModules.map( module => {
        options.push({
          Key: module.InternalName,
          Value: module.ShowName
        })
      })
      return options
    },

    saveChanges: function() {

      if (this.timeWarning != "") {
        alert("Couldn't save letter")
        return
      }

      // the numbers from the input are always strings. We wanna parse them back to ints
      this.config.Modules.forEach(module => {
        let count = 0
        module.Settings.forEach(setting => {
          if (this.modulesInfo[module.Name].Settings[count].Type == "int") {
            setting.Value = parseInt(setting.Value)
          }
          count++
        })
      })

      console.log("saved")
      // console.log(this.config)
      let id = this.$route.query.letter
      db.collection("NewLetters")
      .doc(id)
      .set(this.config)
      .then(() => {
        console.log('letter updated!')
        // alert("Saved all changes")
        this.$router.push({ name: 'letters' })
      })
    },

    loadLetter: function(id) {
      db.collection("NewLetters").doc(id).get()
      .then(snapshot => {
        const data = snapshot.data()
        this.config = data
        this.show = true
      })
      .catch(err => {
        console.log(err.message)
        alert(err.message)
      })
    },

    loadPossibleModules: function() {
      db.collection("misc").doc("newModules").get()
      .then(snapshot => {
        this.possibleModules = snapshot.data().Modules
      })
      .then( () => {
          // make a dict out of the data.

          let tmp = {}
          this.possibleModules.forEach(module => {
            tmp[module.InternalName] = module
          })

          this.modulesInfo = tmp
      })
      .catch(err => {
        console.log(err.message)
        alert(err.message)
      })
    },

    deleteLetter: function() {
      if (!confirm(`Do you really want to delete '${this.config.ShowName}'? This can't be undone.`)) {
        return
      }

      let id = this.$route.query.letter
      db.collection("NewLetters").doc(id).delete()
        .then(() => {
          console.log("deleted letter")
          this.$router.push({ name: 'letters' })
        })
        .catch(err => {
          alert(err.message)
        }) 
    },
    addModule: function() {
      
      console.log(this.chosenModule)
      const m = this.modulesInfo[this.chosenModule];
      // console.log(m)
      if (this.chosenModule == "") {
        return
      }

      // return
      console.log(m);
      let settings = [];
      if (m.Settings != null) {
        m.Settings.forEach(function(setting) {
          let val = null
          if (setting.Type == "str") {
            val = ""
          } else if (setting.Type == "int") {
            val = 5
          } else { // more types in the future
            val = ""
          }

          settings.push({
            Name: setting.InternalName,
            Value: val,
          })
        })
      } 

      this.config.Modules.push({
        Name: this.chosenModule,
        Settings: settings,
      })
      this.chosenModule = ""
    },
    parseType: function (raw) {
      if (raw == "int") {
        return "number"
      } else if (raw == "string") {
        return "text"
      } else {
        return "text"
      }
    },
    removeModule: function (index) {
      console.log(index)
      this.config.Modules.splice(parseInt(index), 1)
      console.log(this.config)
    }
  }, 


}
</script>

<style scoped>

</style>
