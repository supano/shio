import { getStoreType, getAccessorType, mutationTree, actionTree } from 'nuxt-typed-vuex'
import { Context } from '@nuxt/types'

import * as pagedetail from './pagedetail'

export const state = () => ({})

type RootState = ReturnType<typeof state>

export const getters = {}

export const mutations = mutationTree(state, {
  initialiseStore() {
    console.log('initialised')
  }
})

export const actions = actionTree({ state, getters, mutations }, {})

export const accessorType = getAccessorType({
  actions,
  getters,
  mutations,
  state,
  modules: {
    pagedetail
  }
})
