<template>
  <div id="app">
    <div class="main-container">
      <router-view>
      </router-view>
    </div>
    <a class="phone"  v-show="!this.$route.path.includes('console')" :href="contacts[0].href"><inline-svg class="phone-icon" :src="require('./assets/icons/phone-icon.svg')"/></a>
    <notifications group="test"/>
    <Footer v-show="!this.$route.path.includes('console')"/>
  </div>
</template>

<script>
import Footer from "@/components/Footer";
import phones from "@/assets/items/Contacts";

export default {
  name: 'App',
  components: {Footer},
  data: function () {
    return {
      contacts: phones,
    }
  },

}
</script>

<style>

#app {
  font-family: Raleway, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  scroll-behavior: smooth;
  /*color: #2c3e50;*/
}

.main-container {
  min-height: 100vh;
}

a.phone {
  display: flex;
  border: none;
  padding-left: 15px;
  padding-right: 15px;
  position: sticky;
  margin-right: 40px;
  float: right;
  bottom: 60px;
  height: 50px;
  background-color: #EE6723;
  font-size: 22px;
  color: #FFFFFF;
}

.phone-icon {
  margin: auto;
  width: 30px;
  fill: #ffffff;
}

.page-content-container {
  position: relative;
  /*display: flex;*/
  /*flex-direction: row;*/
  display: grid;
  /*grid-template-columns: [head-start nav-start] auto [nav-end main-start] auto [head-end main-end];*/
  /*grid-template-rows:  [head-start] auto [head-end nav-start main-start] auto [nav-end main-end];*/
  grid-template-areas: "head head"
                        "nav  main"
                        "other other";
  /*grid-template-columns: repeat(2, 1fr);*/
  /*flex-wrap: wrap;*/
  left: 50vw;
  transform: translate(-50%);
}

@media screen and (min-width: 320px) {
  .page-content-container {
    max-width: 320px;
    grid-template-areas:
        "head head"
        "nav  ."
        "main main"
        "about about"
        "contactUs contactUs"
  ;
  }
}

@media screen and (min-width: 480px) {
  .page-content-container {
    max-width: 480px;
    grid-template-areas:
        "head head"
        "nav  ."
        "main main"
        "about about"
        "contactUs contactUs";
  }
}

@media screen and (min-width: 768px) {
  .page-content-container {
    max-width: 768px;
    grid-template-areas: "head head"
                        "nav  ."
                        "main main"
                        "about about"
        "contactUs contactUs";
  }
}

@media screen and (min-width: 1000px) {
  .page-content-container {
    max-width: 1000px;
    grid-template-areas: "head head"
                        "nav  main"
                        "about about"
        "contactUs contactUs";
  }
}

@media screen and (min-width: 1440px) {
  .page-content-container {
    max-width: 1440px;
    grid-template-areas: "head head"
                        "nav  main"
                        "about about"
        "contactUs contactUs";
  }
}
</style>
