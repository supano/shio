<template>
  <div>
    <TransactionComponent :phone="this.$route.params.detail" />
  </div>
</template>

<script lang="ts">
import TransactionComponent from '@/components/transaction-table'

export default {
  data: function() {
    return {}
  },
  components: {
    TransactionComponent
  },
  mounted: function() {
    this.$accessor.pagedetail.setHeader('User Transaction Detail')
  }
}
</script>
