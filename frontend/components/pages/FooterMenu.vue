<template>
  <div>
    <atom-menu
      class="atom-menu"
      fixed
      bottom
      :widths="4"
      style="background-color: #3F3F3F"
    >
      <atom-menu-item
        :active="isSearchPage"
        @click="clickSearchPage"
      >
        <i
          class="fas fa-search"
          :style="`color: ${searchPageStyle}`"
        />
      </atom-menu-item>

      <atom-menu-item
        :active="isMyEventsPage"
        @click="clickMyEventsPage"
      >
        <i
          class="far fa-calendar-alt"
          :style="`color: ${myEventsPageStyle}`"
        />
      </atom-menu-item>

      <atom-menu-item
        :active="isInvitePage"
        @click="clickInvitePage"
      >
        <i
          class="fas fa-inbox"
          :style="`color: ${invitePageStyle}`"
        />
      </atom-menu-item>

      <atom-menu-item
        :active="isMyPage"
        @click="clickMyPage"
      >
        <mol-icon
          v-if="me.icon_url"
          :src="me.icon_url"
          :active="isMyPage"
        />
        <i
          v-else
          class="far fa-user"
          :style="`color: ${myPageStyle}`"
        />
      </atom-menu-item>
    </atom-menu>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import AtomMenu from '../atoms/AtomMenu'
  import AtomMenuItem from '../atoms/AtomMenuItem'
  import MolIcon from '../molecules/MolIcon'
  const inactive = '#9e9e9e'
  const active = '#ffffff'

  const noPage = -1
  const myPage = 0
  const invitePage = 1
  const searchPage = 2
  const myEventsPage = 3

  export default {
    name: 'FooterMenu',
    components: { MolIcon, AtomMenuItem, AtomMenu },
    data () {
      return {
        items: [{name: 'Mail'}],
        page: 0,
      }
    },
    computed: {
      ...mapState('auth', {
        me: state => state.me
      }),
      defaultMyPagePath () {
        return this.me ? `/${this.me.identifier}s/${this.me.id}` : ''
      },
      defaultInvitePagePath () {
        return '/invites'
      },
      defaultSearchPagePath () {
        return 'search'
      },
      defaultMyEventsPagePath () {
        return '/events'
      },
      isMyPage () {
        return this.page === myPage
      },
      isInvitePage () {
        return this.page === invitePage
      },
      isSearchPage () {
        return this.page === searchPage
      },
      isMyEventsPage () {
        return this.page === myEventsPage
      },

      myPageStyle () {
        return this.isMyPage ? active : inactive
      },
      invitePageStyle () {
        return this.isInvitePage ? active : inactive
      },
      searchPageStyle () {
        return this.isSearchPage ? active : inactive
      },
      myEventsPageStyle () {
        return this.isMyEventsPage ? active : inactive
      }
    },
    watch: {
      '$route' (to) {
        this.pageChangeHandler(to)
      }
    },
    created () {
      this.pageChangeHandler(this.$route)
      this.myPagePath = this.myPage
    },

    // TODO this.meのundefined判定いい感じにする

    methods: {
      clickMyPage () {
        if (this.me) {
          this.clickButtonHandler(myPage)
        }
      },
      clickMyEventsPage () {
        this.clickButtonHandler(myEventsPage)
      },
      clickInvitePage () {
        this.clickButtonHandler(invitePage)
      },
      clickSearchPage () {
        this.clickButtonHandler(searchPage)
      },
      clickButtonHandler (page) {
        let goPage = ''
        switch (page) {
          case myPage:
            goPage = this.defaultMyPagePath
            break
          case myEventsPage:
            goPage = this.defaultMyEventsPagePath
            break
          case invitePage:
            goPage = this.defaultInvitePagePath
            break
          case searchPage:
            goPage = this.defaultSearchPagePath
            break
        }
        this.$router.push(goPage)
        this.page = page
      },
      pageChangeHandler (route) {
        if (this._isMyPage(route)) this.page = myPage
        else if (this._isInvitePage(route)) this.page = invitePage
        else if (this._isSearchPage(route)) this.page = searchPage
        else if (this._isMyEventsPage(route)) this.page = myEventsPage
        else if (this._isSearchPage(route)) this.page = searchPage
        else this.page = noPage
      },
      _isMyPage (route) {
        if (this.me && (route.name === 'artists-id' || route.name === 'bookers-id')) {
          return route.params.id === this.me.id
        }
        return false
      },
      _isInvitePage (route) {
        return route.path === this.defaultInvitePagePath
      },
      _isSearchPage (route) {
        return route.path === this.defaultSearchPagePath
      },
      _isMyEventsPage (route) {
        return route.path === this.defaultMyEventsPagePath
      }
    },
  }
</script>

<style scoped>
  .atom-menu {
    height: 48px;
  }
  i {
    font-size: 24px;
  }
</style>
