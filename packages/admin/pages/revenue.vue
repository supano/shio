<template>
  <div>
    <div class="searh-wrapper">
      <a-month-picker @change="onChange" style="width: 300px" size="large" format="MMMM YYYY" placeholder="Select month" />
    </div>
    <a-table
      :columns="columns"
      :rowKey="record => record.name"
      :dataSource="this.$accessor.revenue.list"
      :pagination="paginationState"
      :loading="loading"
      @change="handleTableChange"
    >
    </a-table>
  </div>
</template>

<script lang="ts">
import moment from 'moment'
import { IFetchRevenueType } from '../store/revenue'

export interface IRevenueTableState {
  current: number
  pageSize: number
  filterDate: number
}

export default {
  data: function() {
    return {
      loading: false,
      paginationState: {
        current: 1,
        pageSize: 10,
        filterDate: moment().unix()
      } as IRevenueTableState,
      columns: [
        {
          title: 'Book Name',
          dataIndex: 'name',
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
    onChange: async function(date: moment.Moment, dateString: string) {
      this.$data.loading = true
      this.$data.paginationState = { ...this.$data.paginationState, filterDate: date.unix() }
      await this.reload()
      this.$data.loading = false
    },
    handleTableChange: async function(pagination) {
      this.$data.loading = true
      this.$data.paginationState = { ...this.$data.paginationState, ...pagination }
      await this.reload()
      this.$data.loading = false
    },
    reload() {
      this.$accessor.revenue.fetch({
        filterDate: this.$data.paginationState.filterDate,
        currentPage: this.$data.paginationState.current,
        pageSize: this.$data.paginationState.pageSize
      } as IFetchRevenueType)
    }
  },
  mounted: async function() {
    this.$data.loading = true

    this.$accessor.pagedetail.setHeader('Revenue')
    await this.reload()

    this.$data.loading = false
  }
}
</script>
