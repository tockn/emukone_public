<template>
  <div slot="content">
    <croppa-component
      v-model="croppa"
      class="croppa"
      :show-remove-button="false"
      :initial-image="initialImage"
      :width="width"
      :height="height"
    >
      <br>
      <atom-button @click="croppa.rotate(-1)">
        <i class="fas fa-undo" />
      </atom-button>
      <atom-button @click="croppa.rotate(1)">
        <i class="fas fa-redo" />
      </atom-button>
    </croppa-component>
  </div>
</template>

<script>
  import Croppa from 'vue-croppa'
  import AtomButton from '../atoms/AtomButton'
  export default {
    name: 'MolCropper',
    components: { AtomButton, croppaComponent: Croppa.component },
    props: {
      value: {
        type: Object,
        default: () => {}
      },
      initialImage: {
        type: String,
        default: ''
      },
      width: {
        type: Number,
        default: 200
      },
      height: {
        type: Number,
        default: 200
      }
    },
    data () {
      return {
        croppa: {}
      }
    },
    watch: {
      'initialImage'() {
        this.croppa.refresh()
      },
      'croppa'(v) {
        this.$emit('input', v)
      }
    }
  }
</script>

<style scoped>
  .croppa {
    text-align: center;
    margin: auto;
  }
</style>
