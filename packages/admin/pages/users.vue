<template>
  <div>
    <div class="searh-wrapper">
      <a-input-search placeholder="input search text" style="width: 500px" size="large" @search="onSearch" enterButton />
    </div>
    <a-table
      :columns="columns"
      :rowKey="record => record.email"
      :dataSource="this.$accessor.user.list"
      :pagination="paginationState"
      :loading="loading"
      @change="handleTableChange"
    >
      <span slot="action" slot-scope="record" class="flex justify-center">
        <a href="#" @click="handleSeeMore(record)">
          <a-icon type="eye" :style="{ fontSize: '24px' }" />
        </a>
      </span>
    </a-table>
  </div>
</template>

<script lang="ts">
import { IFetchUserType, IUser } from '../store/user'
import moment from 'moment'

export interface IUserTableState {
  current: number
  pageSize: number
  filterName: string | null
}

export interface IUserPageData {
  loading: boolean
  paginationState: IUserTableState
  columns: Array<object>
}

export default {
  data: function(): IUserPageData {
    return {
      loading: false,
      paginationState: {
        current: 1,
        pageSize: 10,
        filterName: null
      },
      columns: [
        {
          title: 'Customer Name',
          dataIndex: 'name'
        },
        {
          title: 'Email',
          dataIndex: 'email'
        },
        {
          title: 'Last Access',
          dataIndex: 'lastAccess'
        },
        {
          title: 'Last Payment',
          dataIndex: 'lastPayment'
        },
        {
          title: 'Phone No.',
          dataIndex: 'phone'
        },
        {
          title: 'Created At',
          dataIndex: 'createdAt'
        },
        {
          title: 'Action',
          scopedSlots: { customRender: 'action' }
        }
      ]
    }
  },
  methods: {
    onSearch: async function(search: string) {
      this.$data.paginationState = { ...this.$data.paginationState, filterName: search }
      await this.reload()
    },
    handleTableChange: async function(pagination) {
      this.$data.pagination = { ...this.$data.paginationState, ...pagination }
      await this.reload()
    },
    reload: async function() {
      this.$data.loading = true

      await this.$accessor.user
        .fetch({
          filterName: this.$data.paginationState.filterName,
          currentPage: this.$data.paginationState.current,
          pageSize: this.$data.paginationState.pageSize
        } as IFetchUserType)
        .finally(() => {
          this.$data.loading = false
        })
    }
  },
  mounted: async function() {
    this.$accessor.pagedetail.setHeader('Users')
    await this.reload()
  }
}
</script>
