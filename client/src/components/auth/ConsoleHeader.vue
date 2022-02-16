<template>
  <header class="console-header-container">
    <div class="console-header-user-block">
      <div class="console-header-user-avatar" @click="menuVisible=!menuVisible">
        <inline-svg class="console-header-user-icon" :src="require('../../assets/icons/user-svgrepo-com.svg')"/>
      </div>
      <div class="console-header-user-name-block not-selectable-text">
        <span class="console-header-user-name">
          {{getUserProfile.firstName}} {{getUserProfile.lastName}}
        </span>
        <span class="console-header-user-email">
          {{getUserProfile.email}}
        </span>
      </div>
      <div class="console-header-menu-block not-selectable-text" v-show="menuVisible">
        <div class="console-header-menu-row">
          <a class="console-header-menu-link">Еще</a>
        </div>
        <div class="console-header-menu-underline"/>
        <div class="console-header-menu-row" @click="logout()">
          <a class="console-header-menu-link">Выйти</a>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import {mapActions, mapGetters, mapMutations} from "vuex";

export default {
  name: "ConsoleHeader",
  data () {
    return {
      menuVisible: false
    }
  },
  computed: {
    ...mapGetters("auth", {
      getUserProfile: "getUserProfile",
      getLogout: "getLogout"
    }),
  },
  methods: {
    ...mapActions("auth", {
      userLogout: "userLogout",
    }),
    ...mapMutations("auth", {
      setLogout: "setLogout",
      setUserProfile: "setUserProfile"
    }),
    async logout() {
      await this.userLogout();
      if (this.getLogout) {
        const resetUser = {
          id: "",
          lastName: "",
          firstName: "",
          email: ""
        };
        this.setUserProfile(resetUser);
        console.log("ksajhdlgj")
        this.setLogout(false);
        console.log('asdasd')
        this.$router.push("/console/login");
      }
    }
  }
}
</script>

<style scoped>
  @import "../../assets/css/auth/console_header.css";
</style>