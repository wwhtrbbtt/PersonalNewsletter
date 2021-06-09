<template>
    <div id="main" class="main">
        <form @submit.prevent="register">     
            <h2>Register</h2>     

        <Input text="Email"
             type="email" 
             v-model="email" />
        <br><br>

        <Input text="Name"
             type="text" 
             v-model="name" />
        <br><br>

      <Input text="Password"
             type="password" 
             v-model="password" />
        <br><br>    

      <Input text="Confirm password"
             type="password" 
             v-model="passwordConfirm" />
        <br><br>  
        <p v-if="error != ''">{{ error }}</p>
        <button type="submit" class="save-button">
            Register
        </button>   
        </form> 
        
    </div>
</template>

<script>
import firebase from "firebase/app";
import "firebase/auth";
import Input from './ui/Input.vue';


export default {
    components: { Input },
    data: () => ({ 
        email: "",
        password: "",
        passwordConfirm: "",
        error: "",
        name: "",
    }),

    methods: {
        // createProfile: function(id) {
        //         console.log(id)
        //         db.collection("users").doc(id).set({
        //             letters: {},
        //             user: id
        //         })
        //         .then(() => {
        //             console.log("Document successfully written!");
        //         })
        //         .catch((error) => {
        //             console.error("Error writing document: ", error);
        //         });
        // },
        setName: function(name) {
            let userNow = firebase.auth().currentUser;
            userNow.updateProfile({
                displayName: name,
            }).then(
                console.log("updated display name")
            ).catch(err => {
                console.log(err.message)
            })
        },

        register: function() {
            console.log(this.password, this.passwordConfirm, this.error)
            this.error = ""
            if (this.password != this.passwordConfirm) {
                this.error = "Passwords do not match"
                return
            }
            if (this.password.lenght < 6) {
                this.error = "Password has to be longer"
                return
            }
            firebase
                .auth()
                .createUserWithEmailAndPassword(this.email, this.password)
                .then(() => {
                    this.$router.push("/letters/")
                    firebase.auth().onAuthStateChanged( user => {
                        if (user) {
                            console.log("logged in")
                            this.setName(this.name)
                        } else {
                            console.log("not logged in")
                        }
                    })
                })
                .catch(error => {
                    this.error = error.message
                })
        }
    }
}
</script>

<style scoped>
#main {
    text-align: center;
}
</style>