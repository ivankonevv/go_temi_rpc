<template>
  <div class="page-content-container">
    <Header />
    <SideBar />
    <Colors />
  </div>
</template>

<script>
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import Colors from "@/components/wood/Colors";
export default {
  name: "WoodColors",
  components: {Colors, SideBar, Header}
}
</script>

<style scoped>

</style>