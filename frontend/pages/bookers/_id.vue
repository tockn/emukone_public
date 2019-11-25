<template>
  <page-booker-profile
    :booker="booker"
    :events="events"
  />
</template>

<script>
  import PageBookerProfile from '../../components/pages/PageBookerProfile'
  import {mapState} from 'vuex'
  export default {
    components: { PageBookerProfile },
    computed: {
      ...mapState({
        booker: state => state.bookers.booker,
        events: state => state.users.events
      })
    },
    asyncData(context) {
      const id = context.route.params.id;
      return Promise.all([
        context.store.dispatch('bookers/getByID', id),
        context.store.dispatch('users/getEventsByUserID', id)
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
