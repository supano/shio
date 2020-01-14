import { getterTree, mutationTree, actionTree } from 'nuxt-typed-vuex'
import { AxiosError } from 'axios'

export interface ITransaction {
  createdAt: Date
  createdBy: string
  assetName: string
  net: number
}

export interface IFetchTransactionType {
  filterDate: number
  filterPhone?: string
  currentPage: number
  pageSize: number
}

export const state = () => ({
  list: [] as ITransaction[]
})

export const getters = getterTree(state, {
  list: state => state.list
})

export const mutations = mutationTree(state, {
  setList(state, list: ITransaction[]) {
    state.list = list
  }
})

export const actions = actionTree(
  { state, getters, mutations },
  {
    fetch({ commit }, filter: IFetchTransactionType): Promise<ITransaction[] | void> {
      return this.$axios
        .$get('/transaction.json', {
          params: {
            createdAt: filter.filterDate,
            skip: filter.currentPage * filter.pageSize,
            limit: filter.pageSize,
            phone: filter.filterPhone
          }
        })
        .then((res: ITransaction[]) => {
          commit('setList', res)
          return res
        })
        .catch((err: AxiosError) => console.error(err))
    }
  }
)
