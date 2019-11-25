<template>
  <div>
    <atom-image
      :src="src"
      shape="circular"
      style="width: 28px !important;"
    />
    <div class="wrap-active">
      <i
        v-show="active"
        class="far fa-circle active-circle"
      />
    </div>
  </div>
</template>

<script>
  import AtomImage from '../atoms/AtomImage'
  export default {
    name: 'MolIcon',
    components: { AtomImage },
    props: {
      src: {
        type: String,
        required: true
      },
      active: {
        type: Boolean,
        default: false
      }
    }
  }
</script>

<style scoped>
  .active-circle {
    position: absolute;
    right: 27px;
    bottom: 3px;
    color: white;
    font-size: 40px;
  }
</style>
