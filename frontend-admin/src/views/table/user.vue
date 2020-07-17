<template>
  <div class="app-container">
    <div class="action-container">
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
      <el-table-column label="ID" prop="id" align="center" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户名" min-width="150px">
        <template slot-scope="{row}">
          <span class="link-type">{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="邮箱" width="250">
        <template slot-scope="{row}">
          <span>{{ row.email }}</span>
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
  </div>
</template>

<script>
import * as api from '@/api/user'

export default {
  name: 'User',
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      pageSizes: [10, 20, 50],
      listQuery: {
        page: 1,
        size: 10,
      },
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      api.getUsers(this.listQuery).then(response => {
        const { code, data } = response
        if (code !== 0) {
          this.$message.error('获取列表失败')
        } else {
          this.list = data.list
          this.total = data.num
        }
      }).finally(() => {
        this.listLoading = false
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
    handleRefresh() {
      this.listLoading = true
      this.getList()
    }
  }
}
</script>

<style scoped>
  .action-container {
    padding: 10px;
  }

  .page-div {
    display: flex;
    margin: 10px 5px;
  }

  .page {
    padding: 10px;
    margin-left: auto;
  }
</style>
