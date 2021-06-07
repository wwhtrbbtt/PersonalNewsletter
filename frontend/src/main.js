import Vue from 'vue'
import App from './App.vue'
import router from "./routes/index";
import firebase from 'firebase'

const firebaseConfig = {
  apiKey: "AIzaSyAjJxeUAweVr4HmfGjIdMpkPYGbk0THr50",
  authDomain: "newsletter-43a2b.firebaseapp.com",
  databaseURL: "https://newsletter-43a2b-default-rtdb.europe-west1.firebasedatabase.app",
  projectId: "newsletter-43a2b",
  storageBucket: "newsletter-43a2b.appspot.com",
  messagingSenderId: "385030942753",
  appId: "1:385030942753:web:c0f4c5e89815b5adb93360",
  measurementId: "G-ZYCMD17ZYS"
};

// Vue.use(firestorePlugin)
Vue.config.productionTip = false
firebase.initializeApp(firebaseConfig);


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

const db = firebase.firestore();

export default db