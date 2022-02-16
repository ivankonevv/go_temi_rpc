<template>
  <div class="content-container">
    <NavBar :nav="navBarLink"/>
    <div class="shipping-page-container">
      <div class="shipping-cards-container">
        <div class="shipping-card" v-for="(item, index) in upItems" :key="index" v-html="item.icon" />
      </div>

      <div class="shipping-info-container">
        <div class="shipping-info-section" v-for="(item, index) in items" :key="index">
          <div class="shipping-info-title" @click="item.opened=!item.opened">
            <span>{{ item.title }}</span>
            <inline-svg v-if="item.opened" :src="require('../assets/icons/shipping_arrow.svg')"/>
            <inline-svg v-else :src="require('../assets/icons/shipping_arrow.svg')" :style="{transform: 'rotate(-180deg)'}"/>
          </div>
          <div class="shipping-info-description" v-html="item.text" v-show="item.opened"></div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import shippingMenu from "@/assets/items/ShippingItems";
import shippingUpItems from "@/assets/items/ShippingUpItems";
export default {
  name: "Shipping",
  components: {NavBar},
  data: function () {
    return {
      navBarLink: "Доставка и оплата",
      upItems: shippingUpItems,
      items: shippingMenu
    }
  },
}
</script>

<style scoped lang="scss">
  @import "../assets/css/base/page.css";
  @import "../assets/css/base/shipping";
</style>