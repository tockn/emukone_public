<template>
  <sui-modal
    v-model="innerValue"
    class="atom-modal"
  >
    <sui-modal-header>
      <slot name="header" />
    </sui-modal-header>

    <sui-modal-content>
      <slot name="content" />

      <sui-modal-description>
        <slot name="description" />
      </sui-modal-description>
    </sui-modal-content>

    <sui-modal-actions style="text-align: center !important;">
      <slot name="action" />
    </sui-modal-actions>
  </sui-modal>
</template>

<script>
  import SuiModalHeader from 'semantic-ui-vue/dist/commonjs/modules/Modal/ModalHeader'
  import SuiModalContent from 'semantic-ui-vue/dist/commonjs/modules/Modal/ModalContent'
  import SuiModalActions from 'semantic-ui-vue/dist/commonjs/modules/Modal/ModalActions'
  import SuiModal from 'semantic-ui-vue/dist/commonjs/modules/Modal/Modal'
  import SuiModalDescription from 'semantic-ui-vue/dist/commonjs/modules/Modal/ModalDescription'
  export default {
    name: 'AtomModal',
    components: { SuiModalDescription, SuiModal, SuiModalActions, SuiModalContent, SuiModalHeader },
    props: {
      value: {
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
  .atom-modal {
    color: black;
  }
</style>
