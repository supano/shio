<template>
  <TransactionComponent />
</template>

<script lang="ts">
import TransactionComponent from '@/components/transaction-table'

export default {
  components: {
    TransactionComponent: TransactionComponent
  },
  mounted: async function() {
    this.$accessor.pagedetail.setHeader('Transaction')
  }
}
</script>
