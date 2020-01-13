import { AxiosError } from 'axios'

export interface IinitialAxiosType {
  $axios: any
}

export default function({ $axios }: IinitialAxiosType) {
  $axios.setBaseURL('https://raw.githubusercontent.com/supano/shio/feature/admin-frontend/packages/admin/mock-data')
  $axios.onError((error: AxiosError) => {
    console.error('custom error', error)
  })
}
