<template>
  <div>
    <div class="searh-wrapper">
      <a-month-picker @change="onChange" style="width: 300px" size="large" format="MMMM YYYY" placeholder="Select month" />
    </div>
    <a-table :columns="columns" :rowKey="record => record.uuid" :dataSource="data" :pagination="pagination" :loading="loading"> </a-table>
  </div>
</template>

<script lang="ts">
import { Moment } from 'moment'

export default {
  data: function() {
    return {
      data: [],
      loading: false,
      pagination: {},
      columns: [
        {
          title: 'Date',
          dataIndex: 'createdAt'
        },
        {
          title: 'Customer Name',
          dataIndex: 'createdBy'
        },
        {
          title: 'Book Name',
          dataIndex: 'assetName'
        },
        {
          title: 'Net Amount (THB)',
          dataIndex: 'net'
        }
      ]
    }
  },
  methods: {
    onChange: function(date: moment, dateString: string) {},
    handleTableChange: function(pagination, filters, sorter) {
      this.$data.loading = true
      this.$data.pagination = { ...pagination }
      this.$data.loading = false
    }
  },
  mounted: function() {
    this.$data.pagination = {
      current: 1,
      total: 500,
      pageSize: 10
    }
    // this.fetchTransactions()

    for (let i = 0; i < 500; i++) {
      this.$data.data.push({ uuid: i, createdAt: new Date().toLocaleString(), createdBy: 'Bat man', assetName: 'Batmobile', net: 700 })
    }

    console.log(this.$data.data)
  }
}
</script>
