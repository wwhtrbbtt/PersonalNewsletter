<template>
    <div id="main" class="main">
        <form @submit.prevent="login">     
            <h2>Login</h2>     

      <Input text="Email"
             type="email" 
             v-model="email" />
        <br><br>

      <Input text="Password"
             type="password" 
             v-model="password" />
        <br><br>    

        <p v-if="error != ''">{{ error }}</p>
            <button type="submit" class="save-button">
                Login
            </button>   
        </form> 
        <p v-on:click="forgotPassword" style="cursor: pointer;">Reset password</p>
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
        error: ""
        }),
    methods: {
        forgotPassword: function() {
            this.error = ""
            if (this.password == "") {
                this.error = "Enter your email"
            }
            firebase.auth().sendPasswordResetEmail(this.email)
                .then( () => {
                    this.error = "Send an email. Please also check your spam folder."
                })
                .catch( err => {
                    this.error = err.message
                })
        },
        login: function() {
        this.error = ""
        if (this.password == "") {
                this.error = "Enter your email"
        }
        firebase
            .auth()
            .signInWithEmailAndPassword(this.email, this.password)
            .then(() => {
              this.$router.push("/letters/")
            })
            .catch(error => {
                if (error.message == 'The password is invalid or the user does not have a password.') {
                    this.error = "The password is invalid"
                } else if (error.message == 'There is no user record corresponding to this identifier. The user may have been deleted.') {
                    this.error = "Unknown email address"
                } else {
                    this.error = error.message
                }
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