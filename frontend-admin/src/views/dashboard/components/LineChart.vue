<template>
  <div class="line-chart" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons')
import resize from './mixins/resize'
import { getRecentData } from '@/api/order'

export default {
  name: 'LineChart',
  mixins: [resize],
  props: {
    autoResize: {
      type: Boolean,
      default: true
    },
  },
  data() {
    return {
      chart: null,
      num: [],
      sum: [],
      date: [],
    }
  },
  created() {
    this.getData()
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      this.setOptions()
    },
    getData() {
      getRecentData().then(response => {
        const { date, num, sum } = response
        this.date = date
        this.num = num
        this.sum = sum
        this.initChart()
      })
    },
    setOptions() {
      this.chart.setOption({
        xAxis: {
          data: this.date,
          boundaryGap: false,
          axisTick: {
            show: false
          }
        },
        grid: {
          left: 10,
          right: 10,
          bottom: 20,
          top: 30,
          containLabel: true
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        yAxis: [{
          name: '金额',
          axisTick: {
            show: false
          }
        }, {
          name: '数量',
          minInterval: 1,
          splitNumber: 5,
          axisTick: {
            show: false
          }
        }],
        legend: {
          data: ['订单数', '收益']
        },
        series: [{
          name: '订单数',
          smooth: true,
          type: 'line',
          data: this.num,
          yAxisIndex: 1,
          animationDuration: 2800,
          animationEasing: 'cubicInOut',
          itemStyle: {
            normal: {
              color: '#FF005A',
              lineStyle: {
                color: '#FF005A',
                width: 2
              }
            }
          },
        }, {
          name: '收益',
          smooth: true,
          type: 'line',
          data: this.sum,
          animationDuration: 2800,
          animationEasing: 'quadraticOut',
          itemStyle: {
            normal: {
              color: '#3888fa',
              lineStyle: {
                color: '#3888fa',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            }
          },
        }]
      })
    }
  }
}
</script>

<style scoped>
  .line-chart {
    height: calc(100vh - 400px);
  }
</style>
