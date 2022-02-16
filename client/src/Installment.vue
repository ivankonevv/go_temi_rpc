<template>
  <div class="page-content-container">
    <Header />
    <SideBar />
    <InstallmentCom />
  </div>
</template>

<script>
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import InstallmentCom from "@/components/InstallmentCom";
export default {
  name: "Installment",
  components: {InstallmentCom, SideBar, Header}
}
</script>

<style scoped>

</style>