<template>
  <date-picker
    ref="picker"
    v-model="innerValue"
    input-class="data-picker-global-class"
    wrapper-class="data-picker-global-wrapper-class"
    @closed="closeHandler"
  />
</template>

<script>
  import DatePicker from 'vuejs-datepicker'
  export default {
    name: 'AtomDatePicker',
    components: { DatePicker },
    props: {
      value: {
        type: [Object, Date, String],
        default: ''
      },
      open: {
        type: Boolean,
        required: true
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
    },
    watch: {
      open () {
        if (this.open) {
          this.$refs.picker.setInitialView()
        } else {
          this.$refs.picker.close(true)
        }
      }
    },
    mounted () {
      if (this.open) this.$refs.picker.setInitialView()
    },
    methods: {
      closeHandler () {
        this.$refs.picker.setInitialView()
      }
    }
  }
</script>

<style>
  .data-picker-global-class {
    display: none;
  }
  .data-picker-global-wrapper-class {
    position: relative;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    width: 304px;
    margin-left: auto;
    margin-right: auto;
  }
</style>
