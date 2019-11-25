<template>
  <mol-content-card
    :title="invite.event.name"
    :image-src="invite.event.header_image_url"
  >
    <atom-card-content slot="top">
      <atom-image
        :src="invite.inviter.icon_url"
        shape="circular"
        size="mini"
      />
      {{ invite.inviter.name }}
    </atom-card-content>
  </mol-content-card>
</template>

<script>
  import MolContentCard from '../molecules/MolContentCard'
  import AtomCardContent from '../atoms/AtomCardContent'
  import AtomImage from '../atoms/AtomImage'
  export default {
    name: 'OrgInviteLargeCard',
    components: { AtomImage, AtomCardContent, MolContentCard },
    props: {
      invite: {
        type: Object,
        required: true
      }
    }
  }
</script>

<style scoped>

</style>
