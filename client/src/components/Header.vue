<template>
  <header class="header-container">
    <div class="header-wrapper">
      <router-link class="header-logo-block" to="/">
        <inline-svg class="header-logo" :src="require(`../assets/icons/logo.svg`)"/>
      </router-link>
      <div class="header-search-container">
        <input class="header-search-input" type="text" placeholder="Поиск"/>
        <button class="header-search-button">
          <inline-svg :src="require(`../assets/icons/search_icon.svg`)"/>
        </button>
      </div>

      <inline-svg class="header-mobile-search" :src="require(`../assets/icons/mobile-search.svg`)"/>

      <div class="header-contacts-block">
        <a class="header-contact-number" v-for="(contact, index) in contacts.slice(0,2)"
           :key="index">{{ contact.short }}</a>
      </div>
      <inline-svg class="header-burger-menu" :src="require(`../assets/icons/burger.svg`)"/>
      <button class="header-contact-button" @click="scrollToElement">Заказать звонок</button>
    </div>
  </header>
</template>

<script>

import phones from "@/assets/items/Contacts";

export default {
  name: "Header",
  data: function () {
    return {
      contacts: phones
    }
  },
  methods: {
    async scrollToElement() {
      if (this.$route.path !== '/')  {
        await this.$router.push({path: '/'})
      }
      const el = document.getElementById("contactUs");
      if (el) {
        el.scrollIntoView({behavior: 'smooth'});
      }
    }
  }
}
</script>

<style scoped>
@import "../assets/css/base/header.css";
</style>