<template>
  <div>
    <atom-card>
      <template slot="header">
        イチオシの曲
      </template>
      <template slot="description">
        <atom-input
          class="input"
          :value="url"
          @input="updateUrl"
        />
        <sui-embed
          v-if="hasEmbed"
          :id="embed"
          :source="service"
          :active="true"
          :autoplay="false"
        />
      </template>
    </atom-card>
  </div>
</template>

<script>
  import { mapState, mapMutations } from 'vuex'
  import AtomCard from '../atoms/AtomCard'
  import AtomInput from '../atoms/AtomInput'

  const youtube1 = new RegExp(/(http|https):\/\/www\.youtube\.com\/watch\?v=.*/)
  const youtube2 = new RegExp(/(http|https):\/\/youtu\.be\//)

  export default {
    name: 'OrgInputMusicCard',
    components: { AtomInput, AtomCard },
    data () {
      return {
        title: '',
        service: '',
        url: '',
        embed: ''
      }
    },
    computed: {
      ...mapState('users', {
        ichioshis: state => state.user.ichioshis
      }),
      hasEmbed () {
        return youtube1.exec(this.url) || youtube2.exec(this.url)
      }
    },
    created () {
      let ichioshi = this.ichioshis[0]
      if (!ichioshi) return
      this.title = ichioshi.title
      this.service = ichioshi.service
      this.url = ichioshi.url
      this.embed = ichioshi.embed
    },
    methods: {
      ...mapMutations('users', [
        'setIchioshi'
      ]),
      updateUrl (v) {
        this.url = v
        if (youtube1.exec(this.url)) {
          this.embed = this.url.split('?v=')[1].split('&')[0]
          this.service = 'youtube'
        } else if (youtube2.exec(this.url)) {
          this.embed = this.url.split('/')[3]
          this.service = 'youtube'
        } else {
          this.embed = ''
          this.service = ''
        }
        let ichioshi = {
          service: this.service,
          url: this.url,
          embed: this.embed
        }
        this.setIchioshi(ichioshi)
      }
    }
  }
</script>

<style scoped>
  .input {
    width: 100%;
  }
  .widget {
    text-align: center;
  }
</style>
