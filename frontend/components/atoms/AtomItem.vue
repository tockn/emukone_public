<template>
  <sui-card class="my-card centered custom-width">
    <sui-image
      :src="src"
      size="tiny"
      class="icon-image"
      style="border-radius: 100px !important;"
    />
  </sui-card>
</template>

<script>
  import SuiCard from 'semantic-ui-vue/dist/commonjs/views/Card/Card'
  import SuiImage from 'semantic-ui-vue/dist/commonjs/elements/Image/Image'
  export default {
    name: 'AtomItem',
    components: { SuiImage, SuiCard },
    props: {
      src: {
        type: String,
        required: true
      }
    }
  }
</script>

<style scoped>
  .my-card {
    padding: 8px;
  }
  .icon-image {
    border-radius: 100px !important;
  }
</style>
