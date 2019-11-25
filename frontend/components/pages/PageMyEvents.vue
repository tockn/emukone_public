<template>
  <div>
    <atom-card>
      <h3
        slot="header"
        style="text-align: center"
      >
        管理しているイベント一覧
      </h3>
    </atom-card>
    <div style="text-align: center;">
      <atom-button
        color="green"
        @click="$router.push('/events/create')"
      >
        イベントを作成
      </atom-button>
    </div>
    <org-my-event-card
      v-for="(event, i) in events"
      :key="i"
      :event="event"
    />
  </div>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import OrgMyEventCard from '../organisms/OrgMyEventCard'
  import AtomButton from '../atoms/AtomButton'
  export default {
    name: 'PageMyEvents',
    components: { AtomButton, OrgMyEventCard, AtomCard },
    props: {
      events: {
        type: Array,
        required: true
      }
    }
  }
</script>

<style scoped>

</style>
