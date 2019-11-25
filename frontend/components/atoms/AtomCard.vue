<template>
  <sui-card
    class="centered card custom-width wrap"
  >
    <slot name="top" />
    <sui-card-content>
      <sui-card-header>
        <div class="header">
          <slot
            class="header"
            name="header"
          />
        </div>
      </sui-card-header>
      <sui-card-description v-show="!noDescription">
        <slot name="description" />
      </sui-card-description>
    </sui-card-content>
    <slot name="bottom" />
  </sui-card>
</template>

<script>
  import SuiCard from 'semantic-ui-vue/dist/commonjs/views/Card/Card'
  import SuiCardContent from 'semantic-ui-vue/dist/commonjs/views/Card/CardContent'
  import SuiCardHeader from 'semantic-ui-vue/dist/commonjs/views/Card/CardHeader'
  import SuiCardDescription from 'semantic-ui-vue/dist/commonjs/views/Card/CardDescription'
  export default {
    name: 'AtomCard',
    components: { SuiCardDescription, SuiCardHeader, SuiCardContent, SuiCard },
    props: {
      noDescription: {
        type: Boolean,
        default: false
      }
    }
  }
</script>

<style scoped>
  .header {
    font-size: 16px;
  }
  .wrap {
    margin-top: 10px !important;
    margin-bottom: 10px !important;
  }
</style>
