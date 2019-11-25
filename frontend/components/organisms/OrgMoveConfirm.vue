<template>
  <mol-confirm-dialog
    v-model="confirmMove"
    title="本当にページを移動しますか？"
    :move-action="next"
  >
    <p slot="description">
      変更した内容は保存されません
    </p>
  </mol-confirm-dialog>
</template>

<script>
  import MolConfirmDialog from '../molecules/MolConfirmDialog'
  export default {
    name: 'OrgMoveConfirm',
    components: { MolConfirmDialog },
    data () {
      return {
        confirmMove: false,
        next: () => {}
      }
    },
    methods: {
      leaveHandler (next) {
        this.next = next
        this.confirmMove = true
      }
    }
  }
</script>

<style scoped>

</style>
