<template>
  <sui-segment>
    <slot />
  </sui-segment>
</template>

<script>
  import SuiSegment from 'semantic-ui-vue/dist/commonjs/elements/Segment/Segment'
  export default {
    name: 'AtomSegment',
    components: { SuiSegment }
  }
</script>

<style scoped>

</style>
