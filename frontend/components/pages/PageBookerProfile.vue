<template>
  <div>
    <org-profile-header :user="booker">
      <nuxt-link
        v-if="me.id === booker.id"
        to="/settings/profile"
      >
        <atom-button>
          編集する
        </atom-button>
      </nuxt-link>
    </org-profile-header>
    <mol-content-card
      title="活動に懸ける想い"
      :description="booker.why_description"
    />
    <mol-content-card
      title="こんなことやっています"
      :description="booker.how_description"
    />
    <org-horizontal-scroll-event-card
      title="今まで関わったイベント"
      :events="events"
    />
    <mol-content-card
      title="詳細"
      :description="booker.free_description"
    />
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import OrgProfileHeader from '../organisms/OrgProfileHeader'
  import MolContentCard from '../molecules/MolContentCard'
  import OrgHorizontalScrollEventCard from '../organisms/OrgHorizontalScrollEventCard'
  import AtomButton from '../atoms/AtomButton'
  export default {
    name: 'PageBookerProfile',
    components: { OrgHorizontalScrollEventCard, MolContentCard, OrgProfileHeader, AtomButton },
    props: {
      booker: {
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
</style>
