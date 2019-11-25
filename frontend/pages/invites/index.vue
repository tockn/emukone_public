<template>
  <page-invites :invites="invites" />
</template>

<script>
  import { mapState } from 'vuex'
  import PageInvites from '../../components/pages/PageInvites'
  export default {
    components: { PageInvites },
    computed: {
      ...mapState({
        invites: state => state.invites.invites
      })
    },
    asyncData(context) {
      return Promise.all([
        context.store.dispatch('invites/getInvites'),
      ])
        .catch((e) => {
          // eslint-disable-next-line
          console.log(e)
          context.error({ statusCode: 404, message: 'not found' })
        })
    }
  }
</script>

<style scoped>

</style>
