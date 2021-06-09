<template>
    <div id="main" v-if="show">
     <!-- Color picker -->
     <!-- <h1>Colorscheme of the website</h1>
        <select v-model="currentColorScheme" >
            <option value="pastel">Pastel</option>
            <option value="slight">Simple light</option>
        </select> 
        <button v-on:click="applyColor">Apply</button>
     -->
        <h1>
            Account
        </h1>
        <h2>Profile information</h2>

        <div v-if="!edit">
            <p>Email: {{ email }}</p>
            <p>Your name: {{ name }}</p>
            <button class="save-button s-b-small" v-on:click="edit = true">Edit</button>
        </div>

        <div v-if="edit">
            <p>Email: {{ email }}</p>
            <p>You can't edit your email (yet)</p>

            <Input v-if="edit" text="Name"
                type="text" 
                v-model="name" />
            <br><br>

            <button v-if="edit" class="save-button s-b-small" v-on:click="saveChanges">Save</button>
        </div>

        <h2>Profile settings</h2>
        <button style="margin-right:20px;" class="save-button s-b-small" v-on:click="logout">Logout</button>
             
        <button class="save-button s-b-small" v-on:click="deleteAccount">Delete account</button>
        <p style="font-size: 10px;">User ID: {{ user.uid }}</p>
    </div>
</template>

<script>
// import { getCookie, setCookie } from 'tiny-cookie'
//import firebase from "firebase";
import firebase from 'firebase/app';
import 'firebase/auth'
import db from "../main"
import Input from './ui/Input.vue';

export default {
    components: { Input },
    created() {
        firebase.auth().onAuthStateChanged(user => {
            if (user) {
                this.user = user
                this.show = true
                this.name = user.displayName
                this.email = user.email
                console.log("logged in")
            } else {
            this.show = false
            console.log("not logged in")
            }
        });

        // if (getCookie("theme") == null) {
        //     setCookie('theme', 'pastel', { expires: '1Y' });
        //     this.currentColorScheme = "pastel"
        // } else {
        //     this.currentColorScheme = getCookie("theme")
        // }
        // this.applyColor()
    },

    methods: {
        // applyColor() {
        //     console.log(this.currentColorScheme)
        //     setCookie('theme', this.currentColorScheme, { expires: '1Y' })

        //     // Thanke WmeAllen :)

        //     let colors = this.colorSchemes[this.currentColorScheme]
        //     document.documentElement.style.setProperty('--text-dark', "#" + colors[0]); // text on page
        //     document.documentElement.style.setProperty('--text-bright', "#" + colors[1]); // text in navbar
        //     document.documentElement.style.setProperty('--background', "#" + colors[2]); // backrgound of site

        //     document.documentElement.style.setProperty('--ui1', "#" + colors[3]); // active input
        //     document.documentElement.style.setProperty('--ui2', "#" + colors[4]); // inactive input

        //     document.documentElement.style.setProperty('--nav1', "#" + colors[5]); // hover navbar
        //     document.documentElement.style.setProperty('--nav2', "#" + colors[6]); // inactive navbar
        //     document.documentElement.style.setProperty('--nav3', "#" + colors[7]); // active navbar
        // },
        logout: function() {
            firebase
                .auth()
                .signOut()
                .then(() => {
                console.log("logged out")
                this.$router.push('/')
                })
                .catch(error => {
                alert(error.message)
                this.$router.push('/');
                })
            },

        deleteAccount: function() {
            if (confirm("Do you really want to delete your account? All of your data will be deleted.")) {
               
                db.collection("NewLetters").where("User", "==", this.user.uid).get()
                .then(docs => {
                    docs.forEach(doc => { 
                        doc.get().then(snap => {
                            snap.delete().then(() => {
                                console.log("deleted letter")
                            })
                            .catch(err => {
                                console.log(err.message)
                            
                            })
                        })
                        .catch(err => {
                            console.log(err.message)
                        })      
                    })
                })
               
               // delete account
                firebase
                    .auth()
                    .currentUser
                    .delete()
                    .then(() => {
                        this.$router.push("/login/")
                    })
                    .catch(err => {
                        if (err.message == "This operation is sensitive and requires recent authentication. Log in again before retrying this request.") {
                            alert("Logout and login before deleting your account")
                            this.$router.push("/login/")
                        }
                    })
                }
        },
        saveChanges: function() {

            // save changes to user
            let userNow = firebase.auth().currentUser;

            userNow.updateProfile({
                displayName: this.name,
                // email: this.email,
            }).then(
                console.log("updated display name"),
            ).catch(err => {
                console.log(err.message)
                return
            })

            // edit every newsletter the user has
            db.collection("NewLetters").where("User", "==", this.user.uid).get()
            .then(docs => {
                    docs.forEach(doc => { 

                        console.log(doc.id)
                        db.collection("NewLetters").doc(doc.id).update({
                            // Email: this.email,
                            Greetingname: this.name,
                        })
                        .then( () => {
                            console.log("updated element")
                        })
                        .catch( err => {
                            console.log(err.message)
                        })
                    })
                })            
            this.edit = false
        }
    },
    data: () => ({ 
        // currentColorScheme: "pastel",
        user: {},
        show: false,
        edit: false,
        email: "",
        name: "",
        // colorSchemes: {
        //     "pastel": 
        //         ["000000", "f8f8f8", "ffe8d6", "cb997e", "ddbea9", "b7b7a4", "a5a58d", "6b705c"],
        //     "slight": 
        //         ["000", "000", "bbb", "ddd", "ccc", "fff", "888", "6b705c"],
        //     "sdark":
        //         ["fff", "000", "000", "bbb", "ccc", "000", "888", "6b705c"]
        //         // ["f8f8f8", "000000", "F6B48E", "CC978E", "F6B48E", "C0BDA5", "261447", "6b705c"]
        // },
   }),
}
</script>

<style scoped>

</style>