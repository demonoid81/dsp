<template>
  <Modal v-model="show" width="700"
         :mask-closable="false"
         @on-ok="saveEvent"
         @on-cancel="closeEvent">
    <p slot="header" style="text-align:center">
      <span>DSP info</span>
    </p>
    <Form :model="dsp" :label-width="100" label-position="left">
      <FormItem label="ID:" prop="id">
        <Input v-model="dsp.id"  placeholder="Enter ssp id..." :disabled="readOnly" @input="$emit('update:dsp.id', $event)"></Input>
      </FormItem>
      <FormItem label="Key:" prop="name">
        <Input v-model="dsp.name" placeholder="Enter ssp key..." :disabled="readOnly" @input="$emit('update:dsp.id', $event)"></Input>
      </FormItem>
      <FormItem label="Name:" prop="endpoint">
        <Input v-model="dsp.endpoint" placeholder="Enter ssp name..." :disabled="readOnly" @input="$emit('update:dsp.endpoint', $event)"></Input>
      </FormItem>
      <FormItem label="Type:" prop="type">
        <Select v-model="dsp.type" :disabled="readOnly" placeholder="Select ..." @change="$emit('update:dsp.type', $event)">
          <Option value="mainstream">mainstream</Option>
          <Option value="adult">adult</Option>
        </Select>
      </FormItem>
      <FormItem label="QPS:" prop="qps">
        <InputNumber style="width: 567px" v-model="dsp.qps" placeholder="Enter qps..." :disabled="readOnly" @on-change="$emit('update:dsp.qps', $event)"></InputNumber>
      </FormItem>
    </Form>
    <div slot="footer">
      <Button type="primary" @click="saveEvent" :style="{hidden:readOnly}">Save</Button>
      <Button type="primary" @click="closeEvent">Close</Button>
    </div>
  </Modal>
</template>
<script>
export default {
  name: 'ModalDSP',
  props: {
    dsp: {},
    readOnly: {},
    show: {},
  },
  methods: {
    saveEvent() {
      this.$emit('save')
    },
    closeEvent() {
      this.$emit('close')
    }
  }
}
</script>