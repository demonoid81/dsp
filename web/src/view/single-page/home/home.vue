<template>
  <div>
      <canvas ref="canvas1"></canvas>
    <canvas ref="canvas2"></canvas>
    <canvas ref="canvas3"></canvas>
  </div>
</template>

<script>
import Chart from 'chart.js'
import ChartDatasourcePrometheusPlugin from 'chartjs-plugin-datasource-prometheus'
Chart.plugins.register(ChartDatasourcePrometheusPlugin);

export default {
  name: 'home',
  components: {
  },
  data () {
    return {}
  },
  mounted () {
    console.log(this.$refs)
    const myChart1 = new Chart(this.$refs['canvas1'].getContext('2d'), {
      type: 'line',
      plugins: [ChartDatasourcePrometheusPlugin],
      options: {
        plugins: {
          'datasource-prometheus': {
            prometheus: {
              endpoint: "http://162.55.244.120:9090",
              baseURL: "/api/v1",   // default value
            },
            query: 'rate(http_requests_by_feed_total{path="click"}[1m])',
            timeRange: {
              type: 'relative',
              start: -1 * 60 * 60 * 1000,
              end: 0,
            },
          },
        },
      },
    });
    const myChart2 = new Chart(this.$refs['canvas2'].getContext('2d'), {
      type: 'line',
      plugins: [ChartDatasourcePrometheusPlugin],
      options: {
        plugins: {
          'datasource-prometheus': {
            prometheus: {
              endpoint: "http://162.55.244.120:9090",
              baseURL: "/api/v1",   // default value
            },
            query: 'increase(response_status[1m])',
            timeRange: {
              type: 'relative',
              start: -1 * 60 * 60 * 1000,
              end: 0,
            },
          },
        },
      },
    });
    const myChart3 = new Chart(this.$refs['canvas3'].getContext('2d'), {
      type: 'line',
      plugins: [ChartDatasourcePrometheusPlugin],
      options: {
        plugins: {
          'datasource-prometheus': {
            prometheus: {
              endpoint: "http://162.55.244.120:9090",
              baseURL: "/api/v1",   // default value
            },
            query: 'rate(http_requests_total{job="dsp", path="/click"}[1m])',
            timeRange: {
              type: 'relative',
              start: -1 * 60 * 60 * 1000,
              end: 0,
            },
          },
        },
      },
    });
    const myChart4 = new Chart(this.$refs['canvas4'].getContext('2d'), {
      type: 'line',
      plugins: [ChartDatasourcePrometheusPlugin],
      options: {
        plugins: {
          'datasource-prometheus': {
            prometheus: {
              endpoint: "http://162.55.244.120:9090",
              baseURL: "/api/v1",   // default value
            },
            query: 'rate(http_requests_total{job="dsp", path="/feed"}[1m])',
            timeRange: {
              type: 'relative',
              start: -1 * 60 * 60 * 1000,
              end: 0,
            },
          },
        },
      },
    });
  }
}
</script>

<style lang="less">
.count-style{
  font-size: 50px;
}
</style>
