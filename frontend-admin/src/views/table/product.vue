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
        style="width: 360px; margin-left:50px;"
      >
        <el-form-item v-if="dialogStatus === 'update'" label="ID" prop="id">
          <el-input v-model="temp.id" disabled />
        </el-form-item>
        <el-form-item label="商品名称" prop="title">
          <el-input v-model="temp.title"/>
        </el-form-item>
        <el-form-item label="价格" prop="price">
          <el-input-number v-model="temp.price" :precision="2" :step="1" :min="0.01"/>
        </el-form-item>
        <el-form-item label="库存" prop="stock">
          <el-input-number v-model="temp.stock" :min="0"/>
        </el-form-item>
        <el-form-item label="图片" prop="imageUrl">
          <el-upload
            :action="imageUploadURL"
            :on-success="handleImageUploadSuccess"
            :file-list="fileList"
            :show-file-list="false"
            list-type="picture"
          >
            <el-button size="small" type="primary">点击上传</el-button>
          </el-upload>
          <el-image
            v-if="temp.imageUrl !== ''"
            style="width: 250px;"
            :src="temp.imageUrl"
            fit="contain"
          />
        </el-form-item>
        <el-form-item label="商品类别" prop="categoryID">
          <el-select
            v-model="temp.categoryID"
            style="width: 280px;"
            placeholder="请选择对应商品类别"
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

    <el-dialog title="商品详情" :visible.sync="dialogTableVisible">
      <div class="detail-info">
        <div>
          商品ID： {{ productInfo.id }}
        </div>
        <div>
          商品名称： {{ productInfo.title }}
        </div>
        <el-upload
          :action="detailUploadURL"
          :on-success="handleDetailUploadSuccess"
          :file-list="fileList"
          :show-file-list="false"
          list-type="picture"
          class="detail-upload-button"
        >
          <el-button size="medium" plain type="primary">
            添加详情图片
          </el-button>
        </el-upload>
      </div>
      <el-table
        ref="dragTable"
        v-loading="detailLoading"
        :data="productDetails"
        row-key="order"
        border
      >
        <el-table-column align="center" label="顺序" width="65">
          <template slot-scope="{row}">
            <span>{{ row.order }}</span>
          </template>
        </el-table-column>

        <el-table-column min-width="300px" label="图片">
          <template slot-scope="{row}">
            <span>{{ row.title }}</span>
            <el-image
              style="height: 100px;"
              :src="row.detailUrl"
              fit="contain"
              :preview-src-list="row.previewList"
            />
          </template>
        </el-table-column>
        <el-table-column align="center" label="拖拽" width="80">
          <template slot-scope="{}">
            <svg-icon class="drag-handler" icon-class="drag" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="删除" width="80">
          <template slot-scope="{row}">
            <i class="el-icon-close" @click="handleDetailDelete(row)"/>
          </template>
        </el-table-column>
      </el-table>
      <span slot="footer">
        <el-button @click="handleDetailInvisible">取 消</el-button>
        <el-button type="primary" @click="handleDetailOrderChange">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import * as productAPI from '@/api/product'
import Sortable from 'sortablejs'

