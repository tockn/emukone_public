<template>
  <page-event-edit
    ref="page"
    is-create-page
  />
</template>

<script>
  import PageEventEdit from '../../components/pages/PageEventEdit'
  export default {
    components: { PageEventEdit },
    beforeRouteEnter (to, from, next) {
      next(vm => {
        let storageData = localStorage.getItem('createEvent')
        if (storageData) {
          try {
            let initialEvent = JSON.parse(storageData)
            vm.$refs.page.enterHandler(initialEvent)
          } catch (e) {
            localStorage.removeItem('createEvent')
          }
        }
      })
    },
    beforeRouteLeave (to, from, next) {
      this.$refs.page.leaveHandler()
      next()
    }
  }
</script>

<style scoped>

</style>
