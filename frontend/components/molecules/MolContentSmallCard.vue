<template>
  <nuxt-link
    class="wrap"
    :to="linkTo"
  >
    <atom-card :no-description="!description">
      <template slot="top">
        <atom-image :src="imgSrc" />
      </template>
      <template slot="header">
        {{ title }}
      </template>
      <template slot="description">
        <pre v-text="description" />
      </template>
      <template slot="bottom">
        <slot name="bottom" />
      </template>
    </atom-card>
  </nuxt-link>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import AtomImage from '../atoms/AtomImage'
  export default {
    name: 'MolContentSmallCard',
    components: { AtomImage, AtomCard },
    props: {
      title: {
        type: String,
        required: true
      },
      description: {
        type: String,
        default: ''
      },
      imgSrc: {
        type: String,
        required: true
      },
      linkTo: {
        type: String,
        default: ''
      }
    }
  }
</script>

<style scoped>
  .wrap {
    width: 144px !important;
  }
  img {
    width: 144px;
    height: 144px;
    object-fit: cover;
  }
</style>

