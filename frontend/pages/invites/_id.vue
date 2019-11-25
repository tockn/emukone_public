<template>
  <page-invite :invite="invite" />
</template>

<script>
  import { mapState } from 'vuex'
  import PageInvite from '../../components/pages/PageInvite'
  export default {
    components: { PageInvite },
    computed: {
      ...mapState({
        invite: state => state.invites.invite,
      })
    },
    asyncData(context) {
      const id = context.route.params.id;
      return Promise.all([
        context.store.dispatch('invites/getByID', id),
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
