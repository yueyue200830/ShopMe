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
              {{product.single_price}}元 × {{product.num}}
            </div>
          </div>
        </div>
      </div>
    </div>
  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';
  export default {
    name: 'UserOrders',
    components: {UserProductComponent},
    data() {
      return {
        orders: [{
          id: 123123,
          date: '2020年3月3日',
          sum: '10000',
          status: '已下单',
          products: [{
            id: 1,
            title: '小米10 Pro',
            num: 2,
            image: '/api/productImage/mi10pro.jpg',
            single_price: 3999
          }, {
            id: 2,
            title: '小米10 Pro',
            num: 1,
            image: '/api/productImage/mi10pro.jpg',
            single_price: 4999
          }]
        }, {
          id: 12312321,
          date: '2020年3月5日',
          sum: '10000',
          status: '已完成',
          products: [{
            id: 4,
            title: '小米10 Pro',
            num: 2,
            image: '/api/productImage/mi10pro.jpg',
            single_price: 3999
          }]
        }]
      }
    },
    methods: {
      showOrderDetail(orderID) {
        this.$router.push({
          path: '/user/order',
          query: {
            orderID: orderID
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
</style>