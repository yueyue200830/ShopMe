<template>
  <div class="app-container">
    <div class="action-container">
      <el-button
        class="action-button"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >
        添加管理员
      </el-button>
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
          <span class="link-type">{{ row.userName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="200" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleReset(row)">
            重置密码
          </el-button>
          <el-button size="mini" type="danger" @click="deleteManager(row,$index)">
            删除
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

    <el-dialog
      title="创建管理员"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="80px"
        style="width: 360px; margin-left:50px;"
      >
        <el-form-item v-if="dialogStatus === 'update'" label="ID" prop="id">
          <el-input v-model="temp.id" disabled />
        </el-form-item>
        <el-form-item label="用户名" prop="userName">
          <el-input v-model="temp.userName" />
        </el-form-item>
        系统会自动生成初始密码，请注意保存，无法重新查看
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="createData()">
          确认
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import * as api from '@/api/manager'

export default {
  name: 'Manager',
  data() {
    const validateName = (rule, value, callback) => {
      // todo: check duplicate name
      if (this.temp.userName === '') {
        callback(new Error('请输入用户名'))
      } else {
        callback()
      }
    }
    return {
      list: null,
      total: 0,
      listLoading: true,
      pageSizes: [10, 20, 50],
      listQuery: {
        page: 1,
        size: 10,
      },
      temp: {
        id: undefined,
        userName: '',
      },
      dialogFormVisible: false,
      dialogStatus: '',
      rules: {
        userName: [{ validator: validateName }],
      },
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      api.getManagers(this.listQuery).then(response => {
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
    resetTemp() {
      this.temp = {
        id: undefined,
        userName: '',
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          api.createManager(this.temp).then(response => {
            const { code, data } = response
            if (code !== 0) {
              this.$message.error('创建失败，请重试')
            } else {
              this.$alert(data, '初始密码', {
                confirmButtonText: '确定',
                callback: () => {
                  this.dialogFormVisible = false
                  this.getList()
                }
              })
            }
          })
        }
      })
    },
    handleReset(row) {
      this.$confirm(`是否重置用户${row.userName}密码`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
      }).then(() => {
        this.resetPassword(row)
      }).catch(() => {})
    },
    resetPassword(row) {
      api.resetManagerPassword(row).then(response => {
        const { code, data } = response
        if (code !== 0) {
          this.$message.error('重置失败，请重试')
        } else {
          this.$alert(data, '重置密码', {
            confirmButtonText: '确定',
            callback: () => {
              this.dialogFormVisible = false
              this.getList()
            }
          })
        }
      })
    },
    deleteManager(row) {
      const deleteQuery = { id: row.id }
      api.deleteManager(deleteQuery).then(response => {
        if (response === 0) {
          this.$message.success('删除成功')
        } else {
          this.$message.error('删除失败，请重试')
        }
        this.getList()
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
