const API_ENDPOINT = process.env.API_ENDPOINT

export const state = () => ({
  event: {},
  relatedUsers: [],
  validationErrorMessages: [],
  loading: {
    event: false,
    relatedUsers: false,
    post: false
  }
})

export const mutations = {
  setEvent (state, event) {
    state.event = event
  },
  setEventLoading (state, loading) {
    state.loading.event = loading
  },
  setEventId (state, id) {
    state.event.id = id
  },
  setRelatedUsers (state, users) {
    state.relatedUsers = users
  },
  setRelatedUsersLoading (state, loading) {
    state.loading.relatedUsers = loading
  },
  setLoadingPost (state, loading) {
    state.loading.post = loading
  }
}

export const actions = {
  async getByID ({ commit }, id) {
    commit('setEventLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/events/${id}`)
    commit('setEvent', res)
    commit('setEventLoading', false)
  },
  async getRelatedUsersByID ({ commit }, id) {
    commit('setRelatedUsersLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/events/${id}/users`)
    commit('setRelatedUsers', res)
    commit('setRelatedUsersLoading', false)
  },
  async postEvent ({ commit }, event) {
    const validateRes = validate(event)
    if (!validateRes.result) {
      throw validateRes.messages
    }
    commit('setLoadingPost', true)
    const res = await this.$axios.post(`${API_ENDPOINT}/events`, event)
    commit('setLoadingPost', false)
    return res.data
  },
  async putEvent ({ commit }, event) {
    const validateRes = validate(event)
    if (!validateRes.result) {
      throw validateRes.messages
    }
    commit('setLoadingPost', true)
    const res = await this.$axios.put(`${API_ENDPOINT}/events`, event)
    commit('setLoadingPost', false)
    return res
  }
}

function validate (event) {
  let returnValue = true
  let errorMsgs = []
  if (!event.name) {
    errorMsgs.push('イベント名を入力してください')
    returnValue = false
  } else if (event.name.length > 50) {
    errorMsgs.push('イベント名は50文字以内で入力してください')
    returnValue = false
  }
  if (!event.why_description) {
    errorMsgs.push('イベントに懸ける想いを入力してください')
    returnValue = false
  }
  if (!event.open_time) {
    errorMsgs.push('OPEN(開場時間)を入力してください')
    returnValue = false
  }
  if (!event.start_time) {
    errorMsgs.push('START(開演時間)を入力してください')
    returnValue = false
  }
  if (!event.place) {
    errorMsgs.push('場所を入力してください')
    returnValue = false
  }
  if (!event.ticket) {
    errorMsgs.push('チケット情報を入力してください')
    returnValue = false
  }
  return {result: returnValue, messages: errorMsgs}
}
