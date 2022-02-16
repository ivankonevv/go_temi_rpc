<template>
  <div class="page-content-container">
    <Header />
    <SideBar />
    <Shipping />
  </div>
</template>

<script>
import SideBar from "@/components/SideBar";
import Shipping from "@/components/Shipping";
import Header from "@/components/Header";
export default {
  name: "ShippingAndPayment",
  components: {Header, Shipping, SideBar}
}
</script>

<style scoped>

</style>