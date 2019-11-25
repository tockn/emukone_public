<template>
  <sui-card-content>
    <slot />
  </sui-card-content>
</template>

<script>
  import SuiCardContent from 'semantic-ui-vue/dist/commonjs/views/Card/CardContent'
  export default {
    name: 'AtomCardContent',
    components: { SuiCardContent }
  }
</script>

<style scoped>

</style>
