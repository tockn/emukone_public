<template>
  <div>
    <div @click="modalOpen = true">
      <atom-header-image
        :src="event.header_image_url"
        :filter="false"
        :max-height="true"
        :absolute="false"
      />
    </div>
    <atom-card>
      <h2 slot="header">
        {{ event.name }}
      </h2>
      <div slot="description">
        <p>"{{ event.meta_description }}"</p>
        <div style="text-align: center;">
          <atom-button
            v-if="editable"
            color="green"
            @click="$router.push(`/events/${event.id}/edit`)"
          >
            編集する
          </atom-button>
        </div>
      </div>
    </atom-card>
    <atom-modal v-model="modalOpen">
      <atom-image
        slot="content"
        :src="event.header_image_url"
      />
      <atom-button
        slot="action"
        @click="modalOpen = false"
      >
        Close
      </atom-button>
    </atom-modal>
  </div>
</template>

<script>
  import AtomHeaderImage from '../atoms/AtomHeaderImage'
  import AtomCard from '../atoms/AtomCard'
  import AtomModal from '../atoms/AtomModal'
  import AtomImage from '../atoms/AtomImage'
  import AtomButton from '../atoms/AtomButton'
  export default {
    name: 'OrgEventHeader',
    components: { AtomImage, AtomModal, AtomCard, AtomHeaderImage, AtomButton },
    props: {
      event: {
        type: Object,
        required: true
      },
      editable: {
        type: Boolean,
        default: false
      }
    },
    data () {
      return {
        modalOpen: false
      }
    }
  }
</script>

<style scoped>
  h2 {
    text-align: center;
  }
  p {
    text-align: center;
    font-style: italic;
  }
</style>
