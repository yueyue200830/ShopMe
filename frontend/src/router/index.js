import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main'
import Category from '../views/Category';
import UserLogin from '../views/UserLogin';
import UserRegister from '../views/UserRegister';

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main
  }, {
    path: '/category/:id',
    name: 'Category',
    component: Category
  }, {
    path: '/login',
    name: 'UserLogin',
    component: UserLogin
  }, {
    path: '/register',
    name: 'UserRegister',
    component: UserRegister
  }
]

const router = new VueRouter({
  // todo: add 404 not found to path '*'
  mode: 'history',
  routes
})

export default router
