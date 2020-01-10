<template>
  <div>
    <div class="flex flex-col flex-start mb-5">
      <a-input-search placeholder="input search text" style="width: 500px" size="large" @search="onSearch" enterButton />
    </div>
    <a-table :columns="columns" :rowKey="record => record.email" :dataSource="data" :pagination="pagination" :loading="loading" @change="handleTableChange">
      <span slot="action" slot-scope="record" class="flex justify-center">
        <a href="#" @click="handleSeeMore(record)">
          <a-icon type="eye" :style="{ fontSize: '24px' }" />
        </a>
      </span>
    </a-table>
  </div>
</template>

<script lang="ts">
interface IUser {
  name: string
  email: string
  lastAccess: string
  lastPayment: string
  phone: string
  createdAt: string
}

export default {
  data: function() {
    return {
      data: [] as IUser[],
      loading: false,
      pagination: {},
      columns: [
        {
          title: 'Customer Name',
          dataIndex: 'name',
          sorter: true
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
    handleTableChange: function(pagination, filters, sorter) {
      this.$data.loading = true
      this.$data.pagination = { ...pagination }
      setTimeout(() => {
        for (let i = pagination.current; i < pagination.current + 10; i++) {
          this.$data.data.push({
            name: 'Name',
            email: `email${i}@email.com`,
            lastAccess: new Date().toLocaleString(),
            lastPayment: new Date().toLocaleString(),
            phone: '0800000000',
            createdAt: new Date().toLocaleString()
          })
        }
        this.$data.loading = false
      }, 2000)
    },
    fetchUsers: function() {
      this.$data.loading = true
      for (let i = 0; i < 10; i++) {
        this.$data.data.push({
          name: 'Name',
          email: `email${i}@email.com`,
          lastAccess: new Date().toLocaleString(),
          lastPayment: new Date().toLocaleString(),
          phone: '0800000000',
          createdAt: new Date().toLocaleString()
        })
      }
      this.$data.loading = false
    },
    handleSeeMore: function(record) {
      alert(record.email)
    }
  },
  mounted: function() {
    this.$data.pagination = {
      current: 1,
      total: 500,
      pageSize: 10
    }
    this.fetchUsers()
  }
}
</script>
