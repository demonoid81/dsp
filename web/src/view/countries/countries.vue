<template>
  <div>
    <p style="padding:10px;">SSP</p>
    <Table :columns="columns" :data="countries" border no-data-text="No data">
      <template slot="action" slot-scope="{ row, index }">
        <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">View</Button>
        <Button size="small" type="error" @click="remove(index)">Delete</Button>
      </template>
    </Table>
  </div>
</template>

<script>
import {mapGetters} from "vuex";

export default {
  name: "countries",
  data() {
    return{
      columnsSSP: [
        {
          title: 'ID',
          key: 'id'
        },
        {
          title: 'Key',
          key: 'key'
        },
      ]
    }
  },
  computed: {
    ...mapGetters('libs', [
      'countries',
    ]),
  }
}
</script>

<style scoped>



</style>