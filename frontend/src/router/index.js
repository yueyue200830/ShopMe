import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main
  }
]

const router = new VueRouter({
  // todo: add 404 not found to path '*'
  mode: 'history',
  routes
})

export default router
