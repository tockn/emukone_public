<template>
  <sui-label
    :color="color"
    :class="klass"
  >
    <slot />
  </sui-label>
</template>

<script>
  import SuiLabel from 'semantic-ui-vue/dist/commonjs/elements/Label/Label'
  export default {
    name: 'AtomLabel',
    components: { SuiLabel },
    props: {
      ribbon: {
        type: String,
        default: ''
      },
      color: {
        type: String,
        default: ''
      }
    },
    computed: {
      klass () {
        let k = ''
        k += this.ribbon ? ` ${this.ribbon}` : ''
        return k
      }
    }
  }
</script>

<style scoped>

</style>
