<template>
  <div @click="clickProductCard(id)">
    <el-card class="product-card" shadow="hover">
      <img :src="image" width="240px" height="240px"/>
      <div class="card-description">
        <div class="card-title">
          {{title}}
        </div>
        <div class="card-price">
          {{price}}元
        </div>
      </div>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: 'ProductCard',
    props: ['id', 'title', 'price', 'image'],
    methods: {
      clickProductCard(id) {
        this.$router.push(`/product/${id}`)
      }
    }
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
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .card-price {
    line-height: 2em;
    font-size: 14px;
  }
</style>