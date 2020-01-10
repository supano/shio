import * as Vuex from 'vuex'

new Vuex.Store({
  modules: {
    pagedetail: {
      namespaced: true,
      state: () => ({
        header: 'Default Header' as string
      }),
      mutations: {
        update(state, { text }) {
          state.header = text
        }
      }
    }
  }
})
