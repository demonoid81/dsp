<template>
  <div>
    <DatePicker type="daterange" :options="options2" format="yyyy-MM-dd" placement="bottom-start" placeholder="Select date" style="width: 200px" :value="value2"></DatePicker>
    <br/>
    <br/>
    <Checkbox v-model="single">Country</Checkbox>
    <br/>
    <Checkbox v-model="single">Compaing</Checkbox>
    <br/>
    <Checkbox v-model="single">OS</Checkbox>
    <br/>
    <Checkbox v-model="single">Browser</Checkbox>
    <br/>
    <Checkbox v-model="single">SID</Checkbox>
    <br/>
    <br/>
    <br/>
    <Table :columns="columns" :data="data" border no-data-text="No data"></Table>
  </div>
</template>

<script>
import axios from "axios";

import dateFormat from "dateformat";

export default {
  name: "statistics",
  data() {
    return {
      columns: [
        {
          title: 'Date',
          key: 'date'
        },
        {
          title: 'Show',
          key: 'shows'
        },
        {
          title: 'Click',
          key: 'click'
        },
        {
          title: 'CTR, %',
          key: 'key'
        },
        {
          title: 'CPC, $',
          key: 'key'
        },
        {
          title: 'Cost',
          key: 'key'
        }
      ],
      data: [],
      value2: [],
      options2: {
        shortcuts: [
          {
            text: '1 week',
            value () {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              return [start, end];
            }
          },
          {
            text: '1 month',
            value () {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              return [start, end];
            }
          },
          {
            text: '3 months',
            value () {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              return [start, end];
            }
          }
        ]
      }
    }
  },
  mounted() {
    const end = new Date();
    const start = new Date();
    start.setTime(start.getTime() - 3600 * 1000 * 24);
    this.value2 = [start, end]
    axios
        .get(window.location.origin + '/api/stat', {
          params: {
            start: dateFormat(start, "yyyy-mm-dd"),
            end: dateFormat(end, "yyyy-mm-dd"),
          }
        })
        .then(response => (this.data = response.data))
  }
}
</script>

<style scoped>

</style>