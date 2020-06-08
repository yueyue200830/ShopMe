<template>
  <user-product-component>
    <div class="product-list">
      <el-row :gutter="20">
        <el-col :span="6" v-for="product in products" :key="product.id">
          <product-card v-bind="product"/>
        </el-col>
      </el-row>
    </div>
    <div class="page-div">
      <el-pagination
          class="page"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="pageSizes"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="productTotal">
      </el-pagination>
    </div>
  </user-product-component>
</template>

<script>
  import UserProductComponent from '../components/UserProductComponent';
  import ProductCard from '../components/ProductCard';

  export default {
    name: 'Category',
    components: {ProductCard, UserProductComponent},
    data() {
      return {
        categoryId: this.$route.params.id,
        products: [],
        pageSize: 8,
        pageSizes: [8, 12, 16, 20],
        productTotal: 0,
        currentPage: 1,
      }
    },
    created() {
      this.getProductNumber()
      this.getProducts()
    },
    watch: {
      '$route': function (to) {
        if (to.name === "Category") {
          this.categoryId = to.params.id
          this.getProductNumber()
          this.getProducts()
        }
      }
    },
    methods: {
      handleSizeChange(newSize) {
        this.pageSize = newSize
        this.getProducts()
      },
      handleCurrentChange(newPage) {
        this.currentPage = newPage
        this.getProducts()
      },
      getProducts() {
        this.$http
          .get('/api/getProducts', {
            params: {
              page: this.currentPage,
              size: this.pageSize,
              category: this.categoryId
            }
          })
          .then(response => {
            let products = response.data
            for (let i = 0; i < products.length; i++) {
              products[i]['image'] = '/api/product/image/' + products[i]['image']
            }
            this.products = products
          })
      },
      getProductNumber() {
        this.$http
          .get('/api/getProductNumber', {
            params: {
              category: this.categoryId
            }
          })
          .then(response => {
            this.productTotal = response.data
          })
      }
    },
  }
</script>

<style scoped>
  .product-list {
    margin: 10px 0;
  }

  .page-div {
    display: flex;
    padding: 10px;
  }

  .page {
    margin-left: auto;
  }
</style>