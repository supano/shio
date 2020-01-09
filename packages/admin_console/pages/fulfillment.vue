<template>
  <div>
    asd
  </div>
</template>

<script lang="ts">
import { createServiceConnector } from '../utils/index'
import * as grpcWeb from 'grpc-web'
import { GetTransactionListRequest, GetTransactionListResponse } from '../shio_rpc_client/__generated__/endpoint/v1/shio_admin_api_pb'

const api = createServiceConnector()

export default {
  data: function() {
    return {
      list: [],
      api: createServiceConnector()
    }
  },
  methods: {
    fetch: function() {
      const req = new GetTransactionListRequest()
      api.getTransactionList(req, (err: grpcWeb.Error | null, response: GetTransactionListResponse | null) => {
        if (err) {
          console.error(err)
          return
        }
        if (response) {
          console.log('response', response.toObject())
        }
      })
    }
  }
}
</script>
