<template>
  <div class="user-header">
    <div class="user-header-inner">
      <div class="user-icon">
        ME
      </div>
      <div class="user-header-info">
        <div class="header-button" @click="login">
          登录
        </div>
        <div class="header-button" @click="register">
          注册
        </div>
        <div class="shopping-cart" @click="shoppingCart">
          购物车
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "UserHeader",
    methods: {
      login() {
        // todo: push to login
      },
      register() {
        // todo: push to register
      },
      shoppingCart() {
        // todo: push to shopping cart
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

  .user-icon {
    color: #eee;
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
</style>