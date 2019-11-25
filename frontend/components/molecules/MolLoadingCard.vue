<template>
  <atom-card class="wrap">
    <atom-loader
      slot="header"
    />
  </atom-card>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import AtomLoader from '../atoms/AtomLoader'
  export default {
    name: 'MolLoadingCard',
    components: { AtomLoader, AtomCard }
  }
</script>

<style scoped>
</style>
