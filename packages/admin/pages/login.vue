<template>
  <div id="login-wrapper">
    <div>
      <a-row id="login-panel" align="middle" type="flex" :gutter="0">
        <a-col :span="12">
          <img src="../assets/img/green-leafed-indoor-plant-1048035.jpg" width="500px" alt="" />
        </a-col>
        <a-col :span="12" id="form-login">
          <h1>CatCat Admin Portal</h1>
          <a-input placeholder="email" type="email" size="large" />
          <a-input placeholder="password" type="password" size="large" />
          <a-button type="primary" id="login-btn" size="large" @click="login">Login</a-button>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  layout: 'empty',
  data: function() {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    login: function() {
      return this.$router.push('/')
    }
  }
}
</script>

<style scoped>
#login-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #cecece;
}
#login-panel {
  width: 1000px;
  background-color: white;

  -webkit-box-shadow: 10px 10px 60px -32px rgba(0, 0, 0, 0.75);
  -moz-box-shadow: 10px 10px 60px -32px rgba(0, 0, 0, 0.75);
  box-shadow: 10px 10px 60px -32px rgba(0, 0, 0, 0.75);
}

#form-login {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 30px;
  height: 100%;

  flex-direction: column;
  background-color: white;
}
#login-btn {
  width: 100%;
}
h1 {
  color: #545454;
}
input {
  margin-bottom: 10px;
}
</style>
