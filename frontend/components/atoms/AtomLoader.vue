<template>
  <div class="my-loader">
    <sui-loader
      :content="content"
      active
    />
  </div>
</template>

<script>
  import SuiLoader from 'semantic-ui-vue/dist/commonjs/elements/Loader/Loader'
  export default {
    name: 'AtomLoader',
    components: { SuiLoader },
    props: {
      content: {
        type: String,
        default: ''
      }
    }
  }
</script>

<style scoped>
  .my-loader {
    padding: 32px !important;
  }
</style>
