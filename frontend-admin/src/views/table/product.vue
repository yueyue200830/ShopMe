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
      <el-table-column label="价格" align="center" width="95">
        <template slot-scope="{row}">
          <span>{{ row.price }}</span>
        </template>
      </el-table-column>
      <el-table-column label="库存" align="center" width="95">
        <template slot-scope="{row}">
          <span>{{ row.stock }}</span>
        </template>
      </el-table-column>
      <el-table-column label="商品类别" align="center" width="95">
        <template slot-scope="{row}">
          <span>{{ row.category }}</span>
        </template>
      </el-table-column>
      <el-table-column label="图片" width="110px" align="center">
        <template slot-scope="{row}">
          <el-image
            style="width: 80px;"
            :src="row.imageUrl"
            fit="contain"
            :preview-src-list="row.imagePreviewList"
          />
        </template>
      </el-table-column>
      <el-table-column label="商品详情" align="center" width="110">
        <template slot-scope="{row}">
          <el-button size="mini" @click="handleDetailClick(row)">
            查看
          </el-button>
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
        <el-form-item label="商品名称" prop="title">
          <el-input v-model="temp.title"/>
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input v-model="temp.price"/>
        </el-form-item>
        <el-form-item label="库存" prop="stock">
          <el-input v-model="temp.stock"/>
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
        <el-form-item label="商品类别" prop="categoryID">
          <el-select
            v-model="temp.categoryID"
            style="width: 280px;"
            placeholder="请选择对应商品类别"
            @change="handleCategoryChange"
          >
            <el-option
              v-for="item in categoryMap"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
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
import * as productAPI from '@/api/product'

export default {
  name: 'product',
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
        callback(new Error('请选择对应产品'))
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
      categoryMap: {},
      categoryOptions: [],
      temp: {
        id: undefined,
        title: '',
        productID: undefined,
        image: '',
        price: undefined,
        stock: undefined,
        category: '',
        categoryID: undefined,
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
    this.getCategoryMap()
    this.getList()
    this.getProductNumber()
  },
  methods: {
    handleImageUploadSuccess(response) {
      const {code, url} = response
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
      productAPI.getProducts(this.listQuery).then(response => {
        const list = response.data
        for (let i = 0; i < list.length; ++i) {
          list[i].imageUrl = process.env.VUE_APP_BASE_API + '/productImage/' + list[i].image
          list[i].imagePreviewList = [list[i].imageUrl]
          list[i].category = this.categoryMap[list[i].categoryID].name
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
    getCategoryMap() {
      productAPI.getCategory().then(response => {
        const categories = response
        const categoryMap = {}
        categories.forEach(category => {
          categoryMap[category.id] = category
        })
        this.categoryMap = categoryMap
      })
    },
    getProductNumber() {
      productAPI.getProductNumber().then(response => {
        this.total = response
      })
    },
    handleCategoryChange(val) {
      // const product = this.categoryOptions[val]
      // this.temp.productID = val
      // this.temp.image = product.image
      // this.temp.price = product.price
      // this.temp.title = product.title
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        title: '',
        productID: undefined,
        image: '',
        price: undefined,
        stock: undefined,
        category: '',
        categoryID: undefined,
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
          // createBanner(this.temp).then(response => {
          //   if (response !== 0) {
          //     this.$message.error('创建失败，请重试')
          //   } else {
          //     this.dialogFormVisible = false
          //     this.getList()
          //   }
          // })
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
          // updateBanner(this.temp).then(response => {
          //   if (response !== 0) {
          //     this.$message.error('更新失败，请重试')
          //   } else {
          //     this.dialogFormVisible = false
          //     this.getList()
          //   }
          // })
        }
      })
    },
    handleDelete(row) {
      const deleteQuery = {id: row.id}
      // deleteBanner(deleteQuery).then(response => {
      //   if (response === 0) {
      //     this.$message.success('删除成功')
      //   } else {
      //     this.$message.error('删除失败，请重试')
      //   }
      //   this.getList()
      // })
    },
    handleSizeChange(newSize) {
      this.listQuery.size = newSize
      this.getList()
    },
    handleCurrentChange(newPage) {
      this.listQuery.page = newPage
      this.getList()
    },
    handleDetailClick(row) {

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