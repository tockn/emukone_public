<template>
  <div class="tags">
    <span
      v-for="(tag, i) in tags"
      :key="i"
    >
      <atom-tag
        class="el"
        :name="tag.name"
        @click="openModal(i)"
      />
      <atom-modal
        v-model="modalFlags[i]"
        class="atom-modal"
      >
        <p slot="header">タグを編集</p>
        <div
          slot="description"
          class="description"
        >
          <atom-droppdown
            :value="tags[i].name"
            :options="allTags"
            search
            fluid
            selection
            @input="funcs[i]"
          />
        </div>
        <atom-button
          slot="action"
          @click="closeModal(i)"
        >
          OK
        </atom-button>
      </atom-modal>
    </span>
  </div>
</template>

<script>
  import { mapState, mapMutations, mapActions } from 'vuex'
  import AtomTag from '../atoms/AtomTag'
  import AtomModal from '../atoms/AtomModal'
  import AtomButton from '../atoms/AtomButton'
  import AtomDroppdown from '../atoms/AtomDroppdown'
  export default {
    name: 'OrgInputTags',
    components: { AtomDroppdown, AtomModal, AtomTag, AtomButton},
    data () {
      return {
        modalFlags: [],
        funcs: []
      }
    },
    computed: {
      ...mapState({
        tags: state => state.users.user.tags,
        stateAllTags: state => state.users.tags
      }),
      allTags () {
        let ts = []
        for (let i = 0; i < this.stateAllTags.length; i++) {
          let t = this.stateAllTags[i]
          ts[i] = {key: t.name, text: t.name, value: t.name}
        }
        return ts
      }
    },
    created () {
      this.tags.forEach((tag, i) => {
        this.modalFlags[i] = false
        this.funcs[i] = (v) => {
          this.setTagName(i, v)
        }
      }, this)
      this.getAllTags()
    },
    methods: {
      ...mapMutations('users', [
        'setTags'
      ]),
      ...mapActions('users', [
        'getAllTags'
      ]),
      openModal (i) {
        this.modalFlags[i] = true
        this.modalFlags = JSON.parse(JSON.stringify(this.modalFlags))
      },
      closeModal (i) {
        this.modalFlags[i] = false
        this.modalFlags = JSON.parse(JSON.stringify(this.modalFlags))
      },
      setTagName(index, name) {
        let tmp = JSON.parse(JSON.stringify(this.tags))
        tmp[index].name = name
        this.setTags(tmp)
      }
    }
  }
</script>

<style scoped>
  .el {
    margin: 8px auto;
  }
  .tags {
    margin: 8px 0 16px 0;
  }
  .description {
    text-align: center;
  }
</style>
