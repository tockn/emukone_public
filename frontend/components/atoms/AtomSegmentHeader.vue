<template>
  <div class="text">
    <slot />
  </div>
</template>

<script>
  export default {
    name: 'AtomSegmentHeader'
  }
</script>

<style scoped>
  .text {
    font-weight: 700;
    font-size: 16px;
    text-align: left;
    color: rgba(0, 0, 0, 0.85);
    margin: 0 0 8px 0;
  }
</style>
