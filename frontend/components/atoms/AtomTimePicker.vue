<template>
  <div
    style="text-align: center"
  >
    <label
      for="time-picker"
    >
      <input
        id="time-picker"
        v-model="innerValue"
        type="time"
      >
    </label>
  </div>
</template>

<script>
  export default {
    name: 'AtomTimePicker',
    props: {
      value: {
        type: [Object, Date, String],
        default: () => {}
      }
    },
    computed: {
      innerValue: {
        get () {
          return this.value
        },
        set (v) {
          this.$emit('input', v)
        }
      }
    }
  }

</script>

<style scoped>
  label {
    position: relative;
    display: inline-block;
    width: 200px;
    height: 36px;
    border: 2px solid #ccc;
    border-radius: 8px;
  }
  input[type="time"] {
    text-align: center !important;
    position: relative;
    padding: 0 10px;
    width: 200px;
    height: 36px;
    border: 0;
    background: transparent;
    box-sizing: border-box;
    font-size: 16px;
  }
</style>
