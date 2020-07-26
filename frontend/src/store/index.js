import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    Authorization: localStorage.getItem('Authorization') ? localStorage.getItem('Authorization') : null,
    Token: localStorage.getItem('Token'),
  },
  getters: {
    hasLoggedIn: (state) => {
      return state.Authorization != null
    },
    getUserID: (state) => {
      if (state.Authorization == null) {
        return null
      }
      return JSON.parse(state.Authorization).key
    },
    getUserName: (state) => {
      if (state.Authorization == null) {
        return null
      }
      return JSON.parse(state.Authorization).name
    },
  },
  mutations: {
    // 将token存入localStorage
    userLogin (state, user) {
      let time = new Date().getTime()
      let authorization = JSON.stringify({ key: user.Authorization, time: time, name: user.name })
      state.Authorization = authorization
      localStorage.setItem('Authorization', authorization)
      localStorage.setItem('Token', user.token)
    },
    // 删除token
    userLogOut (state) {
      state.Authorization = null
      state.Token = null
      localStorage.removeItem('Authorization')
      localStorage.removeItem('Token')
    },
    // 检查token是否过期
    checkLogin (state) {
      let authorization = JSON.parse(state.Authorization)
      if (authorization === null) {
        return
      }
      let expireTime = authorization.time
      let time = new Date().getTime()
      if (time - expireTime > 1000 * 60 * 60 * 2) {
        state.Authorization = null
        localStorage.removeItem('Authorization')
        localStorage.removeItem('Token')
      }
    },
    changeUserName (state, name) {
      let key = JSON.parse(state.Authorization).key
      let time = JSON.parse(state.Authorization).time
      let authorization = JSON.stringify({ key: key, time: time, name: name })
      state.Authorization = authorization
      localStorage.setItem('Authorization', authorization)
    },
  },
  actions: {
  },
  modules: {
  }
})
