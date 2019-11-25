<template>
  <sui-button
    :class="`${klass}`"
    @click="$emit('click')"
  >
    <slot />
  </sui-button>
</template>

<script>
  import SuiButton from 'semantic-ui-vue/dist/commonjs/elements/Button/Button'
  export default {
    name: 'AtomButton',
    components: { SuiButton },
    props: {
      color: {
        type: String,
        default: ''
      },
      loading: {
        type: Boolean,
        default: false
      },
      basic: {
        type: Boolean,
        default: false
      },
      positive: {
        type: Boolean,
        default: false
      },
      negative: {
        type: Boolean,
        default: false
      }
    },
    computed: {
      klass () {
        let klass = ''
        klass += this.color
        klass += this.loading ? ' loading' : ''
        klass += this.basic ? ' basic' : ''
        klass += this.positive ? ' positive' : ''
        klass += this.negative ? ' negative' : ''
        return klass
      }
    }
  }
</script>

<style scoped>

</style>
