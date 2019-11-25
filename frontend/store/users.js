import ImageResizer from '../plugins/image-resizer'

const API_ENDPOINT = process.env.API_ENDPOINT

export const state = () => ({
  userMeta: {},
  user: {
    name: '',
    icon_url: '',
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

  tags: [],

  // ユーザーが関わったイベント
  events: [],

  updated: true,

  loading: {
    userMeta: false,
    events: false,
    user: false,
    editUser: false,
    tags: false,
    iconImage: false
  }
})

export const mutations = {
  setUser (state, user) {
    state.user = user
    state.updated = false
  },
  setUserMeta (state, userMeta) {
    state.userMeta = userMeta
  },
  setUpdated (state, updated) {
    state.updated = updated
  },
  setUserMetaLoading (state, loading) {
    state.loading.userMeta = loading
  },
  setEvents (state, events) {
    state.events = events
  },
  setEventsLoading (state, loading) {
    state.loading.events = loading
  },
  setName (state, name) {
    state.user.name = name
  },
  setIconUrl (state, iconUrl) {
    state.user.icon_url = iconUrl
  },
  setHeaderImageUrl (state, headerImageUrl) {
    state.user.header_image_url = headerImageUrl
  },
  setMetaDescription (state, metaDescription) {
    state.user.meta_description = metaDescription
  },
  setWhyDescription (state, whyDescription) {
    state.user.why_description = whyDescription
  },
  setHowDescription (state, howDescription) {
    state.user.how_description = howDescription
  },
  setFreeDescription (state, freeDescription) {
    state.user.free_description = freeDescription
  },
  setTags (state, tags) {
    state.user.tags = tags
  },
  setLocations (state, locations) {
    state.user.locations = locations
  },
  setWebsiteUrls (state, websiteUrls) {
    state.user.website_urls = websiteUrls
  },
  setImages (state, images) {
    state.user.images = images
  },
  setIchioshi (state, ichioshi) {
    state.user.ichioshis[0] = ichioshi
  },
  setMembers (state, members) {
    state.user.members = members
  },
  setUserLoading (state, loading) {
    state.loading.user = loading
  },
  setEditUserLoading (state, loading) {
    state.loading.editUser = loading
  },
  setTagsLoading (state, loading) {
    state.loading.tags = loading
  },
  setAllTags (state, tags) {
    state.tags = tags
  },
  setIconImageLoading (state, loading) {
    state.loading.iconImage = loading
  }
}

export const actions = {
  async getMetaByID ({ commit }, id) {
    commit('setUserMetaLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/users/${id}/meta`)
    commit('setUserMeta', res)
    commit('setUserMetaLoading', false)
  },
  // 任意のユーザーIDを持つユーザーが関係者になっているイベントを取得
  async getEventsByUserID ({ commit }, id) {
    commit('setEventsLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/users/${id}/events`)
    commit('setEvents', res)
    commit('setEventsLoading', false)
  },
  async getMyEvents ({ dispatch, state }) {
    if (!state.user.id) {
      await dispatch('getMe')
    }
    await dispatch('getEventsByUserID', state.user.id)
  },
  async getMe ({ commit }) {
    const meRes = await this.$axios.$get(`${API_ENDPOINT}/me`)
    commit('setUserLoading', false)
    if (!meRes) return
    commit('setUserLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/${meRes.identifier}s/${meRes.id}`)
    commit('setUser', res)
  },
  async editUser ({ commit, state }) {
    commit('setEditUserLoading', true)
    await this.$axios.$put(`${API_ENDPOINT}/${state.user.identifier}s`, state.user)
    commit('setUpdated', true)
    commit('setEditUserLoading', false)
  },
  async getAllTags ({ commit, state }) {
    if (!state.tags || state.tags.length !== 0) return
    commit('setTagsLoading')
    const res = await this.$axios.$get(`${API_ENDPOINT}/tags`)
    commit('setAllTags', res)
  },
  async updateIconUrl ({ commit }, file) {
    commit('setIconImageLoading', true)

    ImageResizer(file, async (base64) => {
      base64 = base64.replace(/^.*,/, '')
      const res = await this.$axios.$post(`${API_ENDPOINT}/statics/images`,
        {
          content_type: file.type,
          base64: base64,
        })
      commit('setIconUrl', res.url)
      commit('setIconImageLoading', false)
    })
  }
}
