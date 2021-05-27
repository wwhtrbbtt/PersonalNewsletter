import Vue from 'vue'
import App from './App.vue'
import router from "./routes/index";
import { firestorePlugin } from 'vuefire'

Vue.use(firestorePlugin)
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')




// const db = firebase.initializeApp({ projectId: 'newsletter-43a2b' }).firestore()
// const { Timestamp, GeoPoint } = firebase.firestore
// export { Timestamp, GeoPoint }

// new Vue({
//   // setup the reactive todos property
//   data: () => ({ letters: [] }),

//   firestore: {
//     letters: db.collection('letters'),
//   },
// })
