<template>
  <!--TODO コンポーネント切り出せそう-->
  <div>
    <atom-card>
      <atom-card-content slot="description">
        <nuxt-link :to="`/${invite.inviter.identifier}s/${invite.inviter.id}`">
          <atom-image
            :src="invite.inviter.icon_url"
            size="tiny"
            shape="circular"
          />
        </nuxt-link>
        <h3>
          <nuxt-link :to="`/${invite.inviter.identifier}s/${invite.inviter.id}`">
            {{ invite.inviter.name }}
          </nuxt-link>
          さんからのスカウトです！
        </h3>
      </atom-card-content>
    </atom-card>

    <mol-content-card
      title="スカウト詳細"
      :description="invite.description"
    />

    <mol-content-card
      title="スカウトされたイベント"
    >
      <nuxt-link
        slot="bottom"
        :to="`/events/${invite.event.id}`"
      >
        <mol-content-card
          :title="invite.event.name"
          :image-src="invite.event.header_image_url"
        />
      </nuxt-link>
    </mol-content-card>

    <atom-card>
      <atom-container
        v-if="isUnchecked"
        slot="description"
        style="text-align: center"
      >
        <atom-button-group>
          <atom-button
            basic
            positive
            @click="confirmApproveModal = true"
          >
            承認する
          </atom-button>
          <atom-button
            basic
            negative
            @click="confirmDeclineModal = true"
          >
            辞退する
          </atom-button>
        </atom-button-group>
      </atom-container>
      <div
        v-else
        slot="description"
        class="statusMessage"
      >
        <p v-show="isApproved">
          承認済み
        </p>
        <p v-show="isDeclined">
          辞退済み
        </p>
      </div>
    </atom-card>

    <org-confirm-invite-modal
      v-model="confirmApproveModal"
      title="スカウトを承認する"
      approve-button-text="承認する"
      decline-button-text="キャンセル"
      @click="approve"
    >
      <!--TODO 利用規約作成、リンク-->
      <p slot="description">
        利用規約に同意の上、承認してください。
      </p>
    </org-confirm-invite-modal>

    <org-confirm-invite-modal
      v-model="confirmDeclineModal"
      title="スカウトを辞退する"
      approve-button-text="はい"
      decline-button-text="いいえ"
      @click="decline"
    >
      <p slot="description">
        本当に辞退しますか？
      </p>
    </org-confirm-invite-modal>
  </div>
</template>

<script>
  import { mapActions, mapGetters } from 'vuex'
  import AtomCardContent from '../atoms/AtomCardContent'
  import AtomCard from '../atoms/AtomCard'
  import AtomImage from '../atoms/AtomImage'
  import MolContentCard from '../molecules/MolContentCard'
  import AtomButtonGroup from '../atoms/AtomButtonGroup'
  import AtomButton from '../atoms/AtomButton'
  import AtomContainer from '../atoms/AtomContainer'
  import OrgConfirmInviteModal from '../organisms/OrgConfirmInviteModal'
  export default {
    name: 'PageInvite',
    components: { OrgConfirmInviteModal, AtomContainer, AtomButton, AtomButtonGroup, MolContentCard, AtomImage, AtomCard, AtomCardContent },
    props: {
      invite: {
        type: Object,
        required: true
      }
    },
    data () {
      return {
        confirmApproveModal: false,
        confirmDeclineModal: false
      }
    },
    computed: {
      ...mapGetters('invites', [
        'unchecked',
        'approved',
        'declined'
      ]),
      isUnchecked () {
        return this.invite.status === this.unchecked
      },
      isApproved () {
        return this.invite.status === this.approved
      },
      isDeclined () {
        return this.invite.status === this.declined
      }
    },
    methods: {
      ...mapActions('invites', [
        'approveInvite',
        'declineInvite'
      ]),
      approve () {
        this.approveInvite(this.invite.id)
      },
      decline () {
        this.declineInvite(this.invite.id)
      }
    }
  }
</script>

<style scoped>
  .statusMessage {
    font-weight: bold;
    text-align: center;
  }
</style>
