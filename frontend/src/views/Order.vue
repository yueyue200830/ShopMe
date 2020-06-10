<template>
  <user-product-component>
    <div class="order-body">
      <div class="order-title">
        订单详情
      </div>
<!--      <div class="order-divider"/>-->
      <div class="order-no">
        订单号：{{order.id}}
      </div>
      <el-steps :active="step" align-center class="order-step">
        <el-step title="取消" :description="order.cancelTime" v-if="hasCanceled()"/>
        <el-step title="下单" :description="order.orderTime" />
        <el-step title="付款" :description="order.paidTime" />
        <el-step title="完成" :description="order.finishTime" />
      </el-steps>
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
          <div class="product-cal-price">
            {{product.single_price * product.num}}
          </div>
        </div>
      </div>
      <div class="order-cost">
        支付金额：<a class="order-sum">{{order.sum}}</a>元
      </div>
      <div class="press-button" v-if="couldCancel()">
        <button class="cancel-order-button" @click="cancelOrder">
          取消订单
        </button>
      </div>
      <div class="press-button" v-if="couldPay()">
        <button class="pay-button" @click="payOrder">
          去支付
        </button>
      </div>
    </div>
    <el-dialog
        title="付款"
        :visible.sync="paymentVisible"
        width="350px"
        :before-close="handleClose">
      <span>请确认付款</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="paymentVisible = false">取 消</el-button>
        <el-button type="primary" @click="paidOrder">确 定</el-button>
      </span>
    </el-dialog>
  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';
  export default {
    name: 'Order',
    components: {UserProductComponent},
    data() {
      return {
        orderID: this.$route.query.orderID,
        order: {
          id: 123123,
          date: '2020年3月3日',
          orderTime: '2020年5月3日 12:23',
          paidTime: null,
          finishTime: null,
          cancelTime: null,
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
          }, {
            id: 7,
            title: '小米10 Pro',
            num: 1,
            image: '/api/productImage/mi10pro.jpg',
            single_price: 4999
          }, {
            id: 9,
            title: '小米10 Pro',
            num: 1,
            image: '/api/productImage/mi10pro.jpg',
            single_price: 4999
          }]
        },
        step: 0,
        paymentVisible: false
      }
    },
    created() {
      this.updateStep()
    },
    methods: {
      hasCanceled() {
        return this.order.cancelTime !== null
      },
      updateStep() {
        if (this.order.cancelTime != null) {
          this.step = 0
        } else if (this.order.finishTime != null) {
          this.step = 3
        } else if (this.order.paidTime != null) {
          this.step = 2
        } else if (this.order.orderTime != null) {
          this.step = 1
        } else {
          this.step = 0
        }
      },
      couldPay() {
        return this.order.orderTime != null && this.order.paidTime == null
      },
      couldCancel() {
        return this.order.cancelTime == null && this.order.finishTime == null
      },
      cancelOrder() {
        // todo: cancel order
      },
      payOrder() {
        this.paymentVisible = true
      },
      handleClose(done) {
        this.$confirm('确认取消付款？')
          .then(() => {
            done();
          })
          .catch(() => {});
      },
      paidOrder() {
        // todo: send data to backend
        this.paymentVisible = false
      }
    }
  }
</script>

<style scoped>
  .order-body {
    background-color: white;
    color: #666;
    padding: 10px 60px 40px;
    margin: 20px 0 40px;
  }

  .order-title {
    text-align: center;
    font-size: 26px;
    padding: 10px;
    line-height: 2.5em;
  }

  .order-divider {
    width: 100%;
    height: 1px;
    background-color: #ddd;
  }

  .order-step {
    margin: 40px 0;
  }

  .order-no {
    margin: 30px 20px;
    padding: 10px;
  }

  .order-cost {
    padding: 0 10px;
    margin: 20px;
    text-align: right;
  }

  .order-sum {
    color: coral;
    font-size: 32px;
    padding: 10px 5px;
  }

  .product-list {
    margin: 0 10px;
    padding: 20px;
  }

  .product-div {
    padding: 10px 0;
    display: flex;
    border-bottom: 1px solid #ddd;
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
    margin-right: 80px;
  }

  .product-cal-price {
    line-height: 80px;
    margin: 0 40px;
  }

  .press-button {
    padding: 10px 20px;
    display: flex;
  }

  .cancel-order-button {
    margin: 0 10px 0 auto;
    height: 50px;
    border: 1px #bbb solid;
    border-radius: 5px;
    padding: 10px 40px;
    outline: none;
    background-color: transparent;
    color: #bbb;
    font-weight: lighter;
    width: 150px;
  }

  .cancel-order-button:hover {
    cursor: pointer;
  }

  .pay-button {
    height: 50px;
    border: 0;
    border-radius: 5px;
    padding: 10px 45px;
    background-color: coral;
    color: white;
    font-size: 18px;
    margin: 0 10px 0 auto;
    outline: none;
    width: 150px
  }

  .pay-button:hover {
    cursor: pointer;
    background-color: #e67147;
  }
</style>