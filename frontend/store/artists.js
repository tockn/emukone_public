const API_ENDPOINT = process.env.API_ENDPOINT

export const state = () => ({
  artist: {
    name: '',
    icon_url: '',
    identifier: '',
    header_image_url: '',
    meta_description: '',
    why_description: '',
    how_description: '',
    free_description: '',
    tags: [],
    locations: [],
    website_urls: [],
    images: [],
    ichioshis: [],
    members: []
  },
  loading: {
    artist: false
  }
})

export const mutations = {
  setArtist (state, artist) {
    state.artist = artist
  },
  setArtistLoading (state, loading) {
    state.loading.artist = loading
  },
}

export const actions = {
  async getByID ({ commit }, id) {
    commit('setArtistLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/artists/${id}`)
    commit('setArtist', res)
    commit('setArtistLoading', false)
  }
}
