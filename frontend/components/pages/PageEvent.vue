<template>
  <div>
    <org-event-header
      :event="event"
      :editable="isMine"
    />
    <mol-content-card
      title="イベントに懸ける想い"
      :description="event.why_description"
    />

    <org-event-info-card
      :event="event"
    />

    <org-horizontal-scroll-user-card
      title="このイベントの管理者"
      :users="relatedUsers"
    />
  </div>
</template>

<script>
  import MolContentCard from '../molecules/MolContentCard'
  import OrgEventInfoCard from '../organisms/OrgEventInfoCard'
  import OrgEventHeader from '../organisms/OrgEventHeader'
  import OrgHorizontalScrollUserCard from '../organisms/OrgHorizontalScrollUserCard'
  export default {
    name: 'PageEvent',
    components: { OrgHorizontalScrollUserCard, OrgEventHeader, OrgEventInfoCard, MolContentCard },
    props: {
      event: {
        type: Object,
        required: true
      },
      relatedUsers: {
        type: Array,
        required: true
      },
      me: {
        type: Object,
        required: true
      }
    },
    computed: {
      isMine () {
        if (!this.me) return
        return this.relatedUsers.filter(u => u.id === this.me.id).length === 1
      }
    }
  }
</script>

<style scoped>

</style>
