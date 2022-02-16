<template>
  <div class="sidebar-container">
    <div class="sidebar-wrapper">

    <div class="sidebar-dropdown-container">
      <span class="sidebar-dropdown-title" v-on:click="isHidden = !isHidden">
        Каталог
        <inline-svg :src="require(`../assets/icons/sidebar_arrow.svg`)" v-if="isHidden"/>
        <inline-svg :src="require(`../assets/icons/sidebar_arrow.svg`)" v-else style="transform: rotate(90deg)"/>
      </span>
      <div class="sidebar-dropdown-menu" v-if="!isHidden">
<!--        <div class="sidebar-dropdown-parent" v-for="(link, index) in catalogLinks" :key="index">-->
<!--          <router-link class="sidebar-dropdown-item" v-show="!link.child" :to="link.to">{{link.link_title}}</router-link>-->
<!--          <span class="sidebar-item" v-show="link.child">{{link.link_title}}</span>-->
<!--          <router-link class="sidebar-dropdown-item" v-for="(child, i) in link.child" :key="i" :to="child.to">{{child.link_title}}</router-link>-->
<!--        </div>-->
        <div class="sidebar-dropdown-parent" v-for="(link, index) in catalogLinks" :key="index">
          <span class="sidebar-child" v-if="link.child" v-on:click="isHiddenChild = !isHiddenChild">
            {{ link.link_title }}
            <inline-svg :src="require(`../assets/icons/sidebar_arrow.svg`)" v-show="link.child" v-if="isHiddenChild"/>
            <inline-svg :src="require(`../assets/icons/sidebar_arrow.svg`)" v-show="link.child" v-else style="transform: rotate(90deg)"/>
          </span>
          <router-link class="sidebar-child" v-show="!link.child" :to="link.to">{{ link.link_title }}</router-link>
          <div class="sidebar-child-container" v-for="(child, i) in link.child" :key="i" v-show="link.child && !isHiddenChild">
            <router-link class="sidebar-child" :to="child.to">{{ child.link_title }}</router-link>
          </div>
        </div>
<!--        <router-link class="sidebar-dropdown-item" v-for="(link, index) in catalogLinks" :key="index" :to="link.to">{{link.link_title}}</router-link>-->
      </div>
      <a class="sidebar-item" @click="scrollToElement">
        О компании
      </a>
      <router-link class="sidebar-item" v-for="(link, index) in navLinks" :to="link.to" :key="index">{{link.link_title}}</router-link>

    </div>

    </div>

  </div>
</template>

<script>
import catalogLinks from "@/assets/items/CatalogLinks";
import navLinks from "@/assets/items/NavLinks";

export default {
  name: "SideBar",
  data: function () {
    return {
      catalogLinks: catalogLinks,
      navLinks: navLinks,
      isHidden: false,
      isHiddenChild: true,
    }
  },
  methods: {
    async scrollToElement() {
      if (this.$route.path !== '/')  {
        await this.$router.push({path: '/'})
      }
      const el = document.getElementById("aboutUs");
      if (el) {
        el.scrollIntoView({behavior: 'smooth'});
      }
    }
  }
}
</script>

<style scoped>
 @import "../assets/css/base/sidebar.css";
</style>