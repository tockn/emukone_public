<template>
  <div>
    <footer-menu />
    <nuxt class="nuxt" />
  </div>
</template>

<script>
  import FooterMenu from '../components/pages/FooterMenu'
  export default {
    components: { FooterMenu }
  }
</script>

<style>
  .nuxt {
    padding-bottom: 160px;
  }
</style>
