<template>
  <div>
    <user-header/>
    <product-header/>
    <div class="user-body">
      <slot/>
    </div>
    <user-footer/>
  </div>
</template>

<script>
  import UserFooter from './UserFooter';
  import UserHeader from './UserHeader';
  import ProductHeader from './ProductHeader';

  export default {
    name: 'UserProductComponent',
    components: {ProductHeader, UserHeader, UserFooter},
  }
</script>

<style scoped>
  .user-body {
    width: 1220px;
    margin: 0 auto;
    min-height: calc(100vh - 210px);
    padding: 10px 0 20px;
  }
</style>