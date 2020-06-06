<template>
  <div class="product-header-back">
    <el-menu
        class="product-header"
        mode="horizontal"
        text-color="#666"
        router>
      <el-menu-item index="1">全部商品</el-menu-item>
      <el-menu-item index="3">消息中心</el-menu-item>
      <el-menu-item index="4">订单管理</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
  export default {
    name: "ProductHeader"
  }
</script>

<style scoped>
  .product-header {
    width: 1220px;
    margin: 0 auto;
  }

  .product-header-back {
    width: 100%;
    background-color: white;
  }
</style>