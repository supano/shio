import { getterTree, mutationTree, actionTree } from 'nuxt-typed-vuex'
import { AxiosError } from 'axios'

export interface IRevenue {
  assetName: string
  totalAmount: number
  mdr: number
  vat: number
  netAmount: number
}

export interface IFetchRevenueType {
  filterDate: number
  currentPage: number
  pageSize: number
}

export const state = () => ({
  list: [] as IRevenue[]
})

export const getters = getterTree(state, {
  list: state => state.list
})

export const mutations = mutationTree(state, {
  setList(state, list: IRevenue[]) {
    state.list = list
  }
})

export const actions = actionTree(
  { state, getters, mutations },
  {
    fetch({ commit }, filter: IFetchRevenueType): Promise<IRevenue[] | void> {
      return this.$axios
        .$get('/revenue.json', {
          params: {
            createdAt: filter.filterDate,
            skip: filter.currentPage * filter.pageSize,
            limit: filter.pageSize
          }
        })
        .then((res: IRevenue[]) => {
          commit('setList', res)
          return res
        })
        .catch((err: AxiosError) => {
          console.error(err.message)
        })
    }
  }
)
