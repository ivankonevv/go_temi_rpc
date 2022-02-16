<template>
  <div class="page-content-container">
    <Header />
    <SideBar />
    <NotFoundCom/>
  </div>
</template>

<script>
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import NotFoundCom from "@/components/NotFoundCom";
export default {
  name: "NotFound",
  components: {NotFoundCom, SideBar, Header}
}
</script>

<style scoped>

</style>