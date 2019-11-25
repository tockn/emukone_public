<template>
  <page-my-event-invites
    :event="event"
    :invites="invites"
  />
</template>

<script>
  import { mapState } from 'vuex'
  import PageMyEventInvites from '../../../components/pages/PageMyEventInvites'
  export default {
    components: { PageMyEventInvites },
    computed: {
      ...mapState({
        event: state => state.events.event,
        invites: state => state.invites.invites
      })
    },
    asyncData(context) {
      const id = context.route.params.id;
      return Promise.all([
        context.store.dispatch('events/getByID', id),
        context.store.dispatch('invites/getByEventID', id),
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
