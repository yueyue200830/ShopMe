<template>
  <user-product-component>
    <el-carousel height="500px" class="main-carousel" trigger="click">
      <el-carousel-item v-for="banner in banners" :key="banner.id">
        <el-image :src="banner.bannerPath"/>
      </el-carousel-item>
    </el-carousel>
    <div class="product-list" v-for="category in promoteProducts" :key="category.title">
      <div class="product-list-title">
        <div class="product-title">
          {{category.title}}
        </div>
        <div class="product-list-more" @click="showMore(category.category)">
          查看全部
        </div>
      </div>
      <el-row :gutter="20">
        <el-col :span="6" v-for="product in category.products" :key="product.id">
          <product-card v-bind="product"/>
        </el-col>
      </el-row>
    </div>
  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';
  import ProductCard from '../components/ProductCard';

  export default {
    name: 'Main',
    components: {ProductCard, UserProductComponent},
    data() {
      return {
        banners: [],
        promoteProducts: []
      }
    },
    created() {
      this.getPromoteProducts()
      this.getBanner()
    },
    methods: {
      showMore(category) {
        this.$router.push(`/category/${category}`)
      },
      getBanner() {
        this.$http
          .get('/api/mainBanners')
          .then(response => {
            let banners = response.data
            for (let i = 0; i < banners.length; i++) {
              banners[i]['bannerPath'] = '/api/banner/image/' + banners[i]['bannerPath']
            }
            this.banners = banners
          })
      },
      getPromoteProducts() {
        this.$http
          .get('/api/promoteProducts')
          .then(response => {
            let promoteProducts = response.data
            for (let i = 0; i < promoteProducts.length; i++) {
              for (let j = 0; j < promoteProducts[i]['products'].length; j++) {
                promoteProducts[i]['products'][j]['image'] = '/api/productImage/' + promoteProducts[i]['products'][j]['image']
              }
            }
            this.promoteProducts = promoteProducts
          })
      }
    }
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