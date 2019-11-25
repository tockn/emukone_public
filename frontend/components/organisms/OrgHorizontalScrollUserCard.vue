<template>
  <mol-content-card
    :title="title"
  >
    <div
      slot="bottom"
      class="horizontal-list"
    >
      <mol-content-small-card
        v-for="(user, i) in users"
        :key="i"
        class="item"
        :title="user.name"
        :img-src="user.icon_url"
        :link-to="`/${user.identifier}s/${user.id}`"
      />
    </div>
  </mol-content-card>
</template>

<script>
  import MolContentCard from '../molecules/MolContentCard'
  import MolContentSmallCard from '../molecules/MolContentSmallCard'
  export default {
    name: 'OrgHorizontalScrollUserCard',
    components: { MolContentSmallCard, MolContentCard },
    props: {
      title: {
        type: String,
        required: true
      },
      users: {
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

