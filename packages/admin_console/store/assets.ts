export const assets = () => ({
  list: [] as string[]
})

export const mutations = {
  fetch(assets: string[]) {
    console.log(assets)
  }
}
