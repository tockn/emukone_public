<template>
  <div class="org-header-wrap">
    <mol-meta-data-card
      :user="user"
    >
      <slot />
    </mol-meta-data-card>
    <atom-header-image :src="user.icon_url" />
  </div>
</template>

<script>
  import AtomHeaderImage from '../atoms/AtomHeaderImage'
  import MolMetaDataCard from '../molecules/MolMetaDataCard'
  export default {
    name: 'OrgProfileHeader',
    components:{ MolMetaDataCard, AtomHeaderImage },
    props: {
      user: {
        type: Object,
        required: true
      }
    }
  }
</script>

<style scoped>
  .org-header-wrap {
    height: 360px;
    padding-top: 16px;
  }
</style>
