import { getterTree, mutationTree, actionTree } from 'nuxt-typed-vuex'
import { AxiosError } from 'axios'

export interface ITransactionType {
  createdAt: Date
  createdBy: string
  assetName: string
  net: number
}

export interface IFetchTransactionType {
  filterDate: number
  currentPage: number
  pageSize: number
}

export const state = () => ({
  list: [] as ITransactionType[]
})

export const getters = getterTree(state, {
  list: state => state.list
})

export const mutations = mutationTree(state, {
  setList(state, list: ITransactionType[]) {
    state.list = list
  }
})

export const actions = actionTree(
  { state, getters, mutations },
  {
    fetch({ commit }, filter: IFetchTransactionType) {
      this.$axios
        .$get('/transaction.json', {
          params: {
            createdAt: filter.filterDate,
            skip: filter.currentPage * filter.pageSize,
            limit: filter.pageSize
          }
        })
        .then((res: ITransactionType[]) => {
          commit('setList', res)
          return res
        })
        .catch((err: AxiosError) => {
          console.error(err.message)
        })
    }
  }
)
