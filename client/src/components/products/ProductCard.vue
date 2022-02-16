<template>
  <div class="product-card-container" v-on:mouseover="buttonVisible=true" v-on:mouseleave="buttonVisible=false">
    <span class="product-card-title">{{ this.item.title }}</span>
    <img class="product-card-image" alt="product" v-if="!buttonVisible" :src="this.item.outImagesList[0]"/>
    <img class="product-card-image" alt="product" v-else :src="this.item.inImagesList[0]"/>
    <div class="product-card-price-container">
      <span class="after-price">{{ this.item.price }} руб</span>
      <span class="before-price">{{ this.item.price }} руб</span>
    </div>
    <div class="product-card-tags-block">
      <span class="product-card-tag" v-for="(tag, index) in this.item.tagsList" :key="'a' + index">{{ tag }}</span>
    </div>
    <div class="product-card-colors-block">
      <div class="product-card-color" v-for="(color, index) in this.item.inColorsList" :key="index" :title="color.color" v-bind:style="{backgroundColor: color.hex}"/>
    </div>

    <router-link :to="{ name: 'product', params: { id: this.item.id}}" v-if="mobile">
      <button class="product-card-detail-btn"><inline-svg :src="require(`@/assets/icons/product_card_arrow.svg`)"/></button>
    </router-link>
    <router-link :to="{ name: 'product', params: { id: this.item.id}}" v-else>
      <button class="product-card-detail-btn" v-if="buttonVisible">Подробнее</button>
    </router-link>

  </div>
</template>

<script>
export default {
  name: "ProductCard",
  data: function () {
    return {
      buttonVisible: false,
      mobile: window.innerWidth < 1440,
    }
  },
  props: ['item']
}
</script>

<style scoped>
  @import "../../assets/css/base/product_card.css";
</style>