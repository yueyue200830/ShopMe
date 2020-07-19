import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main'
import Category from '../views/Category';
import UserLogin from '../views/UserLogin';
import UserRegister from '../views/UserRegister';
import Product from '../views/Product';
import ShoppingCart from '../views/ShoppingCart';
import Checkout from '../views/Checkout';
import UserInfo from '../views/UserInfo';
import UserOrders from '../views/UserOrders';
import Order from '../views/Order';
import UserPassword from '../views/UserPassword';
import page404 from '../views/404';

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
  }, {
    path: '/product/:id',
    name: 'Product',
    component: Product
  }, {
    path: '/cart',
    name: 'ShoppingCart',
    component: ShoppingCart,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('Authorization') == null) {
        next('/login')
      } else {
        next()
      }
    }
  }, {
    path: '/checkout',
    name: 'Checkout',
    component: Checkout,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('Authorization') == null) {
        next('/login')
      } else {
        next()
      }
    }
  }, {
    path: '/user/info',
    name: 'UserInfo',
    component: UserInfo,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('Authorization') == null) {
        next('/login')
      } else {
        next()
      }
    }
  }, {
    path: '/user/orders',
    name: 'UserOrders',
    component: UserOrders,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('Authorization') == null) {
        next('/login')
      } else {
        next()
      }
    }
  }, {
    path: '/user/order',
    name: 'UserOrder',
    component: Order,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('Authorization') == null) {
        next('/login')
      } else {
        next()
      }
    }
  }, {
    path: '/user/password',
    name: 'UserPassword',
    component: UserPassword,
    beforeEnter: (to, from, next) => {
      if (localStorage.getItem('Authorization') == null) {
        next('/login')
      } else {
        next()
      }
    }
  }, {
    path: '/404',
    component: page404,
    hidden: true
  }, {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const router = new VueRouter({
  // mode: 'history',
  routes
})

export default router
