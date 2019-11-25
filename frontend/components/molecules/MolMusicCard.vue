<template>
  <div>
    <AtomCard>
      <template slot="header">
        イチオシの曲
      </template>
      <template slot="description">
        <sui-embed
          v-if="hasEmbed"
          :id="ichioshi.embed"
          :source="ichioshi.service"
          :active="true"
          :autoplay="false"
        />
      </template>
    </AtomCard>
  </div>
</template>

<script>
  import AtomCard from '../atoms/AtomCard'
  import SuiEmbed from 'semantic-ui-vue/dist/commonjs/modules/Embed/Embed'
  export default {
    name: 'MolMusicCard',
    components: { SuiEmbed, AtomCard },
    props: {
      ichioshi: {
        type: Object,
        required: true
      }
    },
    computed: {
      hasEmbed () {
        // eslint-disable-next-line
        return this.ichioshi.service === 'youtube'
      }
    }
  }
</script>

<style scoped>
  pre {
    white-space: pre-wrap;
  }
  .widget {
    text-align: center;
  }
</style>
