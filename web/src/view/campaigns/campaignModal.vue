<template>
  <Modal v-model="show" width="700" fullscreen
         :mask-closable="false"
         @on-ok="saveEvent"
         @on-cancel="closeEvent">
    <p slot="header" style="color:#f60;text-align:center">
      <Icon type="ios-information-circle"></Icon>
      <span>Create compaing</span>
    </p>
    <Form :model="campaign" :label-width="200" :rules="campaignValidate">
      <FormItem label="Name" prop="name">
        <Input v-model="campaign.name" placeholder="Enter something..."></Input>
      </FormItem>
      <FormItem label="Целевой веб-ресурс" prop="url">
        <Tooltip :content="campaign.url" placement="bottom" max-width="500" :delay="1000" style="width:100%">
          <Input v-model="campaign.url" placeholder="Enter something..."></Input>
        </Tooltip>
        <Tag checkable color="primary" :checked="campaign.url.includes('{SOURCE_ID}')"
             @on-change="toggleTargetItem('{SOURCE_ID}')" style="cursor: pointer;">
          {SOURCE_ID}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{CAMPAIGN_ID}')"
             @on-change="toggleTargetItem('{CAMPAIGN_ID}')" style="cursor: pointer;">
          {CAMPAIGN_ID}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{COST}')"
             @on-change="toggleTargetItem('{COST}')" style="cursor: pointer;">
          {COST}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{COUNTRY}')"
             @on-change="toggleTargetItem('{COUNTRY}')" style="cursor: pointer;">
          {COUNTRY}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{BROWSER}')"
             @on-change="toggleTargetItem('{BROWSER}')" style="cursor: pointer;">
          {BROWSER}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{OS}')" @on-change="toggleTargetItem('{OS}')"
             style="cursor: pointer;">
          {OS}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{FRESHNESS}')"
             @on-change="toggleTargetItem('{FRESHNESS}')" style="cursor: pointer;">
          {FRESHNESS}
        </Tag>
        <Tag checkable color="primary" :checked="campaign.url.includes('{FEED_ID}')"
             @on-change="toggleTargetItem('{FEED_ID}')" style="cursor: pointer;">
          {FEED_ID}
        </Tag>
      </FormItem>
      <FormItem label="Тип Push notification" prop="type">
        <CheckboxGroup v-model="campaign.type">
          <Checkbox label="classic"></Checkbox>
          <Checkbox label="inpage"></Checkbox>
        </CheckboxGroup>
      </FormItem>
    </Form>
    <Collapse simple>
      <Panel>
        Push-уведомление
        <template slot="content">
          <Form :model="campaign" :label-width="150" :rules="campaignValidate">
            <FormItem label="Заголовок" prop="type">
              <Input v-model="campaign.ad.title" maxlength="30" show-word-limit
                     placeholder="Enter something..."></Input>
            </FormItem>
            <FormItem label="Описание" prop="type">
              <Input v-model="campaign.ad.text" maxlength="45" show-word-limit placeholder="Enter something..."></Input>
            </FormItem>
          </Form>
          <Upload
              multiple
              type="drag"
              action="//jsonplaceholder.typicode.com/posts/">
            <div style="padding: 20px 0">
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>Click or drag files here to upload</p>
            </div>
          </Upload>
          <Upload
              multiple
              type="drag"
              action="//jsonplaceholder.typicode.com/posts/">
            <div style="padding: 20px 0">
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>Click or drag files here to upload</p>
            </div>
          </Upload>
        </template>
      </Panel>
      <Panel>
        Страны и цена предложения
        <template slot="content">
          <div>

          </div>
          <Form :model="campaign" :label-width="150" :rules="campaignValidate">
            <FormItem
                v-for="(item, index) in formCampaignCountries"
                :key="index"
                label="Название страны"
                :prop="item.value"
                :rules="{required: true, message: 'Country ' + index +' can not be empty', trigger: 'blur'}">
              <Row>
                <Col span="10">
                  <Select :value="item.value" @on-change="changeCampaignCountry($event, index)" filterable
                          placeholder="Select campaign country">
                    <Option v-for="item in getCountries" :value="item.value" :key="item.value">{{ item.label }}</Option>
                  </Select>
                </Col>
                <Col span="5">
                  <label class="ivu-form-item-label" style="width:100px;">CPC</label>
                  <InputNumber :max="10" :min="0" :step="0.0001" :value="item.cpc"
                               @on-change="changeCampaignCPC($event, index)" style="width:100px;"></InputNumber>
                </Col>
                <Col span="2">
                  <Button @click="removeCampaignCountry(index)">Delete</Button>
                </Col>
              </Row>
            </FormItem>
            <FormItem>
              <Row>
                <Col span="12">
                  <Button type="dashed" long @click="addCampaignCountry" icon="md-add">Add countries</Button>
                </Col>
              </Row>
            </FormItem>
          </Form>
        </template>
      </Panel>
      <Panel>
        Определение целевой аудитории
        <template slot="content">
          <Form :model="campaign" :label-width="200" :rules="campaignValidate">
            <FormItem label="OS">
              <Select :value="campaign.target.os" multiple @on-change="changeCampaignCountry($event, index)" filterable
                      placeholder="Select campaign OS">
                <Option v-for="item in os" :value="item" :key="item">{{ item }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Browser">
              <Select :value="campaign.target.browser" multiple @on-change="changeCampaignCountry($event, index)" filterable
                      placeholder="Select campaign Browser">
                <Option v-for="item in browsers" :value="item" :key="item">{{ item }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Свежесть подписки:">
              <Form inline label-position="left" :label-width="80">
                <FormItem label="c">
                  <InputNumber :max="100" :min="0" v-model="value2"></InputNumber>
                </FormItem>
                <FormItem label="по:">
                  <InputNumber :max="100" :min="0" v-model="value2"></InputNumber>
                </FormItem>
                <FormItem label="тип:" >
                  <Select v-model="model1" style="width:200px">
                    <Option v-for="item in time_range" :value="item.value" :key="item.value">{{ item.label }}</Option>
                  </Select>
                </FormItem>
              </Form>
            </FormItem>
          </Form>
        </template>
      </Panel>
      <Panel>
        График проведения кампании
        <template slot="content">
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Все</Checkbox>
            </Col>
            <Col v-for="item in 24" key="item" flex="24px">
              <div style="position: absolute;left: 30%;top: 50%;transform: translate(-50%, -50%);">
                {{item}}
              </div>
            </Col>
          </Row>

          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Понедельник</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Вторник</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Среда</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Четверг</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Пятница</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Суббота</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox v-model="single">Воскресенье</Checkbox>
            </Col>
            <Col v-for="item in 24" key="vitem" flex="24px">
              <Checkbox v-model="single"></Checkbox>
            </Col>
          </Row>
          <br/>
          <Checkbox v-model="single">Настроить период просмотра</Checkbox>
          <DatePicker type="date" :start-date="new Date(1991, 4, 14)" placeholder="Select date" style="width: 200px"></DatePicker>
          <Checkbox v-model="single">Дата окончания отсутствует</Checkbox>
        </template>
      </Panel>
      <Panel>
        Ограничение рекламы
        <template slot="content">
          <Row>
            <Col span="8">
              <p>Общая сумма бюджета, $:</p>
              <InputNumber v-model="value2" placeholder="без ограничений" style="width:200px;"></InputNumber>
            </Col>
            <Col span="8">
              <p>Ежедневный бюджет, $:</p>
              <InputNumber v-model="value2" placeholder="без ограничений" style="width:200px;"></InputNumber>
            </Col>
          </Row>
          <br/>
          <Row>
            <Col span="8">
              <p>Общее ограничение количества переходов по ссылке:</p>
              <InputNumber v-model="value2"  placeholder="без ограничений" style="width:200px;"></InputNumber>
            </Col>
            <Col span="8">
              <p>Ежедневное ограничение количества переходов по ссылке:</p>
              <InputNumber v-model="value2"  placeholder="без ограничений" style="width:200px;"></InputNumber>
            </Col>
          </Row>
        </template>
      </Panel>
      <Panel>
        Аудитория
        <template slot="content">
          <Row :gutter="50">
            <Col flex="200px">
              <p>Тип (Source):</p>
              <Select v-model="model1" style="width:200px">
                <Option v-for="item in sourceAudience" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </Col>
            <Col flex="auto">
              <p>Источники (Source):</p>
              <Input v-model="value6" type="textarea" :rows="4" placeholder="Enter something..." />
            </Col>
          </Row>
          <br/>
          <Row :gutter="50">
            <Col flex="200px">
              <p>Тип (Feed):</p>
              <Select v-model="model1" style="width:200px">
                <Option v-for="item in feedAudience" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </Col>
            <Col flex="auto">
              <p>Источники (Feed):</p>
              <Input v-model="value6" type="textarea" :rows="4" placeholder="Enter something..." />
            </Col>
          </Row>
        </template>
      </Panel>
    </Collapse>
  </Modal>
</template>

<script>

import parseUrl from 'qhttp/parse-url';
import http_parse_query from 'qhttp/http_parse_query';
import http_build_query from 'qhttp/http_build_query';
import {mapGetters, mapMutations} from "vuex";

export default {
  name: "campaignModal",
  props: {
    show: {
      type: Boolean,
      default: false
    },
  },
  data() {
    return {
      campaignValidate: {
        name: [
          {required: true, message: 'The name cannot be empty', trigger: 'blur'}],
        url: [
          {required: true, message: 'The url cannot be empty', trigger: 'blur'},
          {type: 'url', message: 'Incorrect url format', trigger: 'blur'}
        ],
        type: [
          {required: true, type: 'array', min: 1, message: 'Choose at least one type', trigger: 'change'}
        ],
      },
      urlItems: {
        source_id: "{SOURCE_ID}",
        campaign_id: "{CAMPAIGN_ID}",
        cost: "{COST}",
        country: "{COUNTRY}",
        browser: "{BROWSER}",
        os: "{OS}",
        freshness: "{FRESHNESS}",
        feed_id: "{FEED_ID}",
      },
      daysOfWeek: [
      "Понедельник",
      "Вторник",
      "Среда",
      "Четверг",
      "Пятница",
      "Суббота",
      "Воскресенье"],
      feedAudience: [
        {
          value: "blacklist_feed",
          label: "Blacklist",
        },{
          value: "whitelist_feed",
          label:"Whitelist"
        }],
      sourceAudience: [
        {
          value: "blacklist",
          label: "Blacklist",
        },{
          value: "whitelist",
          label:"Whitelist"
        }],
      time_range: [
        {
          value: 'm',
          label: 'Минуты'
        },
        {
          value: 'h',
          label: 'Часы'
        },
        {
          value: 'd',
          label: 'Дни'
        },
      ]
    }
  },
  methods: {
    ...mapMutations({
      AddCampaignCountry: 'campaigns/AddCampaignCountry',
      UpdateCampaignCountryItemCountry: 'campaigns/UpdateCampaignCountryItemCountry',
      UpdateCampaignCountryItemCPC: 'campaigns/UpdateCampaignCountryItemCPC',
      UpdateCampaignCountryItemRemove: 'campaigns/UpdateCampaignCountryItemRemove'
    }),
    saveEvent() {
      this.$emit('save')
    },
    closeEvent() {
      this.$emit('close')
    },
    toggleTargetItem(item) {
      let target_item;
      if (this.campaign.url && this.isValidUrl(this.campaign.url)) {

        const urlparams = parseUrl(this.campaign.url);
        let params = {};
        let url = "";
        url += urlparams.protocol + "//" + urlparams.hostname

        try {
          params = http_parse_query(urlparams.query)
        } catch (e) {
          this.writeLogs(e.toString())
        }

        if (urlparams.port) {
          url += ":" + urlparams.port
        }

        if (urlparams.pathname) {
          url += urlparams.pathname
        }

        if (this.campaign.url.includes(item)) {
          for (target_item in params) {
            if (params[target_item] === item) {
              delete params[target_item]
              break;
            }
          }
        } else {
          for (target_item in this.urlItems) {
            if (this.urlItems[target_item] === item) {
              params[target_item] = item
              break;
            }
          }
        }

        let query = http_build_query(params, {leave_brackets: true});
        query = query.replace(/%7B/g, '{')
        query = query.replace(/%7D/g, '}')

        if (Object.keys(params).length) {
          this.campaign.url = url + "?" + query
        } else {
          this.campaign.url = url
        }

      }
    },
    isValidUrl(string) {
      try {
        new URL(string);
      } catch (_) {
        return false;
      }
      return true;
    },
    addCampaignCountry() {
      this.AddCampaignCountry()
    },
    changeCampaignCountry(country, index) {
      this.UpdateCampaignCountryItemCountry({country, index})
    },
    changeCampaignCPC(cpc, index) {
      this.UpdateCampaignCountryItemCPC({cpc, index})
    },
    removeCampaignCountry(index) {
      this.UpdateCampaignCountryItemRemove(index)
    }
  },
  computed: {
    ...mapGetters('libs', [
      'countries',
      'os',
      'browsers'
    ]),
    ...mapGetters('campaigns', [
      'campaign',
      'campaignCountries',

    ]),
    getCountries() {
      return this.countries.filter(({value: id1}) => !this.campaignCountries.some(({value: id2}) => id2 === id1));
    },
    formCampaignCountries: {
      get: function () {
        return this.campaignCountries.map(item => {
          console.log(item)
          let c = this.countries.find(o => o.value === item.country);
          return {
            ...item,
            ...c,
          }
        })
      },
      set: function (newValue) {

      }
    }
  },
  mounted() {
    this.$store.dispatch('libs/getCountries')
    this.$store.dispatch('libs/getOS')
    this.$store.dispatch('libs/getCountries')
  }
}
</script>

<style scoped>

</style>
