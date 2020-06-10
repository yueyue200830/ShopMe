<template>
  <div class="product-header-back">
    <ul class="product-header">
      <li class="category-item" @click="handleMenu(0)">
        全部商品
      </li>
      <li
          class="category-item"
          v-for="category in categories"
          :key="category.id"
          @click="handleMenu(category.id)">
        {{category.name}}
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    name: 'ProductHeader',
    data() {
      return {
        categories: []
      }
    },
    created() {
      this.$http
        .get('/api/categories')
        .then(response => {
          this.categories = response.data
        })
    },
    methods: {
      handleMenu(category) {
        this.$router.push(`/category/${category}`)
      }
    }
  }
</script>

<style scoped>
  .product-header {
    width: 1220px;
    margin: 0 auto;
    height: 60px;
    align-items: center;
    text-align: center;
    padding: 0;
  }

  .product-header-back {
    width: 100%;
    background-color: white;
    min-width: 1220px;
  }

  .category-item {
    float: left;
    list-style-type: none;
    line-height: 60px;
    padding: 0 20px;
    font-size: 14px;
    color: #444;
  }

  .category-item:hover {
    cursor: pointer;
    background-color: #f5f5f5;
  }
</style>