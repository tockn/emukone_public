const API_ENDPOINT = process.env.API_ENDPOINT

const unchecked = 0
const approved = 1
const declined = 2

export const getters = {
  unchecked () {
    return unchecked
  },
  approved () {
    return approved
  },
  declined () {
    return declined
  }
}

export const state = () => ({
  invite: {},
  invites: [],
  loading: {
    invite: false,
    invites: false,
    inviteStatus: false
  }
})

export const mutations = {
  setInvites (state, invites) {
    state.invites = invites
  },
  setInvitesLoading (state, loading) {
    state.loading.invites = loading
  },
  setInvite (state, invite) {
    state.invite = invite
  },
  setInviteLoading (state, loading) {
    state.loading.invite = loading
  },
  setInviteStatusLoading (state, loading) {
    state.loading.inviteStatus = loading
  },
  setInviteStatus(state, {id, status}) {
    if (state.invite.id === id) {
      state.invite.status = approved
    }
    state.invites.forEach(invite => {
      if (invite.id === id) {
        invite.status = status
      }
    })
  }
}

export const actions = {
  async getInvites ({ commit }) {
    commit('setInvitesLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/invites`)
    commit('setInvites', res)
    commit('setInvitesLoading', false)
  },
  async getByID({ commit, state }, id) {
    commit('setInviteLoading', true)

    // invitesに該当idのinviteがあったらAPIリクエストしない
    let gotInvite = state.invites.find(invite => invite.id === id)
      || state.invite.id === id
    if (gotInvite) {
      commit('setInvite', gotInvite)
    } else {
      const res = await this.$axios.$get(`${API_ENDPOINT}/invites/${id}`)
      commit('setInvite', res)
    }
    commit('setInviteLoading', false)
  },
  async getByEventID({ commit }, id) {
    commit('setInvitesLoading', true)
    const res = await this.$axios.$get(`${API_ENDPOINT}/events/${id}/invites`)
    commit('setInvites', res)
    commit('setInvitesLoading', false)
  },
  async approveInvite({ commit }, inviteId) {
    commit('setInviteStatusLoading', true)
    await this.$axios.$put(`${API_ENDPOINT}/invites/${inviteId}/status`,
      {
        status: approved
      }
    )
    commit('setInviteStatus', {
      id: inviteId,
      status: approved
    })
    commit('setInviteStatusLoading', false)
  },
  async declineInvite({ commit }, inviteId) {
    commit('setInviteStatusLoading', true)
    await this.$axios.$put(`${API_ENDPOINT}/invites/${inviteId}/status`,
      {
        status: declined
      }
    )
    commit('setInviteStatus', {
      id: inviteId,
      status: declined
    })
    commit('setInviteStatusLoading', false)
  }
}
