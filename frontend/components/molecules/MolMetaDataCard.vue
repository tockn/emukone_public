<template>
  <div class="wrap">
    <atom-icon :src="user.icon_url" />
    <p class="name">
      {{ user.name }}
    </p>
    <p class="description">
      "{{ user.meta_description }}"
    </p>
    <div class="tags">
      <atom-tag
        v-for="(tag, i) in user.tags"
        :key="i"
        :name="tag.name"
      />
    </div>
    <slot />
  </div>
</template>

<script>
  import AtomTag from '../atoms/AtomTag'
  import AtomIcon from '../atoms/AtomIcon'
  export default {
    name: 'MolMetaDataCard',
    components: { AtomIcon, AtomTag},
    props: {
      user: {
        type: Object,
        required: true
      }
    }
  }
</script>

<style scoped>
  .wrap {
    -webkit-border-radius: 8px;
    -moz-border-radius: 8px;
    border-radius: 8px;
    text-align: center;
    padding: 16px 24px;
    color: white;
    width: 100%;
  }
  .name {
    margin: 8px 0;
    font-size: 28px;
  }
  .description {
    font-style: italic;
  }
  .tags {
    margin: 8px auto 24px auto;
  }
</style>
