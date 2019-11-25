<template>
  <div>
    <div class="ui stackable grid container centered">
      <mol-content-small-card
        v-for="(invite, i) in invites"
        :key="i"
        :title="invite.artist.name"
        :img-src="invite.artist.icon_url"
        :link-to="`/artists/${invite.artist.id}`"
        style="width: 200px !important;"
      >
        <atom-label
          slot="bottom"
          :color="statusColor(invite.status)"
          style="transform:translateX(-2px); width: 100%"
        >
          {{ statusText(invite.status) }}
        </atom-label>
      </mol-content-small-card>
      <p style="margin: 16px auto;" v-show="invites.length === 0">
        まだ誰もスカウトしていません。
      </p>
    </div>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import MolContentSmallCard from '../molecules/MolContentSmallCard'
  import AtomLabel from '../atoms/AtomLabel'
  export default {
    name: 'OrgInviteStatuses',
    components: { AtomLabel, MolContentSmallCard },
    props: {
      invites: {
        type: Array,
        required: true
      }
    },
    computed: {
      ...mapGetters('invites',[
        'unchecked',
        'approved',
        'declined'
      ])
    },
    methods: {
      statusText (status) {
        switch (status) {
          case this.unchecked:
            return '未チェック'
          case this.approved:
            return '承認'
          case this.declined:
            return '辞退'
        }
      },
      statusColor (status) {
        switch (status) {
          case this.unchecked:
            return 'grey'
          case this.approved:
            return 'green'
          case this.declined:
            return 'red'
        }
      }
    }
  }
</script>

<style scoped>

</style>
