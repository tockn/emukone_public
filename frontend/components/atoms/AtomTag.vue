<template>
  <!--TODO nuxt-linkにする-->
  <a
    is="sui-label"
    class="small tag atomtag"
    @click="$emit('click')"
  >{{ name }}</a>
</template>

<script>
  export default {
    name: 'AtomTag',
    props: {
      name: {
        type: String,
        required: true
      }
    }
  }
</script>

<style scoped>
  .atomtag {
    margin-top: 0 !important;
    margin-bottom: 4px !important;
  }
</style>
