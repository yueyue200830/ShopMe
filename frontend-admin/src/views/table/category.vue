<template>
  <div class="app-container">
    <div class="action-container">
      <el-button
        class="action-button"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >
        添加类别
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
      <el-table-column label="名称" min-width="150px">
        <template slot-scope="{row}">
          <span class="link-type">{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="产品数量" align="center" width="95">
        <template slot-scope="{row}">
          <span>{{ row.num }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="200" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="deleteProduct(row,$index)">
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
      :title="textMap[dialogStatus]"
      :visible.sync="dialogFormVisible"
    >
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="80px"
        status-icon
        style="width: 360px; margin-left:50px;"
      >
        <el-form-item v-if="dialogStatus === 'update'" label="ID" prop="id">
          <el-input v-model="temp.id" disabled />
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          确认
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import * as api from '@/api/category'

export default {
  name: 'Category',
  data() {
    const validateName = (rule, value, callback) => {
      if (this.temp.name === '') {
        callback(new Error('请输入名称'))
      } else {
        api.checkCategoryName({ id: this.temp.id, name: value }).then(response => {
          if (response !== 0) {
            callback(new Error('类别名称已存在'))
          } else {
            callback()
          }
        }).catch(() => {
          callback(new Error('类别名称检测失败'))
        })
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
        id: 0,
        name: '',
        num: undefined,
      },
      dialogFormVisible: false,
      dialogTableVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        name: [{ validator: validateName, trigger: 'blur' }],
      },
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      api.getCategories(this.listQuery).then(response => {
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
        id: 0,
        name: '',
        num: undefined,
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
          api.createCategory(this.temp).then(response => {
            if (response !== 0) {
              this.$message.error('创建失败，请重试')
            } else {
              this.dialogFormVisible = false
              this.getList()
            }
          })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          api.updateCategory(this.temp).then(response => {
            if (response !== 0) {
              this.$message.error('更新失败，请重试')
            } else {
              this.dialogFormVisible = false
              this.getList()
            }
          })
        }
      })
    },
    deleteProduct(row) {
      const deleteQuery = { id: row.id }
      api.deleteCategory(deleteQuery).then(response => {
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
