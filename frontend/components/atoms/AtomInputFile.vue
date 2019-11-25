<template>
  <label
    for="image-file"
    :class="shape"
  >
    <slot name="content" />
    <input
      id="image-file"
      accept="image/*"
      type="file"
      style="display: none;"
      @change="change"
    >
  </label>
</template>

<script>
  export default {
    name: 'AtomInputFile',
    props: {
      shape: {
        type: String,
        default: 'circular'
      }
    },
    methods: {
      change (e) {
        this.$emit('change', e)
      }
    }
  }
</script>

<style scoped>
  .circular {
    padding-top: 20px;
    -webkit-border-radius: 100px;
    -moz-border-radius: 100px;
    border-radius: 100px;
    position: absolute;
    width: 112px;
    height: 112px;
    background-color: rgba(127, 130, 139, 0.6);
  }
  .rectangle {
    padding-top: 20px;
    border-radius: 8px;
    position: absolute;
    width: 100%;
    height: 360px;
    background-color: rgba(127, 130, 139, 0.6);
  }
</style>
