<template>
  <div>
    <user-header/>
    <div class="cart-body">
      <div class="shopping-cart">
        <span class="shopping-cart-title">
          购物车
        </span>
      </div>
      <div class="cart-divider"/>
      <el-table
          ref="multipleTable"
          :data="products"
          @selection-change="selectProduct">
        <el-table-column align="center" type="selection" width="70">
        </el-table-column>
        <el-table-column width="160" align="center">
          <template slot-scope="scope">
            <el-image :src="scope.row.image" class="cart-image"/>
          </template>
        </el-table-column>
        <el-table-column label="商品名称" show-overflow-tooltip>
          <template slot-scope="scope">
            <div class="cart-title">
              {{ scope.row.title }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="价格" width="140">
          <template slot-scope="scope">
            <div class="cart-price">
              {{ scope.row.price }}元
            </div>
          </template>
        </el-table-column>
        <el-table-column label="数量" align="center" width="180">
          <template slot-scope="scope">
            <el-input-number
                size="small"
                v-model="scope.row.num"
                :min="1"
                :max="10"
                @change="changeProductNumber(scope.row)"
                label="商品数量"/>
          </template>
        </el-table-column>
        <el-table-column align="center" label="小计" width="120">
          <template slot-scope="scope">
            <div>
              {{ scope.row.price * scope.row.num }}元
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="100">
          <template slot-scope="scope">
            {{scope.rowKey}}
            <i class="el-icon-close cart-delete" @click="deleteProduct(scope.row)"/>
          </template>
        </el-table-column>
      </el-table>
      <div class="cart-sum">
        <div class="sum-num">
          已选{{selectedNum}}件商品
        </div>
        <div class="sum-price">
          合计：<a class="sum-price-large">{{sum}}</a>元
        </div>
        <button class="checkout-button" @click="checkout">
          结算
        </button>
      </div>
    </div>
    <user-footer/>
  </div>
</template>

<script>
  import UserHeader from '../components/UserHeader';
  import UserFooter from '../components/UserFooter';

  export default {
    name: 'ShoppingCart',
    components: {UserFooter, UserHeader},
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
        products: [],
        selectedProducts: [],
        selectedNum: 0,
        sum: 0,
      }
    },
    created() {
      // todo: do not allow visitor visit this page
      if (this.hasLoggedIn) {
        this.getCartProducts()
      }
    },
    methods: {
      selectProduct(val) {
        this.selectedProducts = val;
        this.selectedNum = this.selectedProducts.length
        this.calculateSum()
      },
      calculateSum() {
        let s = 0
        this.selectedProducts.forEach(product => {
          s += product.price * product.num
        })
        this.sum = s
      },
      getCartProducts() {
        this.$http
          .get('/api/cardProducts', {
            params: {
              id: this.userID
            }
          })
          .then(response => {
            if (response.data.code !== 0) {
              this.$message.error('加载购物车失败')
            } else {
              let products = response.data.data
              for (let i = 0; i < products.length; ++i) {
                products[i]['image'] = '/api/productImage/' + products[i]['image']
              }
              this.products = products
              // todo: check if length is 0
            }
          })
      },
      changeProductNumber(product) {
        this.$http
          .put('/api/cardProduct', {
            productID: product.productID,
            userID: product.userID,
            num: product.num,
          })
          .then(response => {
            if (response.data !== 0) {
              this.$message.error("修改失败，请重试")
              this.getCartProducts()
            } else {
              this.calculateSum()
            }
          })
      },
      deleteProduct(product) {
        this.$http
          .delete('api/cardProduct', {
            params: {
              productID: product.productID,
              userID: product.userID,
            }
          })
          .then(response => {
            if (response.data !== 0) {
              this.$message.error("删除商品失败，请重试")
              this.getCartProducts()
            } else {
              for (let i = 0; i < this.products.length; i++) {
                if (product === this.products[i]) {
                  this.products.splice(i, 1)
                }
              }
            }
          })
      },
      checkout() {
        this.$router.push({
          name: 'Checkout',
          params: {
            products: this.selectedProducts
          }
        })
      }
    }
  }
</script>

<style scoped>
  .cart-body {
    width: 1220px;
    margin: 0 auto;
    min-height: calc(100vh - 150px);
    padding: 10px 0 20px;
  }

  .shopping-cart {
    background-color: white;
    padding: 20px;
    font-size: 24px;
    margin-top: 20px;
  }

  .shopping-cart-title {
    line-height: 2em;
    margin-left: 30px;
    color: #444;
  }

  .cart-divider {
    width: 100%;
    height: 1px;
    background-color: #ddd;
  }

  .cart-image {
    height: 110px;
    width: 110px;
  }

  .cart-title {
    font-size: 18px;
  }

  .cart-delete:hover {
    cursor: pointer;
    color: coral;
  }

  .cart-sum {
    background-color: white;
    display: flex;
    padding: 20px;
  }

  .sum-num {
    padding: 10px 30px;
    margin-right: auto;
    line-height: 30px;
  }

  .sum-price {
    padding: 0 30px;
    color: coral;
  }

  .sum-price-large {
    font-size: 32px;
    line-height: 50px;
  }

  .checkout-button {
    height: 50px;
    border: 0;
    border-radius: 5px;
    padding: 10px 50px;
    background-color: coral;
    color: white;
    font-size: 18px;
    margin: 0 10px;
    outline: none;
  }

  .checkout-button:hover {
    cursor: pointer;
    background-color: #e67147;
  }
</style>