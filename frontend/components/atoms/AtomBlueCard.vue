<template>
  <div class="wrap">
    <slot />
  </div>
</template>

<script>
  export default {
    name: 'AtomBlueCard'
  }
</script>

<style scoped>
  .wrap {
    margin: 8px;
    padding: 2px;
    background-color: #8CBEF8;
    color: #FFFFFF;
    text-align: center;
    -webkit-border-radius: 8px;
    -moz-border-radius: 8px;
    border-radius: 8px;
  }
</style>
