<template>
    <div id="main" class="main">
        <form @submit.prevent="login">     
            <h2>Login</h2>     
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

        <p v-if="error != ''">{{ error }}</p>
            <button type="submit" class="save-button">
                Login
            </button>   
        </form> 
    </div>
</template>

<script>
import firebase from "firebase/app";
import "firebase/auth";

export default {
    data: () => ({ 
        email: "",
        password: "",
        error: ""
        }),
    methods: {
        login: function() {
        this.error = ""
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