<template>
  <user-product-component>
    <div class="product-header">
      <div>
        <img src="/api/product/image/mi10pro.jpg" class="product-image"/>
      </div>
      <div class="product-info">
        <div class="product-title">
          小米10 Pro
        </div>
        <div class="product-price">
          4999 元
        </div>
        <button class="add-to-cart-button">加入购物车</button>
      </div>
    </div>
    <div class="product-details">
      <div class="product-detail-title">
        产品介绍
      </div>
      <div class="title-divider"/>
      <el-image style="width: 100%" :src="i1" lazy/>
      <el-image style="width: 100%" :src="i2" lazy/>
    </div>
  </user-product-component>
</template>

<script>
  let img1 = require('../assets/1.jpg')
  let img2 = require('../assets/2.jpg')
  import UserProductComponent from '../components/UserProductComponent';
  export default {
    name: 'Product',
    components: {UserProductComponent},
    data() {
      return {
        i1: img1,
        i2: img2
      }
    }
  }
</script>

<style scoped>
  .product-image {
    width: 400px;
    height: 400px;
    margin: 20px;
  }

  .product-header {
    background-color: white;
    display: flex;
    padding: 20px 60px;
    margin: 20px;
    border-radius: 5px;
  }

  .product-info {
    margin-left: 60px;
    padding: 20px;
  }

  .product-title {
    font-size: 28px;
    padding: 10px;
    margin: 20px 0;
  }

  .product-price {
    color: coral;
    font-size: 24px;
    margin-top: 20px;
    padding: 10px
  }

  .add-to-cart-button {
    height: 50px;
    border: 0;
    border-radius: 5px;
    padding: 10px 40px;
    background-color: coral;
    color: #eeeeee;
    font-size: 18px;
    margin: 30px 10px;
    outline:none;
  }

  .add-to-cart-button:hover {
    cursor: pointer;
    background-color: #e67147;
  }

  .product-details {
    margin: 20px;
    background-color: white;
    padding-top: 5px;
  }

  .product-detail-title {
    padding: 10px;
    margin: 10px 30px;
    font-size: 20px;
  }

  .title-divider {
    width: calc(100% - 40px);
    height: 1px;
    margin: 0 20px 10px;
    background-color: #ccc;
  }
</style>