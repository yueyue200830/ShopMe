<template>
  <el-card class="product-card" shadow="hover">
    <img :src="img" width="240px"/>
    <div class="card-description">
      <div class="card-title">
        {{name}}
      </div>
      <div class="card-price">
        {{price}}元
      </div>
    </div>
  </el-card>
</template>

<script>
  export default {
    name: "ProductCard",
    props: ["name", "price", "img"]
  }
</script>

<style scoped>
  .product-card {
    width: 280px;
    margin: 10px auto;
  }

  .product-card:hover {
    cursor: pointer;
  }

  .card-description {
    padding: 10px;
    align-items: center;
    text-align: center;
  }

  .card-title {
    font-size: 16px;
    line-height: 2em;
  }

  .card-price {
    line-height: 2em;
    font-size: 14px;
  }
</style>