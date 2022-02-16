<template>
  <div class="main-console-container">
    <ConsoleHeader/>
    <ConsoleSideBar/>
    <Dashboard />
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import ConsoleSideBar from "@/components/auth/ConsoleSideBar";
import ConsoleHeader from "@/components/auth/ConsoleHeader";
import Dashboard from "@/components/auth/dashboard/Dashboard";

export default {
  name: "DashboardPage",
  components: {Dashboard, ConsoleHeader, ConsoleSideBar},
  computed: {
    ...mapGetters("auth", {
      getUserProfile: "getUserProfile",
    }),
  }
}
</script>

<style scoped>
 @import "../../../assets/css/auth/console.css";
</style>