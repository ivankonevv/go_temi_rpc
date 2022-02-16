<template>
  <div class="page-navbar">
    <router-link class="navbar-value" to="/">Главная</router-link>
    <span class="navbar-value"> / {{this.nav}}</span>
  </div>
</template>

<script>
export default {
  name: "NavBar",
  data: function () {
    return {
      // value: this.$parent.navBarLink
    }
  },
  props: ['nav']
}
</script>

<style scoped>

</style>