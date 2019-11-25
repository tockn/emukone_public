const API_ENDPOINT = process.env.API_ENDPOINT

export const state = () => ({
  me: {},
  loading: {
    me: false
  }
})

export const mutations = {
  setMe (state, me) {
    state.me = me
  },
  setMeLoading (state, loading) {
    state.loading.me = loading
  },
}

export const actions = {
  async checkLogin ({ commit }) {
    commit('setMeLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/me`)
    commit('setMe', res)
    commit('setMeLoading', false)
  }
}
