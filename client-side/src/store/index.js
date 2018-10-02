import Vue from 'vue'
import Vuex from 'vuex'
import global from './modules/global'
import constvar from './modules/const' 

Vue.use(Vuex)

const store = new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    global,
    constvar
  }
})

export default store
