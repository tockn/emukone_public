<template>
  <a
    is="sui-menu-item"
    :active="active"
    @click="$emit('click')"
  >
    <slot />
  </a>
</template>

<script>
  export default {
    name: 'AtomMenuItem',
    props: {
      active: {
        type: Boolean,
        required: true
      }
    }
  }
</script>

<style scoped>

</style>