export default {
  name: 'product',
  data() {
    const validateImageUrl = (rule, value, callback) => {
      if (this.temp.image === '' || this.temp.imageUrl === '') {
        callback(new Error('请上传商品图'))
      } else {
        callback()
      }
    }
    const validateCategoryID = (rule, value, callback) => {
      if (this.temp.categoryID === undefined) {
        callback(new Error('请选择对应类别'))
      } else {
        callback()
      }
    }
    const validateProductTitle = (rule, value, callback) => {
      if (this.temp.title === '') {
        callback(new Error('请输入商品名称'))
      } else {
        callback()
      }
    }
    return {
      imageUploadURL: `${process.env.VUE_APP_BASE_API}/productImage`,
      detailUploadURL: `${process.env.VUE_APP_BASE_API}/productDetail/image`,
      list: null,
      total: 0,
      listLoading: true,
      detailLoading: false,
      pageSizes: [10, 20, 50],
      listQuery: {
        page: 1,
        size: 10,
      },
      categoryMap: {},
      temp: {
        id: undefined,
        title: '',
        productID: undefined,
        image: '',
        imageUrl: '',
        price: undefined,
        stock: 0,
        category: '',
        categoryID: undefined,
      },
      dialogFormVisible: false,
      dialogTableVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },
      rules: {
        title: [{ validator: validateProductTitle }],
        imageUrl: [{ validator: validateImageUrl }],
        categoryID: [{ validator: validateCategoryID }]
      },
      fileList: [],
      productDetails: [{
        productID: undefined,
        order: 1,
        title: 'test',
        detailPath: '',
        detailUrl: '',
      }],
      sortable: null,
      productInfo: {
        id: undefined,
        title: ''
      }
    }
  },
  async created() {
    // todo: sync map and list, map should get before list.
    await this.getCategoryMap()
    this.getList()
    this.getProductNumber()
  },
  methods: {
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
    async getCategoryMap() {
      await productAPI.getCategory().then(response => {
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
    resetTemp() {
      this.temp = {
        id: undefined,
        title: '',
        productID: undefined,
        image: '',
        imageUrl: '',
        price: 0.01,
        stock: 0,
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
          productAPI.createProduct(this.temp).then(response => {
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
          productAPI.updateProduct(this.temp).then(response => {
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
      const deleteQuery = {id: row.id}
      productAPI.deleteProduct(deleteQuery).then(response => {
        if (response === 0) {
          this.$message.success('删除成功')
        } else {
          this.$message.error('删除失败，请重试')
        }
        this.getList()
      })
    },
    handleImageUploadSuccess(response) {
      const {code, url} = response
      if (code !== 0) {
        this.$message.error('上传失败，请重试')
      } else {
        this.temp.image = url
        this.temp.imageUrl = process.env.VUE_APP_BASE_API + url
      }
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
      this.productInfo.id = row.id
      this.productInfo.title = row.title
      this.detailLoading = true
      this.getProductDetails()
      this.$nextTick(() => {
        this.setSort()
      })
    },
    handleDetailInvisible() {
      this.dialogTableVisible = false
      this.productDetails= [{
        productID: undefined,
        order: 1,
        title: 'test',
        detailPath: '',
        detailUrl: '',
      }]
    },
    handleDetailUploadSuccess(response) {
      const {code, url} = response
      if (code !== 0) {
        this.$message.error('上传失败，请重试')
      } else {
        this.detailLoading = true
        const l = this.productDetails.length
        this.productDetails.push({
          productID: this.productInfo.id,
          order: l + 1,
          detailPath: url,
          detailUrl: `${process.env.VUE_APP_BASE_API}/productDetail/${url}`,
        })
        productAPI.updateProductDetails(this.productDetails).then(response => {
          if (response !== 0) {
            this.$message.error("上传失败，请重试")
          }
        }).finally(() => {
          this.getProductDetails()
        })
      }
    },
    handleDetailOrderChange() {
      console.log(this.productDetails)
      if (this.productDetails.length === 0) {
        this.$message.error("请至少上传一张详情图片")
        return
      }
      productAPI.updateProductDetails(this.productDetails).then(response => {
        if (response === 0) {
          this.dialogTableVisible = false;
          this.handleDetailInvisible()
        }
      })
    },
    handleDetailDelete(row) {
      if (this.productDetails.length === 1) {
        this.$message.error("删除失败，请确保至少有一张详情图片")
        return
      }
      let i
      for (let j = 0; j < this.productDetails.length; ++j) {
        if (this.productDetails[j] === row) {
          i = j
          break
        }
      }
      this.productDetails.splice(i, 1)
      console.log(this.productDetails)
      productAPI.updateProductDetails(this.productDetails).then(response => {
        if (response !== 0) {
          this.$message.error("删除失败，请重试")
        }
      }).finally(() => {
        this.detailLoading = true
        this.getProductDetails()
      })
    },
    handleFilter() {
      // todo: handle filter
      // this.currentPage = 1
      // this.getList()
    },
    getProductDetails() {
      productAPI.getProductDetails({id: this.productInfo.id}).then(response => {
        const { code, data } = response
        if (code !== 0) {
          this.$message.error('产品详情加载失败')
        } else {
          for (let i = 0; i < data.length; i++) {
            data[i]['detailUrl'] = `${process.env.VUE_APP_BASE_API}/productDetail/${data[i]['detailPath']}`
            data[i]['previewList'] = [data[i]['detailUrl']]
          }
          this.productDetails = data
        }
      }).finally(() => {
        this.detailLoading = false
      })
    },
    setSort() {
      const el = this.$refs['dragTable'].$el.querySelectorAll('.el-table__body-wrapper > table > tbody')[0]
      this.sortable = Sortable.create(el, {
        ghostClass: 'sortable-ghost',
        dragClass: 'sortable-drag',
        onEnd: evt => {
          const targetRow = this.productDetails.splice(evt.oldIndex, 1)[0]
          this.productDetails.splice(evt.newIndex, 0, targetRow)
        },
        animation: 300
      })
    }
  }
}
</script>

<style>
  .sortable-ghost{
    background: #409EFF !important;
  }

  .sortable-drag {
    opacity: 0;
    color: white;
  }

</style>

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

  .drag-handler{
    width: 20px;
    height: 20px;
    cursor: pointer;
  }

  .detail-upload-button {
    margin: 10px 0;
  }

  .detail-info {
    padding: 10px;
    line-height: 2em;
  }

  .el-icon-close:hover {
    color: #409EFF !important;
    cursor: pointer !important;
  }
</style>
