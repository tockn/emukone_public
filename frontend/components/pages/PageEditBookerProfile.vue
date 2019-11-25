<template>
  <div>
    <org-input-profile-header>
      <atom-button
        :loading="editUserLoading"
        color="green"
        @click="edit"
      >
        保存する
      </atom-button>
    </org-input-profile-header>
    <mol-input-card
      title="活動に懸ける想い"
      :value="user.why_description"
      @input="setWhyDescription"
    />
    <mol-input-card
      title="こんなことやっています"
      :value="user.how_description"
      @input="setHowDescription"
    />
    <mol-input-card
      title="その他の情報"
      :value="user.free_description"
      @input="setFreeDescription"
    />
  </div>
</template>

<script>
  import { mapState, mapMutations, mapActions } from 'vuex'
  import MolInputCard from '../molecules/MolInputCard'
  import OrgInputProfileHeader from '../organisms/OrgInputProfileHeader'
  import AtomButton from '../atoms/AtomButton'
  export default {
    name: 'PageEditBookerProfile',
    components: { OrgInputProfileHeader, MolInputCard, AtomButton },
    computed: {
      ...mapState('users', {
        user: state => state.user,
        editUserLoading: state => state.loading.editUser
      })
    },
    methods: {
      ...mapMutations('users', [
        'setWhyDescription',
        'setFreeDescription',
        'setHowDescription'
      ]),
      ...mapActions('users', [
        'editUser'
      ]),
      async edit () {
        await this.editUser()
        this.$router.push(`/bookers/${this.user.id}`)
      }
    }
  }
</script>

<style scoped>

</style>
