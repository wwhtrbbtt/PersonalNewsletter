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
        <p>Email: {{ user.email }}</p>
        <p>User ID: {{ user.uid }}</p>
        <h2>Profile settings</h2>
        <button class="save-button s-b-small" v-on:click="logout">Logout</button>
        <br><br>
        <button class="save-button s-b-small">Delete account</button>
    </div>
</template>

<script>
// import { getCookie, setCookie } from 'tiny-cookie'
//import firebase from "firebase";
import firebase from 'firebase/app';
import 'firebase/auth'

export default {
    created() {
        firebase.auth().onAuthStateChanged(user => {
            if (user) {
                this.user = user
                this.show = true
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
    },
    data: () => ({ 
        // currentColorScheme: "pastel",
        user: {},
        show: false,
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