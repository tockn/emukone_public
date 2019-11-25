const API_ENDPOINT = process.env.API_ENDPOINT

export const state = () => ({
  booker: {},
  loading: {
    booker: false
  }
})

export const mutations = {
  setBooker (state, booker) {
    state.booker = booker
  },
  setBookerLoading (state, loading) {
    state.loading.booker = loading
  }
}

export const actions = {
  async getByID ({ commit }, id) {
    commit('setBookerLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/bookers/${id}`)
    commit('setBooker', res)
    commit('setBookerLoading', false)
  }
}
