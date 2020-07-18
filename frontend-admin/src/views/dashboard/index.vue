<template>
  <div class="dashboard-container">
    <div class="dashboard-text">Hi: {{ name }}</div>
    <panel-group />
    <el-row class="dashboard-line-chart">
      <line-chart :chart-data="lineChartData" />
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PanelGroup from './components/PanelGroup'
import LineChart from '@/views/dashboard/components/LineChart'

export default {
  name: 'Dashboard',
  components: {
    LineChart,
    PanelGroup
  },
  data() {
    return {
      lineChartData: {
        expectedData: [100, 120, 161, 134, 105, 160, 165],
        actualData: [120, 82, 91, 154, 162, 140, 145]
      },
    }
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
}
</script>

<style lang="scss" scoped>
  .dashboard {
    &-container {
      margin: 30px;
    }

    &-text {
      font-size: 30px;
      line-height: 46px;
    }

    &-line-chart {
      padding: 16px 16px 0;
      margin-bottom: 32px;
      box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    }
  }

</style>
