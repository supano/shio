import { getterTree, mutationTree, actionTree } from 'nuxt-typed-vuex'
import { AxiosError } from 'axios'

export interface IUser {
  name: string
  email: string
  lastAccess: string
  lastPayment: string
  phone: string
  createdAt: string
}

export interface IFetchUserType {
  filterName: string | null
  currentPage: number
  pageSize: number
}

export const state = () => ({
  list: [] as IUser[]
})

export const getters = getterTree(state, {
  list: state => state.list
})

export const mutations = mutationTree(state, {
  setList(state, list: IUser[]) {
    state.list = list
  }
})

export const actions = actionTree(
  { state, getters, mutations },
  {
    fetch({ commit }, filter: IFetchUserType) {
      this.$axios
        .$get('/user.json', {
          params: {
            name: filter.filterName,
            skip: filter.currentPage * filter.pageSize,
            limit: filter.pageSize
          }
        })
        .then((res: IUser[]) => {
          commit('setList', res)
          return res
        })
        .catch((err: AxiosError) => {
          console.error(err.message)
        })
    }
  }
)
