<template>
  <div>
    <input
      type="file"
      @change="selected"
    >
    <croppa
      v-model="cro"
      :image-border-radius="250"
      prevent-white-space
    >
      <img
        slot="initial"
        :src="imgSrc"
      >
    </croppa>
    <div class="btns">
      <button @click="cro.rotate()">
        rotate 90deg
      </button>
      <button @click="cro.rotate(2)">
        rotate 180deg
      </button>
      <button @click="cro.rotate(-1)">
        rotate -90deg
      </button>
      <button @click="cro.flipX()">
        flip horizontally
      </button>
      <button @click="cro.flipY()">
        flip vertically
      </button>
      <button @click="gen">
        gen
      </button>
      <button @click="rot">
        rot
      </button>
    </div>
    <img :src="img">
  </div>
</template>

<script>
  import Croppa from 'vue-croppa'
  import ImageResizer from '../plugins/image-resizer'
  export default {
    components: {croppa: Croppa.component},
    data () {
      return {
        cro: {},
        img: "",
        imgSrc: ""
      }
    },
    methods: {
      gen: function () {
        this.img =this.cro.generateDataUrl()
      },
      rot: function () {
        let ctx = this.cro.getCanvas().getContext('2d')
        ctx.rotate(90)
      },
      selected (e) {
        e.preventDefault()
        let file = e.target.files[0]
        if (!file) return
        ImageResizer(file, (blob) => {
          this.imgSrc = blob
        })
      }
    }
  }
</script>
