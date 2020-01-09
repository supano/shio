import { API } from 'shio_rpc_client/v1'

export function createServiceConnector() {
  return new API.ShioAdminAPIClient('http://localhost:4444')
}
