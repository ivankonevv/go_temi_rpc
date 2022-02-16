<template>
  <div class="main-console-container">
    <ConsoleHeader />
    <ConsoleSideBar />
    <AddNewProduct v-if="this.$route.path === '/console/add/wood'" :items="null"/>
    <AddNewProduct v-if="this.$route.path === '/console/add/metal'" :items="metal"/>
  </div>
</template>

<script>
import ConsoleHeader from "@/components/auth/ConsoleHeader";
import ConsoleSideBar from "@/components/auth/ConsoleSideBar";
import AddNewProduct from "@/components/auth/add/AddNewProduct";
import metalDoorCreationItems from "@/assets/items/MetalDoorsItems";
export default {
  name: "AddPage",
  components: { AddNewProduct, ConsoleSideBar, ConsoleHeader },
  data () {
    return {
      metal: metalDoorCreationItems
    }
  }
}
</script>

<style scoped>
  @import "../../../assets/css/auth/console.css";
</style>