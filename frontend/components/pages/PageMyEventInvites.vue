<template>
  <div>
    <org-event-header :event="event" />
    <atom-card>
      <h3
        slot="header"
        style="text-align: center"
      >
        {{ event.name }}のスカウト状況
      </h3>
      <org-invite-statuses
        slot="bottom"
        class="org-invite-statuses"
        :invites="invites"
      />
    </atom-card>
  </div>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import OrgInviteStatuses from '../organisms/OrgInviteStatuses'
  import OrgEventHeader from '../organisms/OrgEventHeader'
  export default {
    name: 'PageMyEventInvites',
    components: { OrgEventHeader, OrgInviteStatuses, AtomCard },
    props: {
      event: {
        type: Object,
        required: true
      },
      invites: {
        type: Array,
        required: true
      }
    }
  }
</script>

<style scoped>
  .org-invite-statuses {
    padding-bottom: 16px;
  }
</style>
