<template>
  <textarea
    v-model="innerValue"
    :placeholder="placeholder"
  />
</template>

<script>
  export default {
    name: 'AtomTextarea',
    props: {
      placeholder: {
        type: String,
        default: ''
      },
      value: {
        type: String,
        required: true
      }
    },
    computed: {
      innerValue: {
        get() {
          return this.value
        },
        set(v) {
          this.$emit('input', v)
        }
      }
    }
  }
</script>

<style scoped>
  textarea {
    resize: none;
    border: 1px solid rgba(34, 36, 38, 0.15);
    -webkit-border-radius: 0.285714em;
    -moz-border-radius: 0.285714em;
    border-radius: 0.285714em;
    color: rgba(0, 0, 0, 0.87);
    padding: 8px;
    font-size: 16px;
    -webkit-appearance: none;
  }
  textarea:focus {
    -webkit-box-shadow: 0 0 0 1px #8CBEF8;
    -moz-box-shadow: 0 0 0 1px #8CBEF8;
    box-shadow: 0 0 0 1px #8CBEF8;
    outline: none;
  }
</style>
