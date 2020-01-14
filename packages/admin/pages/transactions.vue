<template>
  <div>
    <div class="searh-wrapper">
      <a-month-picker @change="onChange" style="width: 300px" size="large" format="MMMM YYYY" placeholder="Select month" />
    </div>
    <a-table :columns="columns" :rowKey="record => record.uuid" :dataSource="this.$accessor.transaction.list" :pagination="paginationState" :loading="loading"> </a-table>
  </div>
</template>

<script lang="ts">
import moment from 'moment'
import { IFetchTransactionType, ITransactionType } from '../store/transaction'

export interface ITransacionTableState {
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
      } as ITransacionTableState,
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
    onChange: async function(date: moment.Moment, dateString: string) {
      this.$data.paginationState = { ...this.$data.paginationState, filterDate: date.unix() }
      await this.reload()
    },
    handleTableChange: async function(pagination) {
      this.$data.pagination = { ...this.$data.paginationState, ...pagination }
      await this.reload()
    },
    reload: async function() {
      this.$data.loading = true

      await this.$accessor.transaction
        .fetch({
          filterDate: this.$data.paginationState.filterDate,
          currentPage: this.$data.paginationState.current,
          pageSize: this.$data.paginationState.pageSize
        } as IFetchTransactionType)
        .finally(() => {
          this.$data.loading = false
        })
    }
  },
  mounted: async function() {
    this.$accessor.pagedetail.setHeader('Transaction')
    await this.reload()
  }
}
</script>
