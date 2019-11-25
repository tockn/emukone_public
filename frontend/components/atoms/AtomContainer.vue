<template>
  <sui-container>
    <slot />
  </sui-container>
</template>

<script>
  import SuiContainer from 'semantic-ui-vue/dist/commonjs/elements/Container/Container'
  export default {
    name: 'AtomContainer',
    components: { SuiContainer }
  }
</script>

<style scoped>

</style>
