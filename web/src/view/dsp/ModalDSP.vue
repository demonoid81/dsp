<template>
  <Modal v-model="show" width="700"
         :mask-closable="false"
         @on-ok="saveEvent"
         @on-cancel="closeEvent">
    <p slot="header" style="text-align:center">
      <span>DSP</span>
    </p>
    <Form :model="dspItem" :label-width="100" label-position="left">
      <FormItem label="ID:" prop="id">
        <Input v-model="dspID"  placeholder="Enter dsp id..."></Input>
      </FormItem>
      <FormItem label="Name:" prop="name">
        <Input v-model="dspItem.name" placeholder="Enter dsp name..."></Input>
      </FormItem>
      <FormItem label="Endpoint:" prop="endpoint">
        <Input v-model="dspItem.endpoint" placeholder="Enter dsp endpoint..." ></Input>
      </FormItem>
      <FormItem label="Type:" prop="type">
        <Select v-model="dspItem.type" placeholder="Select dsp type..." >
          <Option value="mainstream">mainstream</Option>
          <Option value="adult">adult</Option>
        </Select>
      </FormItem>
      <FormItem label="QPS:" prop="qps">
        <InputNumber style="width: 567px" v-model="dspItem.qps" placeholder="Enter dsp qps..."></InputNumber>
      </FormItem>
    </Form>
    <div slot="footer">
      <Button type="primary" @click="saveEvent">Save</Button>
      <Button type="primary" @click="closeEvent">Close</Button>
    </div>
  </Modal>
</template>

<script>
import {mapGetters, mapMutations, mapActions} from "vuex";

export default {
  name: 'ModalDSP',
  props: {
    show: {},
  },
  methods: {
    ...mapMutations('dsp', [
      'clearDSPItem',
      'setDSPItemField',
    ]),
    ...mapActions('dsp', [
        'setDSP'
    ]),
    saveEvent() {
      this.setDSP()
      this.$emit('save')
    },
    closeEvent() {
      this.clearDSPItem()
      this.$emit('close')
    }
  },
  computed: {
    ...mapGetters('dsp', [
        'dsp',
      'dspItem',
    ]),
    dspID: {
      get: function () {
        if (!this.dspItem.id) {
          let max = this.dsp.reduce(function (prev, current) {
            return (prev.id > current.id) ? prev : current
          })
          this.setDSPItemField({
            value: max.id + 1,
            name: 'id'
          })
        }
        return this.dspItem.id
      },
      set: function (value) {
        this.setDSPItemField({
          value,
          name: 'id'
        })
      }
    },
  }
}
</script>
