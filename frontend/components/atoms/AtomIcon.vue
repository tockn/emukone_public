<template>
  <img :src="src">
</template>

<script>
  export default {
    name: 'AtomIcon',
    props: {
      src: {
        type: String,
        required: true
      }
    }
  }
</script>

<style scoped>
  img {
    width: 112px;
    height: 112px;
    object-fit: cover;
    -webkit-border-radius: 100px;
    -moz-border-radius: 100px;
    border-radius: 100px;
  }
</style>
