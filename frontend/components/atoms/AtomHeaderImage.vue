<template>
  <div>
    <img
      :class="klass"
      :src="src"
    >
  </div>
</template>

<script>
  export default {
    name: 'AtomHeaderImage',
    props: {
      src: {
        type: String,
        required: true
      },
      filter: {
        type: Boolean,
        default: true
      },
      maxHeight: {
        type: Boolean,
        default: false
      },
      absolute: {
        type: Boolean,
        default: true
      }
    },
    computed: {
      klass () {
        let klass = 'background'
        klass += (this.filter ? ' filter' : '')
        klass += (this.maxHeight ? ' max-height' : '')
        klass += (this.absolute ? ' absolute' : '')
        return klass
      }
    }
  }
</script>

<style scoped>
  .background {
    width: 100%;
    height: 360px;
    object-fit: cover;
  }
  .filter {
    filter: contrast(60%) blur(16px);
  }
  .max-height {
    height: auto;
    max-height: 360px;
  }
  .absolute {
    position: absolute;
    z-index: -1;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    margin-left: auto;
    margin-right: auto;
  }
</style>
