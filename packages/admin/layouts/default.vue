<template>
  <a-layout id="main-layout">
    <a-layout-sider :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0 }">
      <div class="logo">
        CatCat
      </div>
      <a-menu theme="dark" mode="inline" :defaultSelectedKeys="[this.$route.name]">
        <a-menu-item key="revenue">
          <nuxt-link to="/revenue">
            <a-icon type="desktop" />
            <span>ยอดขายประจำเดือน</span>
          </nuxt-link>
        </a-menu-item>
        <a-menu-item key="transactions">
          <nuxt-link to="/transactions">
            <a-icon type="desktop" />
            <span>รายละเอียดการขาย</span>
          </nuxt-link>
        </a-menu-item>
        <a-menu-item key="users">
          <nuxt-link to="/users">
            <a-icon type="desktop" />
            <span>รายละเอียดลูกค้า</span>
          </nuxt-link>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout :style="{ marginLeft: '200px' }">
      <a-layout-header :style="{ background: '#fff', paddingLeft: '20px', paddingRight: '20px' }">
        <h1>{{ header }}</h1>
        <a-button type="link" @click="logout"><a-icon type="logout" /> Logout </a-button>
      </a-layout-header>
      <a-layout-content :style="{ margin: '24px 16px 0', overflow: 'initial' }">
        <div :style="{ padding: '24px', background: '#fff', textAlign: 'center' }">
          <nuxt />
        </div>
      </a-layout-content>
      <a-layout-footer :style="{ textAlign: 'center' }">
        Digithun worldwide
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script lang="ts">
export default {
  methods: {
    logout(): void {
      this.$router.push('/login')
    }
  },
  computed: {
    header: function() {
      return this.$accessor.pagedetail.header
    }
  }
}
</script>

<style>
#main-layout .logo {
  height: 32px;
  line-height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
  text-align: center;
  color: white;
  font-weight: bold;
  font-size: 20px;
}

.ant-layout-header {
  display: flex;
  flex-grow: 1;
  align-items: baseline;
  /* padding-right: 20px; */
  justify-content: space-between;
}
</style>
