<template>
  <div>
    <h3 style="text-align: center">
      DATE: {{ date }}
    </h3>

    <mol-two-cell-table>
      <h3 slot="left">
        OPEN: {{ formattedOpenTime }}
      </h3>
      <atom-button
        slot="right"
        @click="$emit('click-edit-open')"
      >
        編集
      </atom-button>
    </mol-two-cell-table>

    <mol-two-cell-table>
      <h3 slot="left">
        START: {{ formattedStartTime }}
      </h3>
      <atom-button
        slot="right"
        @click="$emit('click-edit-start')"
      >
        編集
      </atom-button>
    </mol-two-cell-table>
  </div>
</template>

<script>
  import { formatDate, formatTime } from '../../plugins/date-formatter'
  import AtomButton from '../atoms/AtomButton'
  import MolTwoCellTable from '../molecules/MolTwoCellTable'
  export default {
    name: 'OrgInputEventDate',
    components: { MolTwoCellTable, AtomButton },
    props: {
      event: {
        type: Object,
        default: () => {}
      }
    },
    computed: {
      date () {
        return formatDate(this.event.open_time)
      },
      formattedStartTime () {
        return formatTime(this.event.start_time)
      },
      formattedOpenTime () {
        return formatTime(this.event.open_time)
      }
    }
  }
</script>

<style scoped>
  h3 {
    margin-top: 12px;
    margin-bottom: 12px;
  }
</style>
