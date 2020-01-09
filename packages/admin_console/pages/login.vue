<template>
  <div id="login-wrapper">
    <input type="email" name="" id="" v-model="email" /> <br />
    <input type="password" name="" id="" v-model="password" /> <br />
    <button @click="login">login</button>
  </div>
</template>

<script lang="ts">
import { API } from '../shio_rpc_client/v1'
import * as grpcWeb from 'grpc-web'
import { ShioAdminAPIClient } from '../shio_rpc_client/__generated__/endpoint/v1/shio_admin_api_grpc_web_pb'
import { LoginRequest, LoginResponse } from '../shio_rpc_client/__generated__/endpoint/v1/shio_admin_api_pb'

const api = new API.ShioAdminAPIClient('http://localhost:4444')
export default {
  data: function() {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    login: function() {
      console.log('email', this.$data.email)
      console.log('password', this.$data.password)

      const req = new LoginRequest()
      api.login(req, (err: grpcWeb.Error | null, response: LoginResponse | null) => {
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

<style scoped>
#login-wrapper {
  height: 100vh;
  width: 100vw;
  background-color: bisque;
}
</style>
