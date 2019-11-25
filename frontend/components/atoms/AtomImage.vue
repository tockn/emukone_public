<template>
  <sui-image
    :src="src"
    :class="klass"
  />
</template>

<script>
  import SuiImage from 'semantic-ui-vue/dist/commonjs/elements/Image/Image'
  export default {
    name: 'AtomImage',
    components: { SuiImage },
    props: {
      src: {
        type: String,
        required: true
      },
      shape: {
        type: String,
        default: ''
      },
      size: {
        type: String,
        default: ''
      }
    },
    computed: {
      klass () {
        let k = ''
        k += ` ${this.shape}`
        k += ` ${this.size}`
        return k
      }
    }
  }
</script>

<style scoped>

</style>
