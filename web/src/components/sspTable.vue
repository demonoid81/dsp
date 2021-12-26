<template>
  <div>
    <ModalSSP :countries="countries" :read-only="readOnly" :item="sspItem" :view="viewSSPModal"/>
    <Modal v-model="viewDSPModal" width="700">
      <p slot="header" style="text-align:center">
        <span>DSP info</span>
      </p>
      <Form :model="dspItem" :label-width="100" label-position="left">
        <FormItem label="ID:" prop="id">
          <Input v-model="dspItem.id" placeholder="Enter ssp id..." :disabled="readOnly"></Input>
        </FormItem>
        <FormItem label="Key:" prop="name">
          <Input v-model="dspItem.name" placeholder="Enter ssp key..." :disabled="readOnly"></Input>
        </FormItem>
        <FormItem label="Name:" prop="endpoint">
          <Input v-model="dspItem.endpoint" placeholder="Enter ssp name..." :disabled="readOnly"></Input>
        </FormItem>
        <FormItem label="Type:" prop="type">
          <Input v-model="dspItem.name" placeholder="Enter ssp name..." :disabled="readOnly"></Input>
        </FormItem>
        <FormItem label="Type:" prop="qps">
          <Input v-model="dspItem.qps" placeholder="Enter ssp name..." :disabled="readOnly"></Input>
        </FormItem>
      </Form>
      <div slot="footer">
        <Button type="primary" @click="viewSSPModal=false">Close</Button>
      </div>
    </Modal>
    <Row :gutter="16" style="padding: 10px">
      <Col span="6">
        <Button type="primary" @click="viewSSPModal=true" >Add DSP</Button>
      </Col>
      <Col span="6">
        <Button type="primary" @click="viewSSPModal=true">Add SSP</Button>
      </Col>
    </Row>
    <div>
      <p style="padding:10px;">SSP</p>
      <Table :columns="columnsSSP" :data="tableSSPData" border>
        <template slot="action" slot-scope="{ row, index }">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">View</Button>
          <Button size="small" type="error" @click="remove(index)">Delete</Button>
        </template>
      </Table>
    </div>
    <div>
      <p style="padding:10px;">DSP</p>
      <Table :columns="columnsDSP" :data="dspData" border>
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

export default {
  components: {ModalSSP},
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
      data: []
    }
  },
  methods: {
    show (index) {
      this.sspItem = {
        id:this.data[index].ssp_id,
        key:this.data[index].key,
        name:this.data[index].ssp_name,
        type:this.data[index].type,
        dsp: this.data[index].dsp
      }
      this.readOnly = true
      this.viewSSPModal = true
    },
  },
  computed: {
    tableSSPData : function () {
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

