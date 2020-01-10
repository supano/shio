<template>
  <a-layout id="main-layout">
    <a-layout-sider :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0 }">
      <div class="logo">
        CatCat
      </div>
      <a-menu theme="dark" mode="inline" :defaultSelectedKeys="['4']">
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
      <a-layout-header :style="{ background: '#fff', padding: 0 }">
        {{ this.$store.state.pagedetail.header }}
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
  data() {
    return {
      current: ['mail'],
      openKeys: ['sub1']
    }
  },
  methods: {
    handleClick(e: string) {
      console.log('click', e)
    },
    titleClick(e: string) {
      console.log('titleClick', e)
    }
  },
  watch: {
    openKeys(val: string) {
      console.log('openKeys', val)
    }
  }
}
</script>

<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}

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
</style>
