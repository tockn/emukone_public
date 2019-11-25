<template>
  <atom-card>
    <template slot="header">
      {{ title }}
    </template>
    <template
      v-if="loading"
      slot="description"
    >
      <atom-loader />
    </template>
    <template
      v-else
      slot="description"
    >
      <atom-textarea
        v-model="innerValue"
        class="input"
        :placeholder="placeholder"
      />
    </template>
  </atom-card>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import AtomTextarea from '../atoms/AtomTextarea'
  import AtomLoader from '../atoms/AtomLoader'
  export default {
    name: 'MolInputCard',
    components: { AtomLoader, AtomTextarea, AtomCard },
    props: {
      title: {
        type: String,
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
      loading: {
        type: Boolean,
        default: false
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
  .input {
    width: 100%;
    height: 180px;
    vertical-align: top !important;
  }
</style>
