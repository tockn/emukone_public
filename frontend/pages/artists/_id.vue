<template>
  <page-artist-profile
    :artist="artist"
    :events="events"
  />
</template>

<script>
  import {mapState} from 'vuex'
  import PageArtistProfile from '../../components/pages/PageArtistProfile'
  export default {
    components: { PageArtistProfile },
    computed: {
      ...mapState({
        artist: state => state.artists.artist,
        events: state => state.users.events
      })
    },
    asyncData(context) {
      const id = context.route.params.id;
      return Promise.all([
        context.store.dispatch('artists/getByID', id),
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
