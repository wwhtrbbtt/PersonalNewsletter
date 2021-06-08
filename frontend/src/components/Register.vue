<template>
    <div id="main" class="main">
        <form @submit.prevent="register">     
            <h2>Register</h2>     
        <label class="field field_v1">
            <input class="field__input" type="email" v-model="email" /><br><br>
                <span class="field__label-wrap">
                <span class="field__label">Email</span>
                </span>
        </label>
        <br><br>

        <label class="field field_v1">
            <input class="field__input" type="password" v-model="password" /><br><br>
                <span class="field__label-wrap">
                <span class="field__label">Password</span>
                </span>
        </label>
        <br><br>    

        <label class="field field_v1">
            <input class="field__input" type="password" v-model="passwordConfirm" /><br><br>
                <span class="field__label-wrap">
                <span class="field__label">Confirm password</span>
                </span>
        </label>
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
import db from "../main"

export default {
    data: () => ({ 
        email: "",
        password: "",
        passwordConfirm: "",
        error: ""
    }),

    methods: {
        createProfile: function(id) {
                console.log(id)
                db.collection("users").doc(id).set({
                    letters: {},
                    user: id
                })
                .then(() => {
                    console.log("Document successfully written!");
                })
                .catch((error) => {
                    console.error("Error writing document: ", error);
                });
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
                })
                .catch(error => {
                    this.error = error.message
                })
            const id = firebase.auth().currentUser.id
            this.createProfile(id)

            
        }
    }
}
</script>

<style scoped>
#main {
    text-align: center;
}
</style>