<template>
  <div style="overflow-y: auto;">
    <ModalDSP :dsp="dspItem"
              :show="showModal"
              @close="closeHandler"
              @save="saveHandler"
    />
    <Button type="primary" @click="showModal = true">Add DSP</Button>
    <div>
      <p style="padding:10px;">DSP</p>
      <Table :columns="columns" :data="dsp" border no-data-text="No data">
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
import ModalDSP from "./ModalDSP";
import {mapGetters, mapMutations} from "vuex";

export default {
  components: {ModalDSP},
  data() {
    return {
      showModal: false,
      dspItem: {
        id: "",
        name: "",
        endpoint: "",
        type: "",
        qps: ""
      },
      columns: [
        {
          title: 'ID',
          key: 'id'
        },
        {
          title: 'Name',
          key: 'name'
        },
        {
          title: 'QPS',
          key: 'qps'
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
    }
  },
  methods: {
    ...mapMutations('dsp', [
        'setDSPItem',
        'clearDSPItem'
    ]),
    show(index) {
      this.setDSPItem(this.dsp[index])
      this.showModal = true
    },
    closeHandler() {
      this.clearDSPItem()
      this.showModal = false
    },
    saveHandler() {
      console.log("save")
    }
  },
  computed: {
    ...mapGetters('dsp', [
      'dsp',
    ]),
    dspData: function () {
      return []

    }
  },
  mounted() {
      this.$store.dispatch('dsp/getDSP')
  }
}
</script>

