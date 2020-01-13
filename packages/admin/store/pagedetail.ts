import { getterTree, mutationTree, actionTree } from 'nuxt-typed-vuex'

export interface IPageDetailType {
  header: string
}

export const state = () =>
  ({
    header: ''
  } as IPageDetailType)

export const getters = getterTree(state, {
  header: state => state.header as string
})

export const mutations = mutationTree(state, {
  setHeader(state, newValue: string) {
    state.header = newValue
  }
})

export const actions = actionTree(
  { state, getters, mutations },
  {
    initialise({ commit }) {
      commit('setHeader', 'This is Default Header')
    },
    setHeader({ commit }, header: string) {
      commit('setHeader', header)
    }
  }
)
