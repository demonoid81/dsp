<template>
  <div style="overflow-y: auto;">
    <ModalSSP :ssp="sspItem"
              :read-only="readOnly"
              :show="viewSSPModal"
              @close="viewSSPModal = false"
              @addDSS="addDSP"
              @save="saveSSP"
    />
    <ModalDSP :dsp="dspItem"
              :read-only="readOnly"
              :show="viewDSPModal"
              @close="viewDSPModal = false"
              @save="saveDSP"
    />
    <Row :gutter="16" style="padding: 10px">
      <Col span="6">
        <Button type="primary" @click="addDSPShow" >Add DSP</Button>
      </Col>
      <Col span="6">
        <Button type="primary" @click="addSSPShow">Add SSP</Button>
      </Col>
    </Row>
    <div>
      <p style="padding:10px;">SSP</p>
      <Table :columns="columnsSSP" :data="tableSSPData" border no-data-text="No data">
        <template slot="action" slot-scope="{ row, index }">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">View</Button>
          <Button size="small" type="error" @click="remove(index)">Delete</Button>
        </template>
      </Table>
    </div>
    <div>
      <p style="padding:10px;">DSP</p>
      <Table :columns="columnsDSP" :data="dsp" border no-data-text="No data">
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

import ModalSSP from "./ModalSSP";
import ModalDSP from "./ModalDSP";

export default {
  components: {ModalDSP, ModalSSP},
  data() {
    return {

      viewSSPModal: false,
      viewDSPModal: false,
      readOnly: false,
      sspItem: {
        id:"",
        key:"",
        name:"",
        type:"",
        dsp: []
      },
      columnsSSP: [
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
      dspItem: {
        id:"",
        name:"",
        endpoint:"",
        type: "",
        qps:""
      },
      columnsDSP: [
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
      ssp: [],
      dsp:[]
    }
  },
  methods: {
    addDSPShow() {
      this.viewDSPModal = true
      this.readOnly = false
    },
    addSSPShow() {
      this.viewSSPModal = true
      this.readOnly = false
    },
    saveDSP() {
      axios.post(window.location.origin + '/dsp/add',  this.dspItem)
          .then()
          .catch(error => {
            console.error("There was an error!", error);
          });
    },
    saveSSP() {
      axios.post(window.location.origin + '/ssp/add',  this.dspItem)
          .then()
          .catch(error => {
            console.error("There was an error!", error);
          });
    },
    addDSP () {
      if (this.sspItem.hasOwnProperty('dss')) {
        this.sspItem = {
          ...this.sspItem,
          dsp: this.sspItem.dsp.push({})}
        return
      }
      this.sspItem = {...this.sspItem, dsp: [{}]}
    },
    show (index) {
      this.sspItem = {
        id:this.ssp[index].ssp_id,
        key:this.ssp[index].key,
        name:this.ssp[index].ssp_name,
        type:this.ssp[index].type,
        dsp: {
          ...this.ssp[index].dsp,
          profit: this.ssp[index].dsp.profit / 10000
        }
      }
      this.readOnly = true
      this.viewSSPModal = true
    },
  },
  computed: {
    tableSSPData : function () {
      console.log(this.data)
      return this.ssp.map(item => {
        console.log(item)
        return {
          id: item.ssp_id,
          key: item.key,
          name: item.ssp_name,
          type: item.type
        }
      });
    },
    dspData: function () {
      return []

    }
  },
  mounted() {
    axios
        .get(window.location.origin + 'api/ssp/get')
        .then(response => (this.ssp = response.data))

    axios
        .get(window.location.origin + 'api/dsp/get')
        .then(response => (this.dsp = response.data))
  }
}
</script>

