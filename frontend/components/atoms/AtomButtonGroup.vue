<template>
  <sui-button-group>
    <slot />
  </sui-button-group>
</template>

<script>
  import SuiButtonGroup from 'semantic-ui-vue/dist/commonjs/elements/Button/ButtonGroup'
  export default {
    name: 'AtomButtonGroup',
    components: { SuiButtonGroup }
  }
</script>

<style scoped>

</style>
