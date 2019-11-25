<template>
  <page-my-events
    :events="events"
  />
</template>

<script>
  import { mapState } from 'vuex'
  import PageMyEvents from '../../components/pages/PageMyEvents'
  export default {
    components: { PageMyEvents },
    computed: {
      ...mapState('users',{
        events: state => state.events
      })
    },
    asyncData(context) {
      return Promise.all([
        context.store.dispatch('users/getMyEvents')
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
