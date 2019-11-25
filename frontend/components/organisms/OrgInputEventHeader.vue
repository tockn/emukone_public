<template>
  <div>
    <atom-input-file
      shape="rectangle"
      @change="selected"
    >
      <div
        v-if="loading"
        slot="content"
      >
        <atom-loader />
      </div>
      <div
        v-else
        slot="content"
        class="icon-wrap"
      >
        <span>
          <i class="far fa-image" />
        </span>
        <p>画像を変更</p>
      </div>
    </atom-input-file>

    <atom-header-image
      :src="value"
      :filter="false"
      :max-height="true"
      :absolute="false"
      style="height: 360px"
    />
    <atom-modal v-model="modalFlag">
      <p slot="header">
        画像を編集
      </p>

      <mol-cropper
        slot="content"
        v-model="cro"
        :initial-image="innerSrc"
        :width="280"
        :height="280"
      />
      <atom-button
        slot="action"
        color="green"
        :loading="postLoading"
        @click="uploadHeader"
      >
        OK
      </atom-button>
    </atom-modal>
  </div>
</template>

<script>
  import ImageResizer from '../../plugins/image-resizer'
  import AtomHeaderImage from '../atoms/AtomHeaderImage'
  import MolCropper from '../molecules/MolCropper'
  import AtomModal from '../atoms/AtomModal'
  import AtomButton from '../atoms/AtomButton'
  import AtomInputFile from '../atoms/AtomInputFile'
  import AtomLoader from '../atoms/AtomLoader'
  const API_ENDPOINT = process.env.API_ENDPOINT
  export default {
    name: 'OrgInputEventHeader',
    components: { AtomLoader, AtomInputFile, AtomButton, AtomModal, MolCropper, AtomHeaderImage },
    props: {
      value: {
        type: String,
        default: ''
      },
      loading: {
        type: Boolean,
        default: false
      }
    },
    data () {
      return {
        modalFlag: false,
        innerSrc: '',
        cro: {},
        postLoading: false
      }
    },
    methods: {
      selected(e) {
        e.preventDefault()
        let file = e.target.files[0]
        if (!file) return
        ImageResizer(file, (base64) => {
          this.innerSrc = base64
          this.modalFlag = true
        })
      },
      uploadHeader () {
        let file = this.cro.generateDataUrl()
        this.postLoading = true
        ImageResizer(file, async (base64) => {
          base64 = base64.replace(/^.*,/, '')
          const res = await this.$axios.$post(`${API_ENDPOINT}/statics/images`,
            {
              content_type: file.type,
              base64: base64,
            })
          this.postLoading = false
          this.$emit('input', res.url)
          this.modalFlag = false
        })
      }
    }
  }
</script>

<style scoped>
  .icon-wrap {
    height: 360px;
    padding-top: 124px;
    text-align: center;
    color: white;
  }
  .icon-wrap i {
    margin: auto;
    vertical-align: middle;
    font-size: 48px;
  }
</style>
