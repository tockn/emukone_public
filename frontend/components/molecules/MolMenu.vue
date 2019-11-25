<template>
  <atom-menu
    :fixed="fixed"
    :bottom="bottom"
  >
    <atom-menu-item
      v-for="(item, i) in items"
      :key="i"
      :content="item.name"
      :active="item.active"
      @click="item.action"
    />
  </atom-menu>
</template>

<script>
  import AtomMenu from '../atoms/AtomMenu'
  import AtomMenuItem from '../atoms/AtomMenuItem'
  export default {
    name: 'MolMenu',
    components: { AtomMenuItem, AtomMenu },
    props: {
      items: {
        type: Array, // ex: { name: 'hoge', active: true, action: function() }
        required: true
      },
      fixed: {
        type: Boolean,
        default: false
      },
      bottom: {
        type: Boolean,
        default: false
      }
    }
  }
</script>

<style scoped>

</style>
