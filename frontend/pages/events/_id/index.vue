<template>
  <page-event
    :event="event"
    :related-users="relatedUsers"
    :me="me"
  />
</template>

<script>
  import {mapState} from 'vuex'
  import PageEvent from '../../../components/pages/PageEvent'
  export default {
    components: { PageEvent },
    computed: {
      ...mapState({
        event: state => state.events.event,
        relatedUsers: state => state.events.relatedUsers,
        me: state => state.auth.me
      })
    },
    asyncData(context) {
      const id = context.route.params.id;
      return Promise.all([
        context.store.dispatch('events/getByID', id),
        context.store.dispatch('events/getRelatedUsersByID', id)
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
