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
        @click="handleFilter">
        搜索
      </el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate">
        新增
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;">
      <el-table-column label="ID" prop="id" align="center" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="产品名称" min-width="150px">
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
      <el-table-column label="产品图片" width="110px" align="center">
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
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="pageSizes"
        :page-size="listQuery.size"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
      </el-pagination>
    </div>

    <el-dialog
      :title="textMap[dialogStatus]"
      :visible.sync="dialogFormVisible">
      <el-form ref="dataForm"
               :rules="rules"
               :model="temp"
               label-position="left"
               label-width="70px"
               style="width: 350px; margin-left:50px;">
        <el-form-item label="ID" prop="id">
          <el-input v-model="temp.id" disabled/>
        </el-form-item>
        <el-form-item label="图片" prop="banner">
          <el-upload
            :action="bannerUploadURL"
            :on-success="handleImageUploadSuccess"
            :file-list="fileList"
            :show-file-list="false"
            list-type="picture">
            <el-button size="small" type="primary">点击上传</el-button>
          </el-upload>
          <el-image
            style="width: 250px;"
            :src="bannerUploadImageURL"
            fit="contain"
            v-if="bannerUploadImageURL !== ''"
          />
        </el-form-item>
        <el-form-item label="商品名称">
          <el-select
            v-model="temp.productID"
            @change="handleTitleChange"
            style="width: 280px;"
            placeholder="请选择对应产品名称">
            <el-option
              v-for="item in productOptions"
              :key="item.id"
              :label="item.title"
              :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="单价" prop="price">
          <el-input v-model="temp.price" disabled/>
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
import { getBannerList, deleteBanner, getProductList, updateBanner } from '@/api/banner'

export default {
  name: 'Banner',
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    },
  },
  data() {
    return {
      bannerUploadURL: `${process.env.VUE_APP_BASE_API}/banner/image`,
      bannerUploadImageURL: '',
      srcList: [process.env.VUE_APP_BASE_API + '/banner/image/mi10pro.jpg'],
      list: null,
      total: 0,
      listLoading: true,
      currentPage: 1,
      pageSizes: [10, 20, 50],
      listQuery: {
        page: 1,
        size: 10,
      },
      statusOptions: ['published', 'draft', 'deleted'],
      productOptions: [],
      temp: {
        id: undefined,
        // timestamp: new Date(),
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
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        // timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
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
      console.log(response)
      console.log(this.fileList)
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
      getBannerList(this.listQuery).then(response => {
        this.total = response.data.total
        let list =  response.data.items
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
      this.currentPage = 1
      this.getList()
    },
    getProductList() {
      getProductList().then(response => {
        // this.productOptions = response
        let list = response
        let products = {}
        for (let i = 0; i < list.length; ++i) {
          let p = list[i]
          products[p.id] = p
        }
        this.productOptions = products
      })
    },
    handleTitleChange(val) {
      let product = this.productOptions[val]
      this.temp.productID = val
      this.temp.image = product.image
      this.temp.price = product.price
      this.temp.title = product.title
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        // timestamp: new Date(),
        title: '',
        type: ''
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
          console.log(this.temp)
          // this.temp.id = parseInt(Math.random() * 100) + 1024 // mock a id
          // this.temp.author = 'vue-element-admin'
          updateBanner(this.temp).then(response => {
            console.log(response)
            // this.list.unshift(this.temp)
            // this.dialogFormVisible = false
            // this.$notify({
            //   title: 'Success',
            //   message: 'Created Successfully',
            //   type: 'success',
            //   duration: 2000
            // })
          })
        }
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
      // this.temp.timestamp = new Date(this.temp.timestamp)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          console.log(this.temp)
          // const tempData = Object.assign({}, this.temp)
          // tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          // updateArticle(tempData).then(() => {
          //   const index = this.list.findIndex(v => v.id === this.temp.id)
          //   this.list.splice(index, 1, this.temp)
          //   this.dialogFormVisible = false
          //   this.$notify({
          //     title: 'Success',
          //     message: 'Update Successfully',
          //     type: 'success',
          //     duration: 2000
          //   })
          // })
        }
      })
    },
    handleDelete(row, index) {
      let deleteQuery = {id: row.id}
      deleteBanner(deleteQuery).then(response => {
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
      this.currentPage = newPage
      this.getList()
    },
  }
}
</script>

<style scoped>

</style>
