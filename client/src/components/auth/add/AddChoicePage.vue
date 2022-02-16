<template>
  <div class="main-console-container">
    <ConsoleHeader />
    <ConsoleSideBar />
    <AddChoices />
  </div>
</template>

<script>
import ConsoleHeader from "@/components/auth/ConsoleHeader";
import ConsoleSideBar from "@/components/auth/ConsoleSideBar";
import AddChoices from "@/components/auth/add/AddChoices";
export default {
  name: "AddChoicePage",
  components: {AddChoices, ConsoleSideBar, ConsoleHeader}
}
</script>

<style scoped>
  @import "../../../assets/css/auth/console.css";
</style>