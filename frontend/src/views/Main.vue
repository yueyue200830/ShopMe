<template>
  <div>
    <user-product-component>
      <el-carousel height="500px" class="main-carousel" trigger="click" >
        <el-carousel-item v-for="banner in banners" :key="banner">
          <el-image :src="banner"/>
        </el-carousel-item>
      </el-carousel>
      <div class="product-list" v-for="category in promote_products" :key="category">
        <div class="product-list-title">
          <div class="product-title">
            {{category.title}}
          </div>
          <div class="product-list-more">
            查看全部
          </div>
        </div>
        <el-row :gutter="20">
          <el-col :span="6" v-for="product in category.products" :key="product.name">
            <product-card v-bind="product"/>
          </el-col>
        </el-row>
      </div>
    </user-product-component>
  </div>
</template>

<script>
  import UserProductComponent from "../components/UserProductComponent";
  import ProductCard from "../components/ProductCard";
  let main_img = require('../assets/main.jpg')
  let product_img = require('../assets/product.jpg')
  export default {
    name: "Main",
    components: {ProductCard, UserProductComponent},
    data() {
      return {
        banners: [
          main_img,
          main_img,
          main_img,
          main_img
        ],
        promote_products: [{
          title: "热销产品",
          products: [{
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3799,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }]
        }, {
          title: "手机",
          products: [{
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }, {
            name: "小米10 Pro",
            price: 3999,
            img: product_img,
          }]
        }]
      }
    }
    // data() {
    //   return {
    //     users: [{
    //       id: null,
    //       name: '',
    //       email: '',
    //       avatar_id: null
    //     }]
    //   }
    // },
    // created() {
    //   this.$http
    //     .get('api/getUsers')
    //     .then(response => {
    //       this.users = response.data
    //     })
    // }
  }
</script>

<style scoped>
  .main-carousel {
    margin: 20px 0;
    border-radius: 5px;
    padding: 0 5px;
  }

  .product-list-title {
    height: 70px;
    display: flex;
    align-items: center;
  }

  .product-title {
    font-size: 24px;
    margin: auto 5px;
  }

  .product-list {
    margin: 20px 0;
  }

  .product-list-more {
    margin-left: auto;
    margin-right: 10px;
    display: flex;
    align-items: center;
    padding: 10px;
  }

  .product-list-more:hover {
    cursor: pointer;
    color: dodgerblue;
  }
</style>