<template>
  <div style="overflow-y: auto;">
    <ModalSSP :ssp="sspItem"
              :read-only="readOnly"
              :show="viewSSPModal"
              @close="viewSSPModal = false"
              @addDSS="addDSP"
              @save="viewSSPModal = false"
    />
        <Button type="primary" @click="addSSPShow">Add SSP</Button>

    <div>
      <p style="padding:10px;">SSP</p>
      <Table :columns="columnsSSP" :data="ssp" border no-data-text="No data">
        <template slot="action" slot-scope="{ row, index }">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">Edit</Button>
          <Button size="small" type="error" @click="remove(row.ssp_id)">Delete</Button>
        </template>
      </Table>
    </div>
  </div>

</template>


<script>
import axios from 'axios';

import ModalSSP from "./ModalSSP";
import {mapActions, mapGetters, mapMutations} from "vuex";

export default {
  components: { ModalSSP},
  data() {
    return {

      viewSSPModal: false,
      readOnly: false,
      sspItem: {
        id: "",
        key: "",
        name: "",
        type: "",
        dsp: []
      },
      columnsSSP: [
        {
          title: 'ID',
          key: 'ssp_id'
        },
        {
          title: 'Key',
          key: 'key'
        },
        {
          title: 'Name',
          key: 'ssp_name'
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
    ...mapMutations('ssp', [
        'setSSPItem'
    ]),
    ...mapActions('ssp',[
      'deleteSSP'
    ]),
    addSSPShow() {
      this.viewSSPModal = true
    },
    addDSP() {
      if (this.sspItem.hasOwnProperty('dss')) {
        this.sspItem = {
          ...this.sspItem,
          dsp: this.sspItem.dsp.push({})
        }
        return
      }
      this.sspItem = {...this.sspItem, dsp: [{}]}
    },
    show(index) {
      this.setSSPItem(this.ssp[index])
      this.readOnly = true
      this.viewSSPModal = true
    },
    remove(id) {
      this.deleteSSP(id)
    }
  },
  computed: {
    ...mapGetters('ssp', [
      'ssp',
    ]),
  },
  mounted() {
      this.$store.dispatch('ssp/getSSP')
      this.$store.dispatch('dsp/getDSP')
      this.$store.dispatch('libs/getCountries')
  }
}
</script>

