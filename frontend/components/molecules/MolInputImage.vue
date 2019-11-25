<template>
  <div>
    <atom-input-file @change="selected">
      <template slot="content">
        <span>
          <i class="far fa-image" />
        </span>
        <p>画像を変更</p>
      </template>
    </atom-input-file>

    <atom-icon :src="value" />

    <atom-modal v-model="modalFlags">
      <p slot="header">
        画像を編集
      </p>

      <mol-cropper
        slot="content"
        v-model="cro"
        :initial-image="innerSrc"
      />
      <atom-button
        slot="action"
        @click="change"
      >
        OK
      </atom-button>
    </atom-modal>
  </div>
</template>

<script>
  import AtomIcon from '../atoms/AtomIcon'
  import ImageResizer from '../../plugins/image-resizer'
  import AtomModal from '../atoms/AtomModal'
  // import Croppa from 'vue-croppa'
  import AtomButton from '../atoms/AtomButton'
  import AtomInputFile from '../atoms/AtomInputFile'
  import MolCropper from './MolCropper'
  export default {
    name: 'MolInputImage',
    components: { MolCropper, AtomInputFile, AtomModal, AtomIcon, AtomButton },
    props: {
      value: {
        type: String,
        required: true
      },
      imageLoading: {
        type: Boolean,
        default: false
      }
    },
    data () {
      return {
        innerSrc: '',
        cro: {},
        modalFlag: false
      }
    },
    computed: {
      modalFlags: {
        get () {
          return this.imageLoading || this.modalFlag
        },
        set (v) {
          this.modalFlag = v
        }
      }
    },
    methods: {
      selected (e) {
        e.preventDefault()
        let file = e.target.files[0]
        if (!file) return
        ImageResizer(file, (base64) => {
          this.innerSrc = base64
          this.modalFlag = true
        })
      },
      change () {
        this.$emit('input', this.cro.generateDataUrl())
        this.modalFlag = false
      }
    }
  }
</script>

<style scoped>
  span {
    font-size: 48px;
    margin: auto;
    vertical-align: middle;
  }
</style>
