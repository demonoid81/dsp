<template>
  <Modal v-model="show" width="700" fullscreen
         :mask-closable="false"
         @on-ok="saveEvent"
         @on-cancel="closeEvent">
    <p slot="header" style="color:#f60;text-align:center">
      <Icon type="ios-information-circle"></Icon>
      <span>Create compaing {{ campaignId }}</span>
    </p>
    {{$store.state.campaigns}}
    <Form :model="campaign" :label-width="200" :rules="campaignValidate">
      <FormItem label="Name" prop="name">
        <Input v-model="campaignName" placeholder="Enter something..."></Input>
      </FormItem>
      <FormItem label="Целевой веб-ресурс" prop="url">
        <Tooltip :content="campaignURL" placement="bottom" max-width="500" :delay="1000" style="width:100%">
          <Input v-model="campaignURL" placeholder="Enter something..."></Input>
        </Tooltip>
        <Tag checkable color="primary" :checked="campaignURL.includes('{SOURCE_ID}')"
             @on-change="toggleTargetItem('{SOURCE_ID}')" style="cursor: pointer;">
          {SOURCE_ID}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{CAMPAIGN_ID}')"
             @on-change="toggleTargetItem('{CAMPAIGN_ID}')" style="cursor: pointer;">
          {CAMPAIGN_ID}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{COST}')"
             @on-change="toggleTargetItem('{COST}')" style="cursor: pointer;">
          {COST}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{COUNTRY}')"
             @on-change="toggleTargetItem('{COUNTRY}')" style="cursor: pointer;">
          {COUNTRY}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{BROWSER}')"
             @on-change="toggleTargetItem('{BROWSER}')" style="cursor: pointer;">
          {BROWSER}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{OS}')" @on-change="toggleTargetItem('{OS}')"
             style="cursor: pointer;">
          {OS}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{FRESHNESS}')"
             @on-change="toggleTargetItem('{FRESHNESS}')" style="cursor: pointer;">
          {FRESHNESS}
        </Tag>
        <Tag checkable color="primary" :checked="campaignURL.includes('{FEED_ID}')"
             @on-change="toggleTargetItem('{FEED_ID}')" style="cursor: pointer;">
          {FEED_ID}
        </Tag>
      </FormItem>
      <FormItem label="Тип Push notification" prop="type">
        <CheckboxGroup v-model="campaignType">
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
            <FormItem label="Заголовок" prop="title">
              <Input v-model="campaignTitle" maxlength="30" show-word-limit
                     placeholder="Enter something..."></Input>
            </FormItem>
            <FormItem label="Описание" prop="text">
              <Input v-model="campaignText" maxlength="45" show-word-limit placeholder="Enter something..."></Input>
            </FormItem>
          </Form>
          <Row>
            <Col flex="100px">
              Иконка:
            </Col>
            <Col>
              <Upload
                  multiple
                  type="drag"
                  action="//jsonplaceholder.typicode.com/posts/">
                <div style="padding: 20px 0">
                  <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                  <p>Click or drag files here to upload</p>
                </div>
              </Upload>
            </Col>
            <Col flex="150px">
              Изображение
            </Col>
            <Col>
              <Upload
                  multiple
                  type="drag"
                  action="//jsonplaceholder.typicode.com/posts/">
                <div style="padding: 20px 0">
                  <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                  <p>Click or drag files here to upload</p>
                </div>
              </Upload>
            </Col>
          </Row>
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
              <Select :value="campaignTargetOS" multiple @on-change="setTargetOs($event)" filterable
                      placeholder="All">
                <Option v-for="item in os" :value="item" :key="item">{{ item }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Browser">
              <Select :value="campaignTargetBrowsers" multiple @on-change="setTargetBrowsers($event)" filterable
                      placeholder="All">
                <Option v-for="item in browsers" :value="item" :key="item">{{ item }}</Option>
              </Select>
            </FormItem>
            <FormItem label="Свежесть подписки:">
              <Form inline label-position="left" :label-width="80">
                <FormItem label="c">
                  <InputNumber :max="100" :min="0" v-model="campaignFreshnessInterval1"></InputNumber>
                </FormItem>
                <FormItem label="по:">
                  <InputNumber :max="100" :min="0" v-model="campaignFreshnessInterval2"></InputNumber>
                </FormItem>
                <FormItem label="тип:">
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
              <Checkbox v-model="alldays" @on-change="setAllDays($event)">Все</Checkbox>
            </Col>
            <Col v-for="item in 24" :key="'all_'+item" flex="24px">
              <div style="position: absolute;left: 30%;top: 50%;transform: translate(-50%, -50%);">
                {{ item }}
              </div>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Mon.length === 24" @on-change="setDay('Mon')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Mon.includes(index)" @on-change="setHour('Mon', index)"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Tue.length === 24" @on-change="setDay('Tue')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Tue.includes(index)" @on-change="setHour('Tue', index)"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Wed.length === 24" @on-change="setDay('Wed')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Wed.includes(index)" @on-change="setHour('Wed', index)"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Thu.length === 24" @on-change="setDay('Thu')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Thu.includes(index)" @on-change="setHour('Thu', index)"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Fri.length === 24" @on-change="setDay('Fri')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Fri.includes(index)" @on-change="setHour('Fri', index)"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Sat.length === 24" @on-change="setDay('Sat')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Sat.includes(index)" @on-change="setHour('Sat', index)"></Checkbox>
            </Col>
          </Row>
          <Row>
            <Col flex="150px">
              <Checkbox :value="schedule.Sun.length === 24" @on-change="setDay('Sun')">Понедельник</Checkbox>
            </Col>
            <Col v-for="(n,index) in 24" flex="24px" :key="'mon_'+n">
              <Checkbox :value="schedule.Sun.includes(index)" @on-change="setHour('Sun', index)"></Checkbox>
            </Col>
          </Row>
          <br/>
          <!--          <Checkbox v-model="single">Настроить период просмотра</Checkbox>-->
          <!--          <DatePicker type="date" :start-date="new Date(1991, 4, 14)" placeholder="Select date"-->
          <!--                      style="width: 200px"></DatePicker>-->
          <!--          <Checkbox v-model="single">Дата окончания отсутствует</Checkbox>-->
        </template>
      </Panel>
      <Panel>
        Ограничение рекламы
        <template slot="content">
          <Row>
            <Col span="8">
              <p>Общая сумма бюджета, $:</p>
              <InputNumber v-model="campaignLimitsBudgetTotal" placeholder="без ограничений"
                           style="width:200px;"></InputNumber>
            </Col>
            <Col span="8">
              <p>Ежедневный бюджет, $:</p>
              <InputNumber v-model="campaignLimitsBudgetDaily" placeholder="без ограничений"
                           style="width:200px;"></InputNumber>
            </Col>
          </Row>
          <br/>
          <Row>
            <Col span="8">
              <p>Общее ограничение количества переходов по ссылке:</p>
              <InputNumber v-model="campaignLimitsClickTotal" placeholder="без ограничений"
                           style="width:200px;"></InputNumber>
            </Col>
            <Col span="8">
              <p>Ежедневное ограничение количества переходов по ссылке:</p>
              <InputNumber v-model="campaignLimitsClickTotal" placeholder="без ограничений"
                           style="width:200px;"></InputNumber>
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
              <Select v-model="sourceAudienceType" style="width:200px">
                <Option v-for="item in sourceAudience" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </Col>
            <Col flex="auto">
              <p>Источники (Source):</p>
              <Input v-model="sourceAudienceValue" type="textarea" :rows="4" placeholder="Enter something..."/>
            </Col>
          </Row>
          <br/>
          <Row :gutter="50">
            <Col flex="200px">
              <p>Тип (Feed):</p>
              <Select v-model="feedAudienceType" style="width:200px">
                <Option v-for="item in feedAudience" :value="item.value" :key="item.value">{{ item.label }}</Option>
              </Select>
            </Col>
            <Col flex="auto">
              <p>Источники (Feed):</p>
              <Select v-model="feedAudienceValue" multiple :disabled="(!campaign.blacklist_feed) && (!campaign.whitelist_feed)">
                <Option v-for="item in dsp" :value="item.id" :key="item.id">{{ item.id + ' - ' + item.name }}</Option>
              </Select>
            </Col>
          </Row>
        </template>
      </Panel>
    </Collapse>
    <br/>
    <Checkbox :value="campaignActive">Запустить рекламную кампанию сразу после модерации</Checkbox>
  </Modal>
</template>

<script>

import parseUrl from 'qhttp/parse-url';
import http_parse_query from 'qhttp/http_parse_query';
import http_build_query from 'qhttp/http_build_query';
import {mapActions, mapGetters, mapMutations} from "vuex";

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
      model1: null,
      value2: null,
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
        title: [
          {required: true, message: 'The name cannot be empty', trigger: 'blur'}],
        text: [
          {required: true, message: 'The name cannot be empty', trigger: 'blur'}],
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
        }, {
          value: "whitelist_feed",
          label: "Whitelist"
        }],
      sourceAudience: [
        {
          value: "blacklist",
          label: "Blacklist",
        }, {
          value: "whitelist",
          label: "Whitelist"
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
      ],
      alldays: true,
      schedule: {
        Mon: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23],
        Tue: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23],
        Wed: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23],
        Thu: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23],
        Fri: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23],
        Sat: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23],
        Sun: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23]
      },
    }
  },
  methods: {
    ...mapMutations('campaigns', [
      'addCampaignCountry',
      'updateCampaignCountryItemCountry',
      'updateCampaignCountryItemCPC',
      'updateCampaignCountryItemRemove',
      'setTargetBrowsers',
      'setTargetOs',
      'setCampaignItemField',
      'deleteCampaignItemField',
      'addFeedAudience',
      'setLimitsBudgetTotal',
      'setLimitsBudgetDaily',
      'setLimitsClickTotal',
      'setLimitsClickDaily',
    ]),
    ...mapActions('campaigns', [
      'setCampaign'
    ]),
    saveEvent() {
      this.setCampaign
      this.$emit('save')
    },
    closeEvent() {
      this.$emit('close')
    },
    setAllDays(e) {
      for (const day in this.schedule) {
        if (e) {
          for (let i = 0; i <= 23; i++) {
            this.schedule[day].push(i)
          }
        } else {
          this.schedule[day] = []
        }
      }
      this.alldays = e
    },
    setDay(day) {
      if (this.schedule[day].length === 0) {
        for (let i = 0; i <= 23; i++) {
          this.schedule[day].push(i)
        }
      } else {
        this.schedule[day] = []
      }
    },
    setHour(day, hour) {
      if (this.schedule[day].includes(hour)) {
        this.schedule[day].splice(this.schedule[day].indexOf(hour), 1)
      } else {
        this.schedule[day].push(hour)
      }
    },
    toggleTargetItem(item) {
      let target_item;
      if (this.campaignURL && this.isValidUrl(this.campaignURL)) {

        const urlparams = parseUrl(this.campaignURL);
        let params = {};
        let url = "";
        url += urlparams.protocol + "//" + urlparams.hostname

        try {
          params = http_parse_query(urlparams.query)
        } catch (e) {
          console.log(e)
        }

        if (urlparams.port) {
          url += ":" + urlparams.port
        }

        if (urlparams.pathname) {
          url += urlparams.pathname
        }

        if (this.campaignURL.includes(item)) {
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
          this.setCampaignItemField({
            value: url + "?" + query,
            name: 'url'
          })
        } else {
          this.setCampaignItemField({
            value: url,
            name: 'url'
          })
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
    // addCampaignCountry() {
    //   this.addCampaignCountry()
    // },
    changeCampaignCountry(country, index) {
      this.updateCampaignCountryItemCountry({country, index})
    },
    changeCampaignCPC(cpc, index) {
      this.updateCampaignCountryItemCPC({cpc, index})
    },
    removeCampaignCountry(index) {
      this.updateCampaignCountryItemRemove(index)
    },
    checkFeedAudienceType() {
      if (this.campaign.blacklist_feed) {
        return false
      }
      if (this.campaign.whitelist_feed) {
        return false
      }
      return true
    },
    handleAddFeedAudience() {
      let name = 'blacklist_feed'
      if (this.campaign.whitelist_feed) {
        name = 'whitelist_feed'
      }
      console.log(this.newFeedAudienceValue)
      this.addFeedAudience({name, value: this.newFeedAudienceValue})
      this.newFeedAudienceValue = null
    },
    handleDropFeedAudience(event, name) {
      const index = this.count.indexOf(name);
      this.count.splice(index, 1);
    },
  },
  computed: {
    ...mapGetters('libs', [
      'countries',
      'os',
      'browsers'
    ]),
    ...mapGetters('campaigns', [
      'campaigns',
      'campaign',
      'feedAudienceValue',
      'sourceAudienceValue'
    ]),
    ...mapGetters('dsp', [
      'dsp',
    ]),
    getCountries() {
      if (this.campaign.countries) {
        return this.countries.filter(({value: id1}) => !this.campaign.countries.some(({value: id2}) => id2 === id1));
      }
      return this.countries
    },
    formCampaignCountries: {
      get: function () {
        if (this.campaign.countries) {
          return this.campaign.countries.map(item => {
            let c = this.countries.find(o => o.value === item.country);
            return {
              ...item,
              ...c,
            }
          })
        }
        return []
      },
      set: function (value) {

      }
    },
    campaignId: {
      get: function () {
        if (!this.campaign.id) {
          if (this.campaigns.length === 0) {
            this.setCampaignItemField({value: 1, name: 'id'})
            return 1
          }
          let max = this.campaigns.reduce(function (prev, current) {
            return (prev.id > current.id) ? prev : current
          })
          this.setCampaignItemField({value: max.id + 1, name: 'id'})
        }
        return this.campaign.id
      },
      set: function (value) {
        this.setCampaignItemField({value, name: 'id'})
      }
    },
    campaignName: {
      get: function () {
        return this.campaign.name
      },
      set: function (value) {
        this.setCampaignItemField({
          value,
          name: 'name'
        })
      }
    },
    campaignURL: {
      get: function () {
        if (!this.campaign.url) return ''
        return this.campaign.url
      },
      set: function (value) {
        this.setCampaignItemField({
          value,
          name: 'url'
        })
      }
    },
    campaignType: {
      get: function () {
        return this.campaign.type
      },
      set: function (value) {
        this.setCampaignItemField({
          value,
          name: 'type'
        })
      }
    },
    campaignTitle: {
      get: function () {
        return this.campaign.title
      },
      set: function (value) {
        this.setCampaignItemField({
          value,
          name: 'title'
        })
      }
    },
    campaignText: {
      get: function () {
        return this.campaign.text
      },
      set: function (value) {
        this.setCampaignItemField({
          value,
          name: 'text'
        })
      }
    },
    campaignTargetOS: {
      get: function () {
        if (!this.campaign.target || !this.campaign.target.os) return []
        return this.campaign.target.os
      },
    },
    campaignTargetBrowsers: {
      get: function () {
        if (!this.campaign.target || !this.campaign.target.browsers) return []
        return this.campaign.target.browsers
      },
    },

    campaignFreshnessInterval1: {
      get: function () {
        if (!this.campaign.freshness && !this.campaign.freshness.interval1) return null
        return this.campaign.freshness.interval1
      },
      set: function (value) {

      }
    },
    campaignFreshnessInterval2: {
      get: function () {
        if (!this.campaign.freshness && !this.campaign.freshness.interval1) return null
        return this.campaign.freshness.interval1
      },
      set: function (value) {

      }
    },
    campaignFreshnessType: {
      get: function () {
        if (!this.campaign.freshness && !this.campaign.freshness.type) return null
        return this.campaign.freshness.type
      },
      set: function (value) {

      }
    },
    campaignLimitsBudgetTotal: {
      get: function () {
        if (this.campaign.limits && this.campaign.limits.budget_total) {
          return this.campaign.limits.budget_total
        }
        return null
      },
      set: function (value) {
        this.setLimitsBudgetTotal(value)
      }
    },
    campaignLimitsBudgetDaily: {
      get: function () {
        if (this.campaign.limits && this.campaign.limits.budget_daily) {
          return this.campaign.limits.budget_daily
        }
        return null
      },
      set: function (value) {
        this.setLimitsBudgetDaily(value)
      }
    },
    campaignLimitsClickTotal: {
      get: function () {
        if (this.campaign.limits && this.campaign.limits.click_total) {
          return this.campaign.limits.click_total
        }
        return null
      },
      set: function (value) {
        this.setLimitsClickDaily(value)
      }
    },
    campaignLimitsClickDaily: {
      get: function () {
        if (this.campaign.limits && this.campaign.limits.click_daily) {
          return this.campaign.limits.click_daily
        }
        return null
      },
      set: function (value) {
        this.setLimitsClickDaily(value)
      }
    },
    sourceAudienceType: {
      get: function () {
        if (this.campaign.blacklist) {
          return 'blacklist'
        }
        if (this.campaign.whitelist) {
          return 'whitelist'
        }
      },
      set: function (value) {
        if (this.campaign.blacklist_feed && value !== 'blacklist') {
          this.deleteCampaignItemField('blacklist')
          this.setCampaignItemField({
            value: [],
            name: value
          })
        } else if (this.campaign.whitelist_feed && value !== 'whitelist') {
          this.deleteCampaignItemField('whitelist')
          this.setCampaignItemField({
            value: [],
            name: value
          })
        } else {
          this.setCampaignItemField({
            value: [],
            name: value
          })
        }
        this.$forceUpdate()
      }
    },
    feedAudienceType: {
      get: function () {
        if (this.campaign.blacklist_feed) {
          return 'blacklist_feed'
        }
        if (this.campaign.whitelist_feed) {
          return 'whitelist_feed'
        }
        return null
      },
      set: function (value) {
        if (this.campaign.blacklist_feed && value !== 'blacklist_feed') {
          this.deleteCampaignItemField('blacklist_feed')
          this.setCampaignItemField({
            value: [],
            name: value
          })
        } else if (this.campaign.whitelist_feed && value !== 'whitelist_feed') {
          this.deleteCampaignItemField('whitelist_feed')
          this.setCampaignItemField({
            value: [],
            name: value
          })
        } else {
          this.setCampaignItemField({
            value: [],
            name: value
          })
        }
        this.$forceUpdate()
      }
    },
    feedAudienceValue: {
      get: function () {
        if (this.campaign.blacklist_feed) {
          return this.campaign.blacklist_feed
        }
        if (this.campaign.whitelist_feed) {
          return this.campaign.whitelist_feed
        }
        return null
      },
      set: function (value) {
        this.addFeedAudience(value)
      }
    },
    campaignActive: {
      get: function () {
        if (this.campaign.active) {
          return this.campaign.active
        }
        return false
      },
      set: function (value) {
        this.setCampaignItemField({
          value: value,
          name: 'active'
        })
      }
    }
  },
  mounted() {
    this.$store.dispatch('libs/getCountries')
    this.$store.dispatch('libs/getOS')
    this.$store.dispatch('libs/getCountries')
    this.$store.dispatch('dsp/getDSP')
  },
}
</script>

<style scoped>

</style>
