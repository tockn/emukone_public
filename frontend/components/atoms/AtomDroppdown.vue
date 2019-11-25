<template>
  <sui-dropdown
    v-model="innerValue"
    class="atom-dropdown"
    :options="options"
    :placeholder="placeholder"
    :multiple="multiple"
    :search="search"
    :allow-additions="allowAdditions"
    :fluid="fluid"
    :selection="selection"
  />
</template>

<script>
  import SuiDropdown from 'semantic-ui-vue/dist/commonjs/modules/Dropdown/Dropdown'
  export default {
    name: 'AtomDroppdown',
    components: { SuiDropdown },
    props: {
      options: {
        type: [Function, Array],
        required: true
      },
      value: {
        type: String,
        default: ''
      },
      placeholder: {
        type: String,
        default: ''
      },
      multiple: {
        type: Boolean,
        default: false
      },
      search: {
        type: Boolean,
        default: false
      },
      allowAdditions: {
        type: Boolean,
        default: false
      },
      fluid: {
        type: Boolean,
        default: false
      },
      selection: {
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
      }
    }
  }
</script>

<style scoped>
  .atom-dropdown {
    font-size: 16px !important;
  }
</style>
