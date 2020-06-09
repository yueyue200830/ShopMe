<template>
  <user-product-component>
    <div class="product-header">
      <div>
        <el-image :src="product.image" class="product-image"/>
      </div>
      <div class="product-info">
        <div class="product-title">
          {{product.title}}
        </div>
        <div class="product-price">
          {{product.price}} 元
        </div>
        <button class="add-to-cart-button" v-if="product.stock>0" @click="addToCart">
          加入购物车
        </button>
        <el-button class="cart-disable" type="info" plain disabled v-else>
          已售罄
        </el-button>
      </div>
    </div>
    <div class="product-details">
      <div class="product-detail-title">
        产品介绍
      </div>
      <div class="title-divider"/>
      <el-image
          class="detail-image"
          v-for="img in productDetails"
          :key="img.order"
          :src="img.detailPath"
          lazy/>
    </div>
  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';

  export default {
    name: 'Product',
    components: {UserProductComponent},
    computed: {
      hasLoggedIn() {
        return this.$store.getters.hasLoggedIn
      },
      userID() {
        return this.$store.getters.getUserID
      },
    },
    data() {
      return {
        productID: this.$route.params.id,
        product: {
          id: null,
          title: '',
          image: '',
          price: null,
          stock: null,
        },
        productDetails: []
      }
    },
    created() {
      document.documentElement.scrollTop = 0;
      this.getProductInfo()
      this.getProductDetails()
    },
    methods: {
      getProductInfo() {
        this.$http
          .get(`/api/product/${this.productID}`)
          .then(response => {
            if (response.data.code !== 0) {
              this.$message.error('加载产品信息失败')
              // todo: should add error page
            } else {
              let p = response.data.data
              p['image'] = '/api/productImage/' + p['image']
              this.product = p
            }
          })
      },
      getProductDetails() {
        this.$http
          .get('/api/productDetails', {
            params: {
              id: this.productID
            }
          })
          .then(response => {
            if (response.data.code !== 0) {
              this.$message.error('产品详情加载失败')
            } else {
              let details = response.data.data
              for (let i = 0; i < details.length; i++) {
                details[i]['detailPath'] = '/api/productDetail/' + details[i]['detailPath']
              }
              this.productDetails = details
            }
          })
      },
      addToCart() {
        if (!this.hasLoggedIn) {
          this.$message.warning('您尚未登录，无法加入购物车')
          return
        }
        this.$http
          .post('/api/cardProduct', {
            productID: parseInt(this.productID),
            userID: this.userID,
          })
          .then(response => {
            if (response.data === 0) {
              this.$message.success('成功添加至购物车！')
            } else {
              this.$message.error('添加失败，请重试！')
            }
          })
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
    outline: none;
  }

  .add-to-cart-button:hover {
    cursor: pointer;
    background-color: #e67147;
  }

  .cart-disable {
    height: 50px;
    padding: 10px 40px;
    font-size: 18px;
    margin: 30px 10px;
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

  .detail-image {
    width: 100%;
  }
</style>