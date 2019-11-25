export const actions = {
  async nuxtServerInit ({ dispatch }) {
    await dispatch('auth/checkLogin')
      .catch(() => {

      })
  }
}
