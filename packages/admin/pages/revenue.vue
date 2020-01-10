<template>
  <div>
    <div class="flex flex-row mb-5 justify-end">
      <a-month-picker @change="onChange" style="width: 300px" size="large" format="MMMM YYYY" placeholder="Select month" />
    </div>
    <a-table :columns="columns" :rowKey="record => record.name" :dataSource="data" :pagination="pagination" :loading="loading" @change="handleTableChange"> </a-table>
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
          title: 'Book Name',
          dataIndex: 'name',
          sorter: true,
          width: '50%'
        },
        {
          title: 'Total Amount',
          dataIndex: 'amount'
        },
        {
          title: 'MDR 3% (THB)',
          dataIndex: 'mdr'
        },
        {
          title: 'VAT 7% (THB)',
          dataIndex: 'vat'
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
      console.log('pagination', pagination)
      console.log('filters', filters)
      console.log('sorter', sorter)
      this.$data.pagination = { ...pagination }
      setTimeout(() => {
        for (let i = pagination.current; i < pagination.current + 10; i++) {
          this.$data.data.push({ name: `Book Name ${i}`, amount: 10, mdr: 3, vat: 7, net: 97 })
        }
        this.$data.loading = false
      }, 2000)
    },
    fetchRevenue: function() {
      this.$data.loading = true
      for (let i = 0; i < 10; i++) {
        this.$data.data.push({ name: `Book Name ${i}`, amount: 10, mdr: 3, vat: 7, net: 97 })
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
    this.fetchRevenue()
  }
}
</script>
