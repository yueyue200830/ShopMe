<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.title"
        placeholder="Title"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
      />
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索
      </el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >
        新增
      </el-button>
      <el-button type="primary" icon="el-icon-refresh-right" circle @click="handleRefresh" />
    </div>

    <el-table
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
      <el-table-column label="商品名称" min-width="150px">
        <template slot-scope="{row}">
          <span class="link-type">{{ row.title }}</span>
        </template>
      </el-table-column>
      <el-table-column label="售价" align="center" width="95">
        <template slot-scope="{row}">
          <span>{{ row.price }}</span>
        </template>
      </el-table-column>
      <el-table-column label="广告图" width="200px" align="center">
        <template slot-scope="{row}">
          <el-image
            style="width: 170px;"
            :src="row.bannerImage"
            fit="contain"
            :preview-src-list="row.bannerPreviewList"
          />
        </template>
      </el-table-column>
      <el-table-column label="商品图片" width="110px" align="center">
        <template slot-scope="{row}">
          <el-image
            style="width: 80px;"
            :src="row.imageUrl"
            fit="contain"
            :preview-src-list="row.imagePreviewList"
          />
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="200" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="handleDelete(row,$index)">
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
        @current-change="handleCurrentChange"
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
        style="width: 360px; margin-left:50px;"
      >
        <el-form-item v-if="dialogStatus === 'update'" label="ID" prop="id">
          <el-input v-model="temp.id" disabled />
        </el-form-item>
        <el-form-item label="图片" prop="bannerPath">
          <el-upload
            :action="bannerUploadURL"
            :on-success="handleImageUploadSuccess"
            :file-list="fileList"
            :show-file-list="false"
            list-type="picture"
          >
            <el-button size="small" type="primary">点击上传</el-button>
          </el-upload>
          <el-image
            v-if="bannerUploadImageURL !== ''"
            style="width: 250px;"
            :src="bannerUploadImageURL"
            fit="contain"
          />
        </el-form-item>
        <el-form-item label="商品名称" prop="productID">
          <el-select
            v-model="temp.productID"
            style="width: 280px;"
            placeholder="请选择对应商品名称"
            @change="handleTitleChange"
          >
            <el-option
              v-for="item in productOptions"
              :key="item.id"
              :label="item.title"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="单价" prop="price">
          <el-input v-model="temp.price" disabled />
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
import * as bannerAPI from '@/api/banner'

export default {
  name: 'Banner',
  data() {
    const validateBannerImage = (rule, value, callback) => {
      if (this.temp.bannerPath === '' || this.temp.bannerPath === null) {
        callback(new Error('请上传广告图'))
      } else {
        callback()
      }
    }
    const validateProductID = (rule, value, callback) => {
      if (this.temp.productID === undefined) {
        callback(new Error('请选择对应商品'))
      } else {
        callback()
      }
    }
    return {
      bannerUploadURL: `${process.env.VUE_APP_BASE_API}/banner/image`,
      bannerUploadImageURL: '',
      list: null,
      total: 0,
      listLoading: true,
      pageSizes: [10, 20, 50],
      listQuery: {
        page: 1,
        size: 10,
      },
      productOptions: [],
      temp: {
        id: undefined,
        title: '',
        productID: undefined,
        bannerPath: '',
        bannerImage: '',
        image: '',
        price: undefined,
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        bannerPath: [{ validator: validateBannerImage }],
        productID: [{ validator: validateProductID }]
      },
      fileList: []
    }
  },
  created() {
    this.getList()
    this.getProductList()
  },
  methods: {
    handleImageUploadSuccess(response) {
      const { code, url } = response
      if (code !== 0) {
        this.$message.error('上传失败，请重试')
      } else {
        this.bannerUploadImageURL = process.env.VUE_APP_BASE_API + url
        this.temp.bannerImage = process.env.VUE_APP_BASE_API + url
        this.temp.bannerPath = url
      }
    },
    getList() {
      this.listLoading = true
      bannerAPI.getBannerList(this.listQuery).then(response => {
        this.total = response.data.total
        const list = response.data.items
        for (let i = 0; i < list.length; ++i) {
          list[i].bannerImage = process.env.VUE_APP_BASE_API + list[i].bannerPath
          list[i].bannerPreviewList = [list[i].bannerImage]
          list[i].imageUrl = process.env.VUE_APP_BASE_API + list[i].image
          list[i].imagePreviewList = [list[i].imageUrl]
        }
        this.list = list
      }).finally(() => {
        this.listLoading = false
      })
    },
    handleFilter() {
      // todo: handle filter
      // this.currentPage = 1
      // this.getList()
    },
    getProductList() {
      bannerAPI.getProductList().then(response => {
        const list = response
        const products = {}
        for (let i = 0; i < list.length; ++i) {
          const p = list[i]
          products[p.id] = p
        }
        this.productOptions = products
      })
    },
    handleTitleChange(val) {
      const product = this.productOptions[val]
      this.temp.productID = val
      this.temp.image = product.image
      this.temp.price = product.price
      this.temp.title = product.title
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        title: '',
        productID: undefined,
        bannerPath: '',
        bannerImage: '',
        image: '',
        price: undefined,
      }
      this.bannerUploadImageURL = ''
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
          bannerAPI.createBanner(this.temp).then(response => {
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
      this.bannerUploadImageURL = this.temp.bannerImage
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          bannerAPI.updateBanner(this.temp).then(response => {
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
    handleDelete(row) {
      const deleteQuery = { id: row.id }
      bannerAPI.deleteBanner(deleteQuery).then(response => {
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
    handleCurrentChange(newPage) {
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
</style>
