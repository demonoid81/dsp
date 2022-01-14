<template>
  <Modal v-model="show" width="700"
         :mask-closable="false"
         @on-ok="saveEvent"
         @on-cancel="closeEvent">
    <p slot="header" style="text-align:center">
      <span>SSP info</span>
    </p>
    <Form :model="ssp" :label-width="100" label-position="left">
      <FormItem label="ID:" prop="id">
        <InputNumber style="width: 567px" v-model="ssp.id" placeholder="Enter id..." :disabled="readOnly" @on-change="$emit('update:ssp.id', $event)"></InputNumber>
      </FormItem>
      <FormItem label="Key:" prop="key">
        <Input v-model="ssp.key" placeholder="Enter ssp key..." :disabled="readOnly"></Input>
      </FormItem>
      <FormItem label="Name:" prop="name">
        <Input v-model="ssp.name" placeholder="Enter ssp name..." :disabled="readOnly"></Input>
      </FormItem>
      <FormItem label="Type:" prop="type">
        <Select v-model="ssp.type" :disabled="readOnly" placeholder="Select ...">
          <Option value="mainstream">mainstream</Option>
          <Option value="adult">adult</Option>
        </Select>
      </FormItem>
      <FormItem label="DSP:" prop="dsp">
        <Button icon="ios-add" type="dashed" size="small" @click="addDSS" :disabled="readOnly">Add DSS</Button>
        <div v-for="(dsp, index) in ssp.dsp"
             style="border: 1px dashed #dcdee2; border-radius: 3px; padding: 10px; margin-top: 10px; margin-bottom: 10px">
          <Form :model="dsp" :label-width="125" label-position="left">
            <FormItem label="Enabled:" prop="dsp_id">
              <Row type="flex" justify="space-between">
                <Col span="3">
                  <i-switch v-model="dsp.enabled" :disabled="readOnly"/>
                </Col>
                <Col span="5">
                  <Button icon="ios-remove" type="error" ghost size="small" @click="removeDSP" :disabled="readOnly">
                    remove
                  </Button>
                </Col>
              </Row>
            </FormItem>
            <FormItem label="ID:" prop="dsp_id">
              <InputNumber style="width: 420px" v-model="dsp.dsp_id" placeholder="Enter id..." :disabled="readOnly"></InputNumber>
            </FormItem>
            <FormItem label="Revshare:" prop="profit">
              <InputNumber style="width: 420px" v-model="dsp.profit" placeholder="Enter id..." :disabled="readOnly"
                           :formatter="value => `$ ${value}`.replace(/\B(?=(\d{4})+(?!\d))/g, ',')"
                           :parser="value => value.replace(/\$\s?|(,*)/g, '')"></InputNumber>
            </FormItem>
            <FormItem label="Source ID Blacklist:" prop="source_id_blacklist">
              <Input v-model="dsp.source_id_blacklist" placeholder="Enter ..." type="textarea" :disabled="readOnly"
                     style="width:420px"></Input>
            </FormItem>
            <FormItem label="Country Blacklist:" prop="country_blacklist">
              <Select v-model="dsp.country_blacklist" multiple style="width:420px" placeholder="Select ..."
                      :disabled="readOnly">
                <Option v-for="item in countries" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Country Whitelist:" prop="country_whitelist">
              <Select v-model="dsp.country_whitelist" multiple style="width:420px" placeholder="Select ..."
                      :disabled="readOnly">
                <Option v-for="item in countries" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </FormItem>
          </Form>
        </div>
      </FormItem>
    </Form>
    <div slot="footer">
      <Button type="primary" @click="closeEvent">Close</Button>
    </div>
  </Modal>
</template>
<script>
import countries from "../../libs/countries";

export default {
  name: 'ModalSSP',
  props: {
    readOnly: {
      type: Boolean,
      default: false
    },
    ssp: Object,
    show: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
    countries: countries,
    }
  },
  methods: {
    addDSS() {
      this.$emit('addDSS')
    },
    saveEvent() {
      this.$emit('save')
    },
    closeEvent() {
      console.log("close")
      this.$emit('close')
    }
  }
}
</script>
