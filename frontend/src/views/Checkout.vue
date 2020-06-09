<template>
  <div>
    <user-header/>
    <div class="cart-body">
      <div class="checkout-title">
        结算
      </div>
      <div class="checkout-divider"/>
      <el-table
          ref="multipleTable"
          :data="products">
        <el-table-column width="200" align="center">
          <template slot-scope="scope">
            <el-image :src="scope.row.image" class="product-image"/>
          </template>
        </el-table-column>
        <el-table-column label="商品名称" show-overflow-tooltip>
          <template slot-scope="scope">
            <div class="product-title">
              {{ scope.row.title }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="价格" width="120">
          <template slot-scope="scope">
            <div class="cart-price">
              {{ scope.row.price }}元
            </div>
          </template>
        </el-table-column>
        <el-table-column label="数量" align="center" width="120">
          <template slot-scope="scope">
            <div>
              {{scope.row.num}}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="center" label="小计" width="150">
          <template slot-scope="scope">
            <div>
              {{ scope.row.price * scope.row.num }}元
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="money">
        <div class="sum">
          应付金额：<a class="price">{{ sum }}</a>元
        </div>
      </div>
      <div class="press-button">
        <button class="back-cart-button" @click="backToCart">
          返回购物车
        </button>
        <button class="checkout-button" @click="makeOrder">
          去结算
        </button>
      </div>
    </div>
    <user-footer/>
  </div>

</template>

<script>
  import UserFooter from '../components/UserFooter';
  import UserHeader from '../components/UserHeader';

  export default {
    name: 'Checkout',
    components: {UserHeader, UserFooter},
    computed: {
      userID() {
        return this.$store.getters.getUserID
      },
    },
    props: ['pro'],
    data() {
      return {
        products: [],
        sum: 0,
      }
    },
    created() {
      this.products = this.$route.params.products
      this.calculateSum()
    },
    methods: {
      backToCart() {
        this.$router.push('/cart')
      },
      calculateSum() {
        let s = 0
        this.products.forEach(product => {
          s += product.price * product.num
        })
        this.sum = s
      },
      makeOrder() {

      }
    }
  }
</script>

<style scoped>
  .cart-body {
    width: 1220px;
    margin: 0 auto;
    min-height: calc(100vh - 150px);
    padding: 40px 0 60px;
  }

  .checkout-title {
    text-align: center;
    font-size: 32px;
    padding: 10px;
    line-height: 2.5em;
    color: #333333;
    background-color: white;
  }

  .checkout-divider {
    width: 100%;
    height: 1px;
    background-color: #ddd;
  }

  .product-image {
    height: 100px;
    width: 100px;
  }

  .product-title {
    font-size: 18px;
  }

  .money {
    background-color: white;
    display: flex;
    padding: 10px 20px;
  }

  .sum {
    margin-left: auto;
    line-height: 50px;
    margin-right: 20px;
  }

  .price {
    color: coral;
    font-size: 32px;
    padding: 0 5px;
  }

  .press-button {
    background-color: white;
    padding: 10px 20px 20px;
    display: flex;
  }

  .back-cart-button {
    margin: 0 10px 0 auto;
    height: 50px;
    border: 1px #bbb solid;
    border-radius: 5px;
    padding: 10px 40px;
    outline: none;
    background-color: transparent;
    color: #bbb;
    font-weight: lighter;
  }

  .back-cart-button:hover {
    cursor: pointer;
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