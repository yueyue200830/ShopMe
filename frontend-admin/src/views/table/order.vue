<template>
  <div class="app-container">
    <div class="filter-container">
      <el-button type="primary" icon="el-icon-refresh-right" circle @click="handleRefresh" />
    </div>

    <el-table
      ref="productTable"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column label="ID" prop="id" align="center" width="100">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户ID" align="center" width="100">
        <template slot-scope="{row}">
          <span>{{ row.userID }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户名" min-width="150">
        <template slot-scope="{row}">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="金额" align="center" width="110">
        <template slot-scope="{row}">
          <span>{{ row.sum }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" width="110">
        <template slot-scope="{row}">
          <el-tag v-if="row.status === 'ordered'" type="warning">已下单</el-tag>
          <el-tag v-else-if="row.status === 'paid'">已付款</el-tag>
          <el-tag v-else-if="row.status === 'finished'" type="success">已完成</el-tag>
          <el-tag v-else type="info">已取消</el-tag>

        </template>
      </el-table-column>
      <el-table-column label="最后更新时间" align="center" width="180">
        <template slot-scope="{row}">
          <span>{{ getLastUpdate(row) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="订单详情" align="center" width="120">
        <template slot-scope="{row}">
          <el-button size="mini" @click="handleDetailClick(row)">
            查看
          </el-button>
        </template>
      </el-table-column>
      <el-table-column label="完成订单" align="center" width="120">
        <template slot-scope="{row}">
          <el-button :disabled="cannotFinishOrder(row)" size="mini" type="primary" @click="handleFinishOrderClick(row)">
            完成
          </el-button>
        </template>
      </el-table-column>
      <el-table-column label="取消订单" align="center" width="120" class-name="small-padding fixed-width">
        <template slot-scope="{row}">
          <el-button :disabled="cannotCancelOrder(row)" size="mini" type="danger" @click="handleCancelOrderClick(row)">
            取消
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="page-div">
      <el-pagination
        class="page"
        :current-page="listQuery.page"
        :page-sizes="pageSizes"
        :page-size="listQuery.size"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </div>

    <el-dialog title="订单详情" :visible.sync="dialogTableVisible">
      <div class="detail-info">
        <div>
          ID： {{ orderInfo.id }}
        </div>
        <div>
          用户ID： {{ orderInfo.userID }}
        </div>
        <div>
          用户名： {{ orderInfo.name }}
        </div>
        <div v-if="orderInfo.orderTime !== null">
          下单时间： {{ convertTime(orderInfo.orderTime) }}
        </div>
        <div v-if="orderInfo.payTime !== null">
          付款时间： {{ convertTime(orderInfo.payTime) }}
        </div>
        <div v-if="orderInfo.finishTime !== null">
          完成时间： {{ convertTime(orderInfo.finishTime) }}
        </div>
        <div v-if="orderInfo.cancelTime !== null">
          取消时间： {{ convertTime(orderInfo.cancelTime) }}
        </div>
      </div>
      <el-table
        ref="detailTable"
        v-loading="detailLoading"
        :data="productDetails"
        border
      >
        <el-table-column align="center" label="图片" width="130">
          <template slot-scope="{row}">
            <el-image
              style="height: 80px;"
              :src="row.image"
              fit="contain"
            />
          </template>
        </el-table-column>
        <el-table-column label="商品名称" min-width="80">
          <template slot-scope="{row}">
            <span>{{ row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="单价" width="80">
          <template slot-scope="{row}">
            <span>{{ row.price }}元</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="数量" width="80">
          <template slot-scope="{row}">
            <span>{{ row.num }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="小计" width="100">
          <template slot-scope="{row}">
            <span>{{ row.price * row.num }}元</span>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>

<script>
import * as api from '@/api/order'

export default {
  name: 'Order',
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      detailLoading: false,
      pageSizes: [10, 20, 50],
      listQuery: {
        page: 1,
        size: 10,
      },
      temp: {
        id: undefined,
        name: '',
        userID: undefined,
        orderTime: null,
        payTime: null,
        finishTime: null,
        cancelTime: null,
        status: '',
        sum: undefined,
      },
      dialogTableVisible: false,
      productDetails: [{
        image: '',
        title: '',
        sum: undefined,
        num: undefined
      }],
      orderInfo: {
        id: undefined,
        name: '',
        userID: undefined,
        orderTime: null,
        payTime: null,
        finishTime: null,
        cancelTime: null,
        status: '',
        sum: undefined,
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      api.getOrders(this.listQuery).then(response => {
        const { code, data } = response
        if (code !== 0) {
          this.$message.error('获取订单失败')
        } else {
          this.list = data.list
          this.total = data.total
        }
      }).finally(() => {
        this.listLoading = false
      })
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        name: '',
        userID: undefined,
        orderTime: null,
        payTime: null,
        finishTime: null,
        cancelTime: null,
        status: '',
        sum: undefined,
      }
    },
    handleCancelOrderClick(row) {
      this.$confirm(`确认取消订单${row.id}?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
      }).then(() => {
        this.cancelOrder(row)
      }).catch(() => {})
    },
    cancelOrder(row) {
      api.cancelOrder(row).then(response => {
        if (response !== 0) {
          this.$message.error('取消订单失败，请重试')
        }
      }).finally(() => {
        this.handleRefresh()
      })
    },
    handleSizeChange(newSize) {
      this.listQuery.size = newSize
      this.getList()
    },
    handlePageChange(newPage) {
      this.listQuery.page = newPage
      this.getList()
    },
    handleDetailClick(row) {
      this.dialogTableVisible = true
      this.orderInfo = row
      this.detailLoading = true
      this.getOrderDetail()
    },
    handleRefresh() {
      this.listLoading = true
      this.getList()
    },
    handleFinishOrderClick(row) {
      this.$confirm(`确认完成订单${row.id}?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
      }).then(() => {
        this.finishOrder(row)
      }).catch(() => {})
    },
    finishOrder(row) {
      api.finishOrder(row).then(response => {
        if (response !== 0) {
          this.$message.error('操作失败，请重试')
        }
      }).finally(() => {
        this.handleRefresh()
      })
    },
    getOrderDetail() {
      this.detailLoading = true
      api.getOrderDetail({ id: this.orderInfo.id }).then(response => {
        const { code, data } = response
        if (code !== 0) {
          this.$message.error('订单详情加载失败')
        } else {
          const products = data.products
          for (let i = 0; i < products.length; i++) {
            products[i]['image'] = `${process.env.VUE_APP_BASE_API}/productImage/${products[i]['image']}`
          }
          this.productDetails = products
        }
      }).finally(() => {
        this.detailLoading = false
      })
    },
    getLastUpdate(row) {
      let time
      if (row.finishTime !== null) {
        time = row.finishTime
      } else if (row.cancelTime !== null) {
        time = row.cancelTime
      } else if (row.payTime !== null) {
        time = row.payTime
      } else {
        time = row.orderTime
      }
      return this.convertTime(time)
    },
    convertTime(time) {
      if (time == null) {
        return null
      }
      let t = time.substr(0, 19)
      t = t.replace('T', ' ')
      return t
    },
    cannotFinishOrder(row) {
      return !(row.finishTime == null && row.cancelTime == null && row.payTime != null)
    },
    cannotCancelOrder(row) {
      return !(row.finishTime == null && row.cancelTime == null)
    }
  }
}
</script>

<style scoped>
  .filter-container {
    padding: 10px;
  }

  .filter-item {
    margin: 5px 10px;
  }

  .page-div {
    display: flex;
    margin: 10px 5px;
  }

  .page {
    padding: 10px;
    margin-left: auto;
  }

  .detail-info {
    padding: 10px;
    line-height: 2em;
  }
</style>
