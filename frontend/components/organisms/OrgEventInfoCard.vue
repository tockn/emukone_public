<template>
  <atom-card>
    <p slot="header">
      開催情報
    </p>
    <div
      slot="description"
      style="text-align: center"
    >
      <table style="margin: auto">
        <tr>
          <td><span class="regular">DATE:</span></td>
          <td><span class="huge">{{ date }}</span></td>
        </tr>
        <tr>
          <td><span class="regular">OPEN:</span></td>
          <td><span class="huge">{{ formattedOpenTime }}</span></td>
        </tr>
        <tr>
          <td><span class="regular">START:</span></td>
          <td><span class="huge">{{ formattedStartTime }}</span></td>
        </tr>
        <tr>
          <td><span class="regular">PLACE:</span></td>
          <td><span class="huge">{{ event.place }}</span></td>
        </tr>
        <tr>
          <td><span class="regular">TICKET:</span></td>
          <td><span class="huge">{{ event.ticket }}</span></td>
        </tr>
      </table>
    </div>
  </atom-card>
</template>

<script>
  import { formatDate, formatTime } from '../../plugins/date-formatter'
  import AtomCard from '../atoms/AtomCard'
  export default {
    name: 'OrgEventInfoCard',
    components: { AtomCard },
    props: {
      event: {
        type: Object,
        required: true
      }
    },
    computed: {
      date () {
        return formatDate(this.event.start_time)
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
  p {
    margin: 0 !important;
  }
  h3 {
    margin: 4px auto;
  }
  .regular {
    font-size: 14px;
  }
  .huge {
    font-size: 18px;
    font-weight: bold;
  }
</style>
