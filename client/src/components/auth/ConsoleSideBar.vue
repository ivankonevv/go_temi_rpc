<template>
  <div class="con-sidebar-main-container" v-bind:class="{active: opened}">
    <inline-svg class="con-sidebar-icon" v-on:click="opened=!opened" :src="require('../../assets/icons/console_sidebar_logo.svg')"/>
    <div class="con-sidebar-nav-container" v-bind:class="{active: opened}">
      <router-link class="con-sidebar-link" to="/console/dashboard" active-class="con-sidebar-link-active">
        <div class="con-sidebar-link-icon-block" >
          <inline-svg class="con-sidebar-link-icon" :src="require('../../assets/icons/console_sidebar_home.svg')"/>
        </div>
        <span class="con-sidebar-link-title" v-bind:class="{active: opened}">Консоль</span>
      </router-link>
      <router-link class="con-sidebar-link" to="/console/storage" active-class="con-sidebar-link-active">
        <div class="con-sidebar-link-icon-block" >
          <inline-svg class="con-sidebar-link-icon" :src="require('../../assets/icons/console_sidebar_storage.svg')"/>
        </div>
        <span class="con-sidebar-link-title" v-bind:class="{active: opened}">База</span>
      </router-link>
      <router-link class="con-sidebar-link" to="/console/add" active-class="con-sidebar-link-active">
        <div class="con-sidebar-link-icon-block" >
          <inline-svg class="con-sidebar-link-icon" :src="require('../../assets/icons/console_sidebar_add.svg')"/>
        </div>
        <span class="con-sidebar-link-title" v-bind:class="{active: opened}">Добавить</span>
      </router-link>
    </div>
  </div>
</template>

<script>

export default {
  name: "ConsoleSideBar",
  data: function () {
    return {
      opened: true,
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/auth/console_sidebar.css";
</style>