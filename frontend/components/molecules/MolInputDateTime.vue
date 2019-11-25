<template>
  <atom-modal
    v-model="innerDisplay"
  >
    <p slot="header">
      {{ title }}
    </p>
    <div
      slot="content"
      style="text-align: center !important; margin: auto !important;"
    >
      <h3>
        Date
      </h3>
      <atom-date-picker
        v-model="date"
        open
        class="date-picker"
      />
      <div class="dummy-rectangle" />
      <h3
        style="margin-top: 0 !important;"
      >
        Time
      </h3>
      <atom-time-picker
        v-model="time"
      />
    </div>
    <atom-button
      slot="action"
      color="green"
      @click="clickCloseModal"
    >
      OK
    </atom-button>
  </atom-modal>
</template>

<script>
  import AtomDatePicker from '../atoms/AtomDatePicker'
  import AtomTimePicker from '../atoms/AtomTimePicker'
  import AtomModal from '../atoms/AtomModal'
  import AtomButton from '../atoms/AtomButton'
  export default {
    name: 'OrgInputDateTime',
    components: { AtomButton, AtomModal, AtomTimePicker, AtomDatePicker },
    props: {
      value: {
        type: [Date, Object, String],
        default: ''
      },
      title: {
        type: String,
        required: true
      },
      display: {
        type: Boolean,
        required: true
      },
      closeModalMutation: {
        type: Function,
        required: true
      }
    },
    data () {
      return {
        date: new Date(),
        time: '18:30'
      }
    },
    computed: {
      innerDisplay: {
        get () {
          return this.display
        },
        set (v) {
          this.closeModalMutation(v)
        }
      }
    },
    watch: {
      date () {
        this.handleChange()
      },
      time () {
        this.handleChange()
      }
    },
    mounted () {
      this.handleChange()
    },
    methods: {
      handleChange () {
        let d = new Date(this.date.getTime())
        let hour = Number(this.time.split(':')[0])
        let min = Number(this.time.split(':')[1])
        d.setHours(hour)
        d.setMinutes(min)
        this.$emit('input', d)
      },
      clickCloseModal () {
        this.innerDisplay = false
        this.handleChange()
      }
    }
  }
</script>

<style scoped>
  .dummy-rectangle {
    width: 360px;
    height: 300px;
  }
</style>
