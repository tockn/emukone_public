<template>
  <div>
    <AtomCard>
      <template slot="top">
        <slot name="top" />
        <atom-image
          v-if="imageSrc"
          :src="imageSrc"
        />
      </template>
      <template slot="header">
        {{ title }}
      </template>
      <div
        v-if="description"
        slot="description"
      >
        <pre v-text="description" />
      </div>
      <div slot="bottom">
        <slot name="bottom" />
      </div>
    </AtomCard>
  </div>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import AtomImage from '../atoms/AtomImage'
  export default {
    name: 'MolContentCard',
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
      imageSrc: {
        type: String,
        default: ''
      }
    }
  }
</script>

<style scoped>
  pre {
    white-space: pre-wrap;
  }
</style>
