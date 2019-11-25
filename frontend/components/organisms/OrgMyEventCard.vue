<template>
  <atom-card>
    <h3 slot="header">
      {{ event.name }}
    </h3>
    <div slot="description">
      <nuxt-link
        :to="`/events/${event.id}`"
      >
        <org-event-info-card
          :event="event"
        />
      </nuxt-link>
      <div style="text-align: center !important;">
        <atom-button
          color="blue"
          @click="$router.push(`/events/${event.id}/invites`)"
        >
          スカウト一覧
        </atom-button>
      </div>
    </div>
  </atom-card>
</template>

<script>
  import AtomButton from '../atoms/AtomButton'
  import OrgEventInfoCard from './OrgEventInfoCard'
  import AtomCard from '../atoms/AtomCard'
  export default {
    name: 'OrgMyEventCard',
    components: { AtomCard, OrgEventInfoCard, AtomButton },
    props: {
      event: {
        type: Object,
        required: true
      }
    }
  }
</script>

<style scoped>

</style>
