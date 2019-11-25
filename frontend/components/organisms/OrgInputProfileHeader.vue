<template>
  <div>
    <div class="wrap">
      <mol-input-image
        :value="user.icon_url"
        :modal-flag="imageLoading"
        @input="updateIconUrl"
      />
      <atom-input
        class="name"
        placeholder="名前"
        :white="true"
        :value="user.name"
        @input="setName"
      />
      <atom-input
        class="description"
        placeholder="ひとこと"
        :white="true"
        :value="user.meta_description"
        @input="setMetaDescription"
      />

      <org-input-tags />

      <slot />
    </div>
    <atom-header-image
      class="header-img"
      :src="user.icon_url"
    />
  </div>
</template>

<script>
  import { mapState, mapMutations, mapActions } from 'vuex'
  import AtomHeaderImage from '../atoms/AtomHeaderImage'
  import AtomInput from '../atoms/AtomInput'
  import OrgInputTags from './OrgInputTags'
  import MolInputImage from '../molecules/MolInputImage'
  export default {
    name: 'OrgInputProfileHeader',
    components: { MolInputImage, OrgInputTags, AtomInput, AtomHeaderImage },
    data () {
      return {
        iconImageFile: {},
        imageLoaded: false,
        imagePreview: ""
      }
    },
    computed: {
      ...mapState('users', {
        user: state => state.user,
        imageLoading: state => state.loading.iconImage
      })
    },
    methods: {
      ...mapMutations('users', [
        'setName',
        'setMetaDescription'
      ]),
      ...mapActions('users', [
        'updateIconUrl'
      ])
    }
  }
</script>

<style scoped>
  .wrap {
    -webkit-border-radius: 8px;
    -moz-border-radius: 8px;
    border-radius: 8px;
    text-align: center;
    padding: 32px 24px;
    color: white;
    width: 100%;
    height: 360px;
  }
  .name {
    font-size: 28px;
    padding: 2px;
    display: block;
    margin: auto;
  }
  .description {
    font-style: italic;
    margin-top: 2px !important;
    margin-bottom: 2px !important;
  }
</style>

