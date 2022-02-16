<template>
  <div class="blog-card-container">
    <span class="blog-card-title">{{ this.item.title }}</span>
    <div class="blog-card-content-container">
      <div class="blog-text-container">
        <div class="blog-text-block" v-for="(block, index) in this.item.text" :key="index">{{ block }}</div>
      </div>
      <div class="blog-primary-image" v-bind:style="{ backgroundImage: `url(${this.item.images.primary})` }"/>
    </div>
    <button class="blog-read-more-button">Читать далее...</button>
  </div>
</template>

<script>
export default {
  name: "BlogCard",
  props: ['item'],
  mounted() {

  }
}
</script>

<style scoped>
  @import "../../assets/css/base/blog_card.css";
</style>