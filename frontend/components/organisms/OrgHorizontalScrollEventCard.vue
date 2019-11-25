<template>
  <mol-content-card
    :title="title"
  >
    <div
      slot="bottom"
      class="horizontal-list"
    >
      <mol-content-small-card
        v-for="(event, i) in events"
        :key="i"
        class="item"
        :title="event.name"
        :img-src="event.header_image_url"
        :link-to="`/events/${event.id}`"
      />
    </div>
  </mol-content-card>
</template>

<script>
  import MolContentCard from '../molecules/MolContentCard'
  import MolContentSmallCard from '../molecules/MolContentSmallCard'
  export default {
    name: 'OrgHorizontalScrollEventCard',
    components: { MolContentSmallCard, MolContentCard },
    props: {
      title: {
        type: String,
        required: true
      },
      events: {
        type: Array,
        required: true
      }
    }
  }
</script>

<style scoped>
  img {
    width: 100%;
    height: 160px;
    object-fit: cover;
  }
  .horizontal-list {
    overflow-x: auto;
    white-space: nowrap;
    -webkit-overflow-scrolling: touch;
  }

  .item {
    display: inline-block;
    width: 90%;
    margin-left: 4px !important;
    margin-right: 4px !important;
  }
</style>
