<template>
  <div>
    <org-profile-header :user="artist">
      <nuxt-link
        v-if="me.id === artist.id"
        to="/settings/profile"
      >
        <atom-button>
          編集する
        </atom-button>
      </nuxt-link>
      <atom-button
        v-else
        color="green"
      >
        スカウトする
      </atom-button>
    </org-profile-header>
    <mol-content-card
      title="活動に懸ける想い"
      :description="artist.why_description"
    />
    <mol-music-card
      v-if="artist.ichioshis[0]"
      :ichioshi="artist.ichioshis[0]"
    />
    <mol-content-card
      title="詳細"
      :description="artist.free_description"
    />
    <org-horizontal-scroll-event-card
      title="今まで関わったイベント"
      :events="events"
    />
  </div>
</template>

<script>
  import MolContentCard from '../molecules/MolContentCard'
  import MolMusicCard from '../molecules/MolMusicCard'
  import OrgHorizontalScrollEventCard from '../organisms/OrgHorizontalScrollEventCard'
  import OrgProfileHeader from '../organisms/OrgProfileHeader'
  import AtomButton from '../atoms/AtomButton'
  import { mapState } from 'vuex'
  export default {
    name: 'PageArtistProfile',
    components: { OrgProfileHeader, OrgHorizontalScrollEventCard, MolMusicCard, MolContentCard, AtomButton},
    props: {
      artist: {
        type: Object,
        required: true
      },
      events: {
        type: Array,
        required: true
      }
    },
    computed: {
      ...mapState({
        me: state => state.auth.me
      })
    }
  }
</script>

<style scoped>
  pre {
    white-space: pre-wrap;
  }
</style>
