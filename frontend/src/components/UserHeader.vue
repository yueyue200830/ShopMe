<template>
  <div class="user-header">
    <div class="user-header-inner">
      <div class="me-icon" @click="goMainPage">
        ME
      </div>
      <div class="user-header-info">
        <div class="header-button" v-if="hasLoggedIn">
          <el-dropdown @command="handleUserMenuClick">
            <span class="user-name-header">
              {{ userName }}<i class="el-icon-arrow-down el-icon--right"/>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="info">个人信息</el-dropdown-item>
              <el-dropdown-item command="password">修改密码</el-dropdown-item>
              <el-dropdown-item command="orders">我的订单</el-dropdown-item>
              <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
        <div class="user-header-info" v-else>
          <div class="header-button" @click="login">
            登录
          </div>
          <div class="header-button" @click="register">
            注册
          </div>
        </div>
        <div class="shopping-cart" @click="shoppingCart">
          购物车
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import {mapMutations} from 'vuex'

  export default {
    name: 'UserHeader',
    computed: {
      hasLoggedIn() {
        return this.$store.getters.hasLoggedIn
      },
      userID() {
        return this.$store.getters.getUserID
      },
      userName() {
        return this.$store.getters.getUserName
      }
    },
    created() {
      this.checkLogin()
    },
    methods: {
      ...mapMutations(['userLogOut', 'checkLogin']),
      login() {
        this.$router.push('/login')
      },
      register() {
        this.$router.push('/register')
      },
      shoppingCart() {
        this.$router.push('/cart')
      },
      goMainPage() {
        this.$router.push('/')
      },
      handleUserMenuClick(command) {
        if (command === "info") {
          this.$router.push('/userInfo')
        } else if (command === "password") {
          // todo: push to password
        } else if (command === "orders") {
          // todo: push to orders
        } else if (command === "logout") {
          this.userLogOut()
          // todo: push to main page if in user info page.
        }
      }
    }
  }
</script>

<style scoped>
  .user-header {
    background-color: #333;
    color: #aaa;
    height: 50px;
    padding: 5px;
  }

  .user-header-inner {
    width: 1220px;
    height: 100%;
    margin: auto;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .me-icon {
    color: #eee;
    padding: 5px 10px;
    font-size: 18px;
  }

  .me-icon:hover {
    cursor: pointer;
  }

  .user-header-info {
    display: flex;
  }

  .header-button {
    margin: auto 5px;
    padding: 5px;
  }

  .header-button:hover {
    cursor: pointer;
  }

  .shopping-cart {
    margin: auto 20px;
    padding: 10px;
    align-items: center;
  }

  .shopping-cart:hover {
    cursor: pointer;
  }

  .user-name-header {
    color: #aaa;
  }
</style>