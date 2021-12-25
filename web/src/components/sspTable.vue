<template>
  <div>
    <Row :gutter="16" style="padding: 10px">
      <Col span="6">
        <Button type="primary">Add DSP</Button>
      </Col>
      <Col span="6">
        <Button type="primary">Add SSP</Button>
      </Col>
    </Row>
    <div>
      <p style="padding:10px;">SSP</p>
      <Table :columns="columns" :data="tableData" border>
        <template slot="action" slot-scope="{ row, index }">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">View</Button>
          <Button size="small" type="error" @click="remove(index)">Delete</Button>
        </template>
      </Table>
    </div>
    <div>
      <p style="padding:10px;">DSP</p>
      <Table :columns="columns" :data="tableData" border>
        <template slot="action" slot-scope="{ row, index }">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">View</Button>
          <Button size="small" type="error" @click="remove(index)">Delete</Button>
        </template>
      </Table>
    </div>
  </div>

</template>

<script>
import axios from 'axios';


export default {
  data() {
    return {
      columns: [
        {
          title: 'ID',
          key: 'id'
        },
        {
          title: 'Key',
          key: 'key'
        },
        {
          title: 'Name',
          key: 'name'
        },
        {
          title: 'Type',
          key: 'type'
        },
        {
          title: 'Action',
          slot: 'action',
          width: 150,
          align: 'center'
        }
      ],
      data: []
    }
  },
  methods: {},
  computed: {
    tableData : function () {
      console.log(this.data)
      return this.data.map(item => {
        return {
          id: item.ssp_id,
          key: item.key,
          name: item.ssp_name,
          type: item.type
        }
      });
    }
  },
  mounted() {
    axios
        .get(window.location.origin + '/ssp/get')
        .then(response => (this.data = response.data))
  }
}
</script>

<style>

</style>