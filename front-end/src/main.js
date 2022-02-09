import Vue from 'vue'
import App from './App.vue'

// import axios from 'axios'
// axios.defaults.baseURL = process.env.VUE_APP_API_URL;
// Vue.prototype.$http = axios;

// import axios from 'axios'
// import VueAxios from 'vue-axios'

// Vue.use(axios);
// Vue.use(VueAxios);
// Vue.prototype.$http = axios
// app.axios = axios;
// app.$http = axios;
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
