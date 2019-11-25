<template>
  <div>
    <org-input-event-header
      v-model="event.header_image_url"
      :loading="eventLoading"
    />

    <mol-loading-card v-show="eventLoading" />
    <atom-card
      v-show="!eventLoading"
      style="text-align: center;"
    >
      <atom-input
        slot="header"
        v-model="event.name"
        placeholder="イベント名"
        class="name-input"
      />
      <atom-input
        slot="description"
        v-model="event.meta_description"
        placeholder="ひとこと"
        class="meta-input"
      />
    </atom-card>

    <mol-input-card
      v-model="event.why_description"
      title="イベントに懸ける想い"
      :loading="eventLoading"
    />

    <atom-card>
      <p slot="header">
        開催情報
      </p>
      <atom-loader
        v-if="eventLoading"
        slot="description"
      />
      <div
        v-else
        slot="description"
      >
        <org-input-event-date
          :event="event"
          @click-edit-open="openTimeModal = true"
          @click-edit-start="startTimeModal = true"
        />
        <mol-two-cell-table>
          <h3 slot="left">
            場所:
          </h3>
          <atom-input
            slot="right"
            v-model="event.place"
          />
        </mol-two-cell-table>
        <mol-two-cell-table>
          <h3 slot="left">
            チケット情報:
          </h3>
          <atom-input
            slot="right"
            v-model="event.ticket"
          />
        </mol-two-cell-table>
      </div>
    </atom-card>

    <mol-input-date-time
      v-model="event.open_time"
      title="開場時間を入力"
      :display="openTimeModal"
      :close-modal-mutation="openTimeModalMutation"
    />
    <mol-input-date-time
      v-model="event.start_time"
      title="開演時間を入力"
      :display="startTimeModal"
      :close-modal-mutation="startTimeModalMutation"
    />

    <atom-card>
      <div
        slot="description"
        style="text-align: center"
      >
        <div
          v-show="errorMsgs.length > 0"
          style="margin-bottom: 16px;"
        >
          <p>以下の入力エラーがあります...！</p>
          <p
            v-for="(msg, i) in errorMsgs"
            :key="i"
            style="color: red"
          >
            {{ msg }}
          </p>
        </div>

        <atom-button
          color="green"
          :loading="loading.post"
          @click="clickPost"
        >
          イベントを投稿
        </atom-button>
      </div>
    </atom-card>

    <org-move-confirm
      ref="confirm"
    />
  </div>
</template>

<script>
  import { mapActions } from 'vuex'
  import OrgInputEventHeader from '../organisms/OrgInputEventHeader'
  import AtomCard from '../atoms/AtomCard'
  import AtomInput from '../atoms/AtomInput'
  import MolInputCard from '../molecules/MolInputCard'
  import MolInputDateTime from '../molecules/MolInputDateTime'
  import OrgInputEventDate from '../organisms/OrgInputEventDate'
  import MolTwoCellTable from '../molecules/MolTwoCellTable'
  import AtomButton from '../atoms/AtomButton'
  import OrgMoveConfirm from '../organisms/OrgMoveConfirm'
  import MolLoadingCard from '../molecules/MolLoadingCard'
  import AtomLoader from '../atoms/AtomLoader'
  export default {
    name: 'PageEventEdit',
    components: { AtomLoader, MolLoadingCard, OrgMoveConfirm, AtomButton, MolTwoCellTable, OrgInputEventDate, MolInputDateTime, MolInputCard, AtomInput, AtomCard, OrgInputEventHeader },
    props: {
      isCreatePage: {
        type: Boolean,
        default: false
      }
    },
    data () {
      return {
        event: {},
        openTimeModal: false,
        startTimeModal: false,
        errorMsgs: [],
        loading: {
          post: false
        },
        pollingId: 0,
        updated: false
      }
    },
    computed: {
      eventLoading () {
        return !this.isCreatePage && !this.event.id
      }
    },
    methods: {
      ...mapActions('events',[
        'postEvent',
        'putEvent'
      ]),
      openTimeModalMutation (v) {
        this.openTimeModal = v
      },
      startTimeModalMutation (v) {
        this.startTimeModal = v
      },
      async clickPost () {
        this.loading.post = true
        try {
          let id = ''
          if (this.isCreatePage) {
            id = await this.postEvent(this.event)
            localStorage.removeItem('createEvent')
          } else {
            await this.putEvent(this.event)
            this.updated = true
            id = this.event.id
          }
          this.$router.push(`/events/${id}`)
        } catch (msgs) {
          this.errorMsgs = msgs
        }
        this.loading.post = false
      },
      // localStorageへ編集内容を保存。ポーリングで行う。
      // enterHandler, leaveHandlerは親コンポーネントから呼ばれる
      // enterHandlerが呼ばれる前にeventが初期化されるので、
      // event自身をdeep watchしてしまうと初期化時に発火してしまい、
      // localStorageが初期状態に更新されてしまう
      enterHandler (initialEvent) {
        this.event = initialEvent
        this.updated = false

        if (this.isCreatePage) {
          this.pollingId = setInterval(() => {
            localStorage.setItem('createEvent', JSON.stringify(this.event))
          }, 1000)
        }
      },
      leaveHandler (next = () => {}) {
        if (this.isCreatePage) {
          localStorage.setItem('createEvent', JSON.stringify(this.event))
          clearInterval(this.pollingId)
        } else if (!this.updated) {
          this.$refs.confirm.leaveHandler(next)
        } else {
          next()
        }
      }
    }
  }
</script>

<style scoped>
  .name-input {
    width: 92%;
    font-size: 24px !important;
    font-weight: bold;
  }
  .meta-input {
    font-style: italic;
  }
</style>
