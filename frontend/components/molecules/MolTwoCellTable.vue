<template>
  <div class="my-table">
    <div class="my-cell">
      <slot name="left" />
    </div>
    <div class="my-cell">
      <slot name="right" />
    </div>
  </div>
</template>

<script>
  export default {
    name: 'MolTwoCellTable'
  }
</script>

<style scoped>
  .my-table {
    margin: auto;
    display: table;
    table-layout: fixed;
  }
  .my-cell {
    display: table-cell;
    text-align: center;
    vertical-align: middle;
    width: 180px;
  }
</style>
