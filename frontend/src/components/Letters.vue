<template>
    <div id="main" class="main" v-if="show">
        <h1>Your letters</h1>
        <p v-if="!letters.length">
            No letters yet - create one!
        </p>


        <div v-for="(letter, index) in letters" :key="index">
            <h3>{{ letter.Feedname }}</h3>
            <p>{{ letter.Modules.length }}  Modules, comes at {{ letter.Time }}</p>
            <button class="save-button s-b-small" v-on:click="$router.push({ name: 'dashboard', query: { letter: letter.id }})">
                Edit
            </button>
        </div>

        <br><br><br>
        <div style="text-align: center;">
            <button class="save-button" v-on:click="addLetter">Add letter</button>
        </div>
    </div>
</template>

<script>

import db from "../main"
import firebase from 'firebase/app';
import 'firebase/auth';


export default {
    data() {
        return {
            letters: [],
            user: {},
            show: false
        }
    },
    methods: {
        getLetters: function() {
            db.collection("NewLetters").where("User", "==", this.user.uid).get()
                .then(docs => {
                    docs.forEach(doc => {
                        let tmp = doc.data()
                        tmp.id = doc.id
                        this.letters.push(tmp)
                    })
                    this.show = true
                })
                .catch(err => {
                    console.log(err.message)
                    alert(err.message)
                })
        },

        addLetter: function() {
            console.log(this.user)
            if (this.user == {}) {
                console.log("not logged in")
                return
            }
            db.collection("NewLetters").add({
                Email: this.user.email,
                Feedname: "Newsletter " + (this.letters.length + 1),
                Greetingname: "name",
                Time: "10:00",
                User: this.user.uid,
                Modules: []
            })
            .then(ref => {
                console.log("Document written with ID: ", ref.id);
                this.$router.push({ name: 'dashboard', query: { letter: ref.id }})
            })
            .catch(err => {
                console.log(err.message)
            })
        }
    },

    created() {
        firebase.auth().onAuthStateChanged( user => {
        if (user) {
            this.user = user
            console.log(user.uid)
            this.getLetters()
        } else {
            console.log("not logged in")
            this.user = {}
        }
        });
    }
    
}
</script>