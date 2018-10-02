// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import store from './store'
import axios from 'axios'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.prototype.$http = axios.create({
  headers: {
    'X-Requested-With': 'XMLHttpRequest'
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  components: {
    App
  },
  store,
  template: '<App/>'
})
