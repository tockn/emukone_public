<template>
  <atom-modal v-model="innerValue">
    <h3
      slot="header"
      style="text-align: center"
    >
      {{ title }}
    </h3>
    <div
      slot="description"
      style="text-align: center"
    >
      <slot name="description" />
    </div>
    <atom-button-group slot="action">
      <atom-button
        basic
        positive
        @click="positiveClick"
      >
        {{ approveButtonText }}
      </atom-button>
      <atom-button
        basic
        negative
        @click="$emit('input', false)"
      >
        {{ declineButtonText }}
      </atom-button>
    </atom-button-group>
  </atom-modal>
</template>

<script>
  import AtomButtonGroup from '../atoms/AtomButtonGroup'
  import AtomButton from '../atoms/AtomButton'
  import AtomModal from '../atoms/AtomModal'
  export default {
    name: 'OrgConfirmInviteModal',
    components: { AtomModal, AtomButtonGroup, AtomButton },
    props: {
      title: {
        type: String,
        required: true
      },
      approveButtonText: {
        type: String,
        required: true
      },
      declineButtonText: {
        type: String,
        required: true
      },
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
    },
    methods: {
      positiveClick () {
        this.$emit('input', false)
        this.$emit('click')
      }
    }
  }
</script>

<style scoped>

</style>
