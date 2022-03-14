import { env, starport, blocks, wallet, transfers } from '@starport/vuex'

// init modules you need
export default function init(store) {
  transfers(store)
  starport(store)
  blocks(store)
  env(store)
  wallet(store)
}