<template>
  <div>
    <DatePicker type="daterange" :options="options" format="yyyy-MM-dd" placement="bottom-start"
                placeholder="Select date" style="width: 200px" :value="date"></DatePicker>
    <Button type="primary" @click="getStat">Show</Button>
    <br/>
    <br/>
    <RadioGroup v-model="filter">
    <Radio label="date">
      <span>Date</span>
    </Radio>
      <Radio label="country">
        <span>Country</span>
      </Radio>
      <Radio label="compaing">
        <span>Compaing</span>
      </Radio>
      <Radio label="os">
        <span>OS</span>
      </Radio>
      <Radio label="browser">
        <span>Browser</span>
      </Radio>
      <Radio label="sid">
        <span>SID</span>
      </Radio>
      <Radio label="feed_id">
        <span>Feed</span>
      </Radio>
    </RadioGroup>
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
          title: 'Metrica',
          key: 'metrica',
          sortable: true
        },
        {
          title: 'Requested feeds',
          key: 'req_feed',
          sortable: true
        },
        {
          title: 'Clicks',
          key: 'clicks',
          sortable: true
        },
        {
          title: 'CTR, %',
          key: 'key',
          sortable: true
        },
        {
          title: 'CPC, $',
          key: 'key',
          sortable: true
        },
        {
          title: 'Cost',
          key: 'rate',
          sortable: true
        }
      ],
      data: [],
      date: [],
      filter: 'date',
      options: {
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
  methods: {
    getStat() {
      axios
          .get(window.location.origin + '/api/stat', {
            params: {
              start: dateFormat(this.date[0], "yyyy-mm-dd"),
              end: dateFormat(this.date[1], "yyyy-mm-dd"),
              filter: this.filter
            }
          })
          .then(response => (this.data = response.data))
    }
  },
  mounted() {
    const end = new Date();
    const start = new Date();
    start.setTime(start.getTime() - 3600 * 1000 * 24);
    this.date = [start, end]
  }
}
</script>

<style scoped>

</style>
