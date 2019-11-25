<template>
  <sui-menu
    :class="klass"
    :widths="widths"
  >
    <slot />
  </sui-menu>
</template>

<script>
  import SuiMenu from 'semantic-ui-vue/dist/commonjs/collections/Menu/Menu'
  export default {
    name: 'AtomMenu',
    components: { SuiMenu },
    props: {
      widths: {
        type: Number,
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
    },
    computed: {
      klass () {
        let k = ''
        k += this.bottom ? ' bottom' : ''
        k += this.fixed ? ' fixed' : ''
        return k
      }
    }
  }
</script>

<style scoped>

</style>
