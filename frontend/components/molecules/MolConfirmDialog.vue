<template>
  <atom-modal
    v-model="innerValue"
  >
    <p slot="header">
      {{ title }}
    </p>
    <slot
      slot="description"
      name="description"
    />
    <atom-button-group slot="action">
      <atom-button
        style="width: 128px !important;"
        basic
        positive
        @click="clickApprove"
      >
        {{ approveButtonText }}
      </atom-button>
      <atom-button
        style="width: 128px !important;"
        basic
        negative
        @click="clickDecline"
      >
        {{ declineButtonText }}
      </atom-button>
    </atom-button-group>
  </atom-modal>
</template>

<script>
  import AtomModal from '../atoms/AtomModal'
  import AtomButton from '../atoms/AtomButton'
  import AtomButtonGroup from '../atoms/AtomButtonGroup'
  export default {
    name: 'MolConfirmDialog',
    components: { AtomButtonGroup, AtomModal, AtomButton },
    props: {
      value: {
        type: Boolean,
        default: false
      },
      title: {
        type: String,
        required: true
      },
      moveAction: {
        type: Function,
        required: true
      },
      notMoveAction: {
        type: Function,
        default: () => {}
      },
      approveButtonText: {
        type: String,
        default: 'OK'
      },
      declineButtonText: {
        type: String,
        default: 'キャンセル'
      }
    },
    computed: {
      innerValue: {
        get () {
          return this.value
        },
        set (v) {
          this.notMoveAction()
          this.$emit('input', v)
        }
      }
    },
    methods: {
      clickApprove () {
        this.moveAction()
        this.$emit('input', false)
      },
      clickDecline () {
        this.notMoveAction()
        this.$emit('input', false)
      }
    }
  }
</script>

<style scoped>

</style>
