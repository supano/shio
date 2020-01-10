<template>
  <div id="login-wrapper">
    <div id="login-panel">
      <a-row>
        <a-col :span="12">
          <img src="../assets/img/photo-of-plant-on-pot-1022922.jpg" width="500px" alt="" />
        </a-col>
        <a-col :span="12" id="form-login">
          <h1>CatCat Admin Portal</h1>
          <a-input placeholder="email" type="email" size="large" />
          <a-input placeholder="password" type="password" size="large" />
          <a-button type="primary" id="login-btn" size="large">Login</a-button>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  layout: 'empty'
}
</script>

<style scoped>
#login-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #e7e7e7;
}
#login-panel {
  width: 1000px;
  border: 1px solid #e0e0e0;
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
