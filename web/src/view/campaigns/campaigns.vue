<template>
  <div>
    <CompanyModal :show="showModal"
                  @close="showModal = false"
                  @save="saveCompany"
    >

    </CompanyModal>
    <div>
      <Button size="small" style="margin-right: 5px" type="primary" @click="showModal = true">
        Создать Кампанию
      </Button>
    </div>
    <br/>
    <Table :columns="columns" :data="data" border no-data-text="No data">
      <template slot="action" slot-scope="{ row, index }">
        <Tooltip content="Edit" placement="bottom" theme="light">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">
            <Icon type="ios-settings" />
          </Button>
        </Tooltip>
        <Tooltip content="Copy" placement="bottom" theme="light">
          <Button size="small" style="margin-right: 5px" type="primary" @click="show(index)">
            <Icon type="md-copy" />
          </Button>
        </Tooltip>
        <Tooltip content="Remove" placement="bottom-end" theme="light">
          <Button size="small" type="error" @click="remove(index)">
            <Icon type="md-remove-circle" />
          </Button>
        </Tooltip>
      </template>
      <template slot="footer">
        <br/>
      </template>
    </Table>
  </div>
</template>


<script>
import CompanyModal from "./campaignModal";
export default {
  components: {CompanyModal},
  data() {
    return {
      showModal: false,
      columns: [
        {type: 'selection', width: 60, align: 'center'},
        {title: 'ID', key: 'id'},
        {title: 'Название кампании', key: 'name'},
        {title: 'Изображение', key: 'ad_icon'},
        {title: 'Статус', key: 'status'},
        {title: 'Ограничение бюджета', key: 'limit_click'},
        {title: 'Ограничение переходов по ссылке', key: 'limit_budget'},
        {title: 'Создано', key: 'created'},
        {title: 'Действия', slot: 'action', width: 150, align: 'center'}
      ],
    }
  },
  methods: {
    saveCompany() {
      this.showModal = false
    }
  },
  mounted () {
    this.$store.dispatch('libs/getCountries')
    this.$store.dispatch('libs/getOS')
    this.$store.dispatch('libs/getBrowsers')
  }
}
</script>

<style lang="less">

</style>
