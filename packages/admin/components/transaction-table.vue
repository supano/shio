<template>
  <div>
    <div class="searh-wrapper">
      <a-month-picker @change="onFilter" style="width: 300px" size="large" format="MMMM YYYY" placeholder="Select month" />
    </div>
    <a-table
      :columns="columns"
      :rowKey="record => record.uuid"
      :dataSource="this.$accessor.transaction.list"
      :pagination="paginationState"
      :loading="loading"
      @change="onTableChange"
    >
      <template slot="createdAt" slot-scope="createdAt"> {{ formatTime(createdAt) }} </template>
    </a-table>
  </div>
</template>

<script lang="ts">
import moment from 'moment'
import { IFetchTransactionType, ITransaction } from '@/store/transaction'
import timeUtils from '@/utils/time'

export interface ITransacionTableState {
  current: number
  pageSize: number
}

export interface ITransactionPageData {
  loading: boolean
  paginationState: ITransacionTableState
  columns: Array<object>
  filterDate: number
}

export default {
  props: {
    phone: {
      type: String
    }
  },
  data: function(): ITransactionPageData {
    return {
      loading: false,
      filterDate: moment().unix(),
      paginationState: {
        current: 1,
        pageSize: 10
      },
      columns: [
        {
          title: 'Date',
          dataIndex: 'createdAt',
          scopedSlots: { customRender: 'createdAt' }
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
    onFilter: async function(date: moment.Moment, dateString: string) {
      this.$data.filterDate = date.unix()
      await this.reload()
    },
    onTableChange: async function(pagination) {
      this.$data.paginationState = { ...this.$data.paginationState, ...pagination }
      await this.reload()
    },
    reload: async function() {
      this.$data.loading = true

      await this.$accessor.transaction
        .fetch({
          filterDate: this.$data.filterDate,
          filterPhone: this.$props.phone,
          currentPage: this.$data.paginationState.current,
          pageSize: this.$data.paginationState.pageSize
        } as IFetchTransactionType)
        .finally(() => {
          this.$data.loading = false
        })
    },
    formatTime(t: number) {
      return timeUtils.timestampToDateString(t)
    }
  },
  mounted: function() {
    this.reload()
  }
}
</script>
