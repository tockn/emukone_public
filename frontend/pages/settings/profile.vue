<template>
  <div>
    <page-edit-artist-profile v-if="me.identifier === 'artist'" />
    <page-edit-booker-profile v-else />
    <org-move-confirm ref="confirm" />
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import PageEditArtistProfile from '../../components/pages/PageEditArtistProfile'
  import PageEditBookerProfile from '../../components/pages/PageEditBookerProfile'
  import OrgMoveConfirm from '../../components/organisms/OrgMoveConfirm'
  export default {
    name: 'Profile',
    components: { OrgMoveConfirm, PageEditBookerProfile, PageEditArtistProfile },
    data () {
      return {
        moveAction: () => {},
        notMoveAction: () => {},
        history: {}
      }
    },
    computed: {
      ...mapState({
        user: state => state.users.user,
        me: state => state.auth.me,
        updated: state => state.users.updated
      })
    },
    asyncData(context) {
      return Promise.all([
        context.store.dispatch('users/getMe'),
      ])
        .catch((e) => {
          // eslint-disable-next-line
          console.log(e)
          context.error({ statusCode: e.status, message: 'not found' })
        })
    },
    beforeRouteLeave (to, from, next) {
      if (!this.updated) {
        this.$refs.confirm.leaveHandler(next)
      } else {
        next()
      }
    }
  }
</script>

<style scoped>

</style>
