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
    component: ShoppingCart
  }, {
    path: '/checkout',
    name: 'Checkout',
    component: Checkout,
  }, {
    path: '/user/info',
    name: 'UserInfo',
    component: UserInfo
  }, {
    path: '/user/orders',
    name: 'UserOrders',
    component: UserOrders
  }, {
    path: '/user/order',
    name: 'UserOrder',
    component: Order,
  }, {
    path: '/user/password',
    name: 'UserPassword',
    component: UserPassword
  }
]

const router = new VueRouter({
  // todo: add 404 not found to path '*'
  mode: 'history',
  routes
})

export default router
