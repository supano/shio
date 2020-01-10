<template>
  <div>
    <div class="flex flex-row mb-5 justify-end">
      <a-month-picker @change="onChange" style="width: 300px" size="large" format="MMMM YYYY" placeholder="Select month" />
    </div>
    <a-table :columns="columns" :rowKey="record => record.uuid" :dataSource="data" :pagination="pagination" :loading="loading" @change="handleTableChange"> </a-table>
  </div>
</template>

<script lang="ts">
export default {
  data: function() {
    return {
      data: [],
      loading: false,
      pagination: {},
      columns: [
        {
          title: 'Date',
          dataIndex: 'createdAt',
          sorter: true
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
    handleTableChange: function(pagination, filters, sorter) {
      this.$data.loading = true
      this.$data.pagination = { ...pagination }
      setTimeout(() => {
        for (let i = pagination.current; i < pagination.current + 10; i++) {
          this.$data.data.push({ uuid: i, createdAt: new Date().toLocaleString(), createdBy: 'Bat man', assetName: 'Batmobile', net: 700 })
        }
        this.$data.loading = false
      }, 2000)
    },
    fetchTransactions: function() {
      this.$data.loading = true
      for (let i = 0; i < 10; i++) {
        this.$data.data.push({ uuid: i, createdAt: new Date().toLocaleString(), createdBy: 'Bat man', assetName: 'Batmobile', net: 700 })
      }
      this.$data.loading = false
    }
  },
  mounted: function() {
    this.$data.pagination = {
      current: 1,
      total: 500,
      pageSize: 10
    }
    this.fetchTransactions()
  }
}
</script>
