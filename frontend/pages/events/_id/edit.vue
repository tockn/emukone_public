<template>
  <page-event-edit
    ref="page"
  />
</template>

<script>
  import { mapState } from 'vuex'
  import PageEventEdit from '../../../components/pages/PageEventEdit'
  export default {
    components: { PageEventEdit },
    computed: {
      ...mapState('events',{
        event: state => state.event
      })
    },
    async asyncData(context) {
      const id = context.route.params.id;
      try {
        await context.store.dispatch('events/getByID', id)
        await context.store.dispatch('events/getRelatedUsersByID', id)
      } catch (e) {
        // eslint-disable-next-line
        console.log(e)
        context.error({ statusCode: 404, message: 'not found' })
      }
      let relatedUsers = context.store.state.events.relatedUsers
      let myId = context.store.state.auth.me.id
      let isMine = relatedUsers.filter(u => u.id === myId)
        .length === 1
      if (!isMine) {
        context.error({ statusCode: 403, message: '編集権限がありません'})
      }
    },
    beforeRouteEnter (to, from, next) {
      next(vm => {
        let copy = JSON.parse(JSON.stringify(vm.event))
        vm.$refs.page.enterHandler(copy)
      })
    },
    beforeRouteLeave (to, from, next) {
      this.$refs.page.leaveHandler(next)
    }
  }
</script>

<style scoped>

</style>
