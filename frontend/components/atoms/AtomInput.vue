<template>
  <input
    v-model="innerValue"
    :class="klass"
    :placeholder="placeholder"
  >
</template>

<script>
  export default {
    name: 'AtomInput',
    props: {
      placeholder: {
        type: String,
        default: ''
      },
      value: {
        type: String,
        default: ''
      },
      white: {
        type: Boolean,
        default: false
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
      },
      klass () {
        return this.white ? ' white' : ''
      }
    }
  }
</script>

<style scoped>
  input {
    -webkit-border-radius: 0.285714em;
    -moz-border-radius: 0.285714em;
    border-radius: 0.285714em;
    color: rgba(0, 0, 0, 0.87);
    padding: 8px;
    font-size: 16px;
    -webkit-appearance: none;
    background-color: rgba(0, 0, 0, 0);
    text-align: center;
    margin: 8px 0;
    border: 1px solid rgba(34, 36, 38, 0.15);
  }
  input:focus {
    -webkit-box-shadow: 0 0 0 1px #8CBEF8;
    -moz-box-shadow: 0 0 0 1px #8CBEF8;
    box-shadow: 0 0 0 1px #8CBEF8;
    outline: none;
  }
  .white {
    border: 1px solid rgba(255, 255, 255, 0.8);
    color: #ffffff;
  }
</style>
