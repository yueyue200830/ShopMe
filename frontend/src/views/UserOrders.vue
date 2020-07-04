<template>
  <user-product-component>
    <div class="order-title">
      我的订单
    </div>
    <div class="order-divider"/>
    <div class="order-list">
      <div class="order-div" v-for="order in orders" :key="order.id">
        <div class="order-status">
          {{order.status}}
        </div>
        <div class="order-detail-line">
          <div class="order-date">
            {{order.date}}
          </div>
          <el-divider direction="vertical"/>
          <div class="order-no">
            订单号：{{order.id}}
          </div>
          <div class="order-cost">
            支付金额：{{order.sum}}元
          </div>
          <el-divider direction="vertical"/>
          <div class="order-more" @click="showOrderDetail(order.id)">
            查看详情
          </div>
        </div>
        <div class="order-divider"/>
        <div class="product-list">
          <div class="product-div" v-for="product in order.products" :key="product.id">
            <el-image :src="product.image" class="product-img"/>
            <div class="product-title">
              {{product.title}}
            </div>
            <div class="product-price">
              {{product.price}}元 × {{product.num}}
            </div>
          </div>
        </div>
      </div>
      <div class="page-div">
        <el-pagination
            class="page"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-sizes="pageSizes"
            :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="orderTotal">
        </el-pagination>
      </div>
    </div>

  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';
  export default {
    name: 'UserOrders',
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
        orders: [],
        pageSize: 5,
        currentPage: 1,
        pageSizes: [5, 10, 20],
        orderTotal: 0,
      }
    },
    created() {
      this.getOrderNumber()
      this.getOrders()
    },
    methods: {
      showOrderDetail(orderID) {
        this.$router.push({
          path: '/user/order',
          query: {
            orderID: orderID
          }
        })
      },
      handleSizeChange(newSize) {
        this.pageSize = newSize
        this.getOrders()
      },
      handleCurrentChange(newPage) {
        this.currentPage = newPage
        this.getOrders()
      },
      getOrderNumber() {
        this.$http
          .get('/api/orderNumber', {
            params: {
              userID: this.userID
            }
          })
          .then(response => {
            this.orderTotal = response.data
          })
      },
      getOrders() {
        this.$http
          .get('/api/orders', {
            params: {
              page: this.currentPage,
              size: this.pageSize,
              userID: this.userID
            }
          })
          .then(response => {
            if (response.data.code !== 0) {
              this.$message.error('加载订单失败')
            } else {
              let orders = response.data.data
              for (let i = 0; i < orders.length; ++i) {
                if (orders[i].status === 'ordered') {
                  orders[i].status = '已下单'
                } else if (orders[i].status === 'paid') {
                  orders[i].status = '已付款'
                } else if (orders[i].status === 'finished') {
                  orders[i].status = '已完成'
                } else if (orders[i].status === 'canceled') {
                  orders[i].status = '已取消'
                }
                for (let j = 0; j < orders[i]['products'].length; ++j) {
                  orders[i]['products'][j]['image'] = '/api/productImage/' + orders[i]['products'][j]['image']
                }
              }
              this.orders = orders
            }
          })
      }
    }
  }
</script>

<style scoped>
  .order-title {
    text-align: center;
    font-size: 26px;
    padding: 10px;
    line-height: 2.5em;
    color: #666;
    background-color: white;
    margin-top: 10px;
  }

  .order-divider {
    width: 100%;
    height: 1px;
    background-color: #ddd;
  }

  .order-list {
    background-color: white;
    padding: 10px 20px;
    margin-bottom: 20px;
  }

  .order-div {
    margin: 20px 10px;
    padding: 10px 30px;
    border: 1px solid #ccc;
    color: #666;
  }

  .order-status {
    padding: 20px 20px 10px 0;
    font-size: 24px;
    line-height: 1.5em;
    margin-left: 20px;
  }

  .order-detail-line {
    display: flex;
    padding: 5px 10px 10px;
    margin-bottom: 5px;
  }

  .order-date {
    padding-left: 10px;
    margin-right: 10px;
  }

  .el-divider--vertical {
    height: auto;
  }

  .order-no {
    margin-left: 10px;
    margin-right: auto;
  }

  .order-cost {
    margin-left: auto;
    padding: 0 10px;
  }

  .order-more {
    padding: 0 10px;
  }

  .order-more:hover {
    cursor: pointer;
    color: coral;
  }

  .product-list {
    margin: 0 10px;
    padding: 20px;
  }

  .product-div {
    padding: 10px 0;
    display: flex;
  }

  .product-img {
    width: 80px;
    height: 80px;
    margin-right: 30px;
  }

  .product-title {
    text-align: center;
    line-height: 80px;
    font-size: 18px;
  }

  .product-price {
    line-height: 80px;
    margin-left: auto;
  }

  .page-div {
    display: flex;
    padding: 10px;
  }

  .page {
    margin-left: auto;
  }
</style>