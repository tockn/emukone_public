const API_ENDPOINT = process.env.API_ENDPOINT

export const state = () => ({
  event: {},
  loading: {
    uploadImage: false
  }
})

export const mutations = {
  setEvent (state, event) {
    state.event = event
  },
  setLoadingEvent (state, loading) {
    state.loading.event = loading
  }
}

export const actions = {
  async getByID ({ commit }, id) {
    commit('setLoadingEvent', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/events/${id}`)
    commit('setEvent', res)
    commit('setLoadingEvent', false)
  }
}
