<template>
  <Modal v-model="show" width="700"
         :mask-closable="false"
         @on-ok="saveEvent"
         @on-cancel="closeEvent">
    <p slot="header" style="text-align:center">
      <span>SSP</span>
    </p>
    {{ sspItem }}
    <Form :model="sspItem" :label-width="100" label-position="left">
      <FormItem label="ID:" prop="id">
        <InputNumber style="width: 567px" v-model="sspID" placeholder="Enter id..."
                     @on-change="$emit('update:ssp.id', $event)"></InputNumber>
      </FormItem>
      <FormItem label="Key:" prop="key">
        <Input v-model="sspKey" placeholder="Enter ssp key..."></Input>
      </FormItem>
      <FormItem label="Name:" prop="name">
        <Input v-model="sspName" placeholder="Enter ssp name..."></Input>
      </FormItem>
      <FormItem label="Type:" prop="type">
        <Select v-model="sspType" placeholder="Select ...">
          <Option value="mainstream">mainstream</Option>
          <Option value="adult">adult</Option>
        </Select>
      </FormItem>
      <FormItem label="DSP:" prop="dsp">
        <Button icon="ios-add" type="dashed" size="small" @click="addDSP">Add DSS</Button>
        <div v-for="(dspItem, index) in sspItem.dsp"
             style="border: 1px dashed #dcdee2; border-radius: 3px; padding: 10px; margin-top: 10px; margin-bottom: 10px">
          <Form :model="dspItem" :label-width="125" label-position="left">
            <FormItem label="Enabled:" prop="dsp_id">
              <Row type="flex" justify="space-between">
                <Col span="3">
                  <i-switch v-model="dspItem.enabled" @on-change="updateDSPItem( index, 'enabled', $event)"/>
                </Col>
                <Col span="5">
                  <Button icon="ios-remove" type="error" ghost size="small" @click="removeDSP(index)">
                    remove
                  </Button>
                </Col>
              </Row>
            </FormItem>
            <FormItem label="ID:">
              <Select :value="dspItem.id" placeholder="Select ..." style="width: 420px"
                      @on-change="updateDSPItem( index, 'id', $event)">
                <Option v-for="item in dsp" :value="item.id" :key="item.id">{{ item.id + " " + item.name }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Revshare:" prop="profit">
              <InputNumber style="width: 420px" :value="dspItem.profit" placeholder="Enter id..."
                           :formatter="value => `$ ${value}`.replace(/\B(?=(\d{4})+(?!\d))/g, ',')"
                           :parser="value => value.replace(/\$\s?|(,*)/g, '')"
                           @on-change="updateDSPItem( index, 'profit', $event)"
              ></InputNumber>
            </FormItem>
            <FormItem label="Source ID Blacklist:" prop="source_id_blacklist">
              <Input :value="dspItem.source_id_blacklist" placeholder="Enter ..." type="textarea"
                     style="width:420px" @on-change="updateDSPItem( index, 'source_id_blacklist', $event)"></Input>
            </FormItem>
            <FormItem label="Country Blacklist:" prop="country_blacklist">
              <Select :value="dspItem.country_blacklist" multiple style="width:420px" placeholder="Select ..."
                      @on-change="updateDSPItem( index, 'country_blacklist', $event)">
                <Option v-for="item in countries" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Country Whitelist:" prop="country_whitelist">
              <Select :value="dspItem.country_whitelist" multiple style="width:420px" placeholder="Select ..."
                      @on-change="updateDSPItem( index, 'country_whitelist', $event)">
                <Option v-for="item in countries" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </FormItem>
          </Form>
        </div>
      </FormItem>
    </Form>
    <div slot="footer">
      <Button type="warning" @click="saveEvent">Save</Button>
      <Button type="primary" @click="closeEvent">Close</Button>
    </div>
  </Modal>
</template>

<script>
import {mapActions, mapGetters, mapMutations} from "vuex";

const makeKey = (length) => {
  var result = '';
  var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  var charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() *
        charactersLength));
  }
  return result;
}

export default {
  name: 'ModalSSP',
  props: {
    show: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {}
  },
  methods: {
    ...mapMutations('ssp', [
      'clearCurSSP',
      'setCurSSPItem',
      'addDPSInCurSSP',
      'removeDSPinCurSSP',
      'dspItemUpdate',
    ]),
    ...mapActions('ssp', [
      'setSSP'
    ]),
    addDSP() {
      this.addDPSInCurSSP()
    },
    removeDSP(index) {
      this.removeDSPinCurSSP(index)
    },
    saveEvent() {
      this.setSSP()
      this.clearCurSSP()
      this.$emit('save')
    },
    closeEvent() {
      this.clearCurSSP()
      this.$emit('close')
    },
    updateDSPItem(index, name, value) {
      this.dspItemUpdate({index, name, value})
    }
  },
  computed: {
    ...mapGetters('ssp', [
      'ssp',
      'sspItem',
      'dspItem'
    ]),
    ...mapGetters('dsp', [
      'dsp'
    ]),
    ...mapGetters('libs', [
      'countries'
    ]),
    sspID: {
      get: function () {
        if (!this.sspItem.ssp_id) {
          let max = this.ssp.reduce(function (prev, current) {
            return (prev.ssp_id > current.ssp_id) ? prev : current
          })
          this.setCurSSPItem({
            value: max.ssp_id + 1,
            name: 'ssp_id'
          })
        }
        return this.sspItem.ssp_id
      },
      set: function (value) {
        this.setCurSSPItem({
          value,
          name: 'ssp_id'
        })
      }
    },
    sspKey: {
      get: function () {
        if (!this.sspItem.key) {
          let key = makeKey(13)
          this.setCurSSPItem({
            value: key,
            name: 'key'
          })
        }
        return this.sspItem.key
      },
      set: function (value) {
        this.setCurSSPItem({
          value,
          name: 'key'
        })
      }
    },
    sspName: {
      get: function () {
        return this.sspItem.ssp_name
      },
      set: function (value) {
        this.setCurSSPItem({
          value,
          name: 'ssp_name'
        })
      }
    },
    sspType: {
      get: function () {
        return this.sspItem.type
      },
      set: function (value) {
        this.setCurSSPItem({
          value,
          name: 'type'
        })
      }
    },
  }
}
</script>
