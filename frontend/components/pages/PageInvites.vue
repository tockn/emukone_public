<template>
  <div>
    <atom-card>
      <h3
        slot="header"
        style="text-align: center"
      >
        スカウトされているイベント一覧
      </h3>
    </atom-card>
    <atom-menu
      :widths="2"
    >
      <atom-menu-item
        :active="isDisplayUnchecked"
        @click="clickUnchecked"
      >
        未承認のスカウト
      </atom-menu-item>
      <atom-menu-item
        :active="isDisplayAll"
        @click="clickAll"
      >
        全てのスカウト
      </atom-menu-item>
    </atom-menu>
    <div v-if="filteredInvites.length > 0">
      <nuxt-link
        v-for="(invite, i) in filteredInvites"
        :key="i"
        :to="`/invites/${invite.event.id}`"
      >
        <org-invite-large-card
          :invite="invite"
        />
      </nuxt-link>
    </div>
    <div v-else>
      <atom-card
        style="text-align: center;"
      >
        <p
          v-show="isDisplayAll"
          slot="header"
        >
          まだスカウトはありません。プロフィールを充実させ、スカウトを獲得しましょう！
        </p>
        <p
          v-show="isDisplayUnchecked"
          slot="header"
        >
          未承認のスカウトはありません
        </p>
      </atom-card>
    </div>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import AtomCard from '../atoms/AtomCard'
  import OrgInviteLargeCard from '../organisms/OrgInviteLargeCard'
  import AtomMenu from '../atoms/AtomMenu'
  import AtomMenuItem from '../atoms/AtomMenuItem'
  const displayUnchecked = 0
  const displayAll = 1
  export default {
    name: 'PageInvites',
    components: { AtomMenuItem, AtomMenu, OrgInviteLargeCard, AtomCard },
    props: {
      invites: {
        type: Array,
        required: true
      }
    },
    data () {
      return {
        state: displayUnchecked
      }
    },
    computed: {
      ...mapGetters('invites', [
        'unchecked'
      ]),
      isDisplayUnchecked () {
        return this.state === displayUnchecked
      },
      isDisplayAll () {
        return this.state === displayAll
      },
      filteredInvites () {
        return this.isDisplayAll ? this.invites
          : this.invites.filter(invite => invite.status === 'unchecked')
      }
    },
    methods: {
      clickUnchecked () {
        this.state = displayUnchecked
      },
      clickAll () {
        this.state = displayAll
      }
    }
  }
</script>

<style scoped>
  p {
    font-weight: normal;
  }
</style>
