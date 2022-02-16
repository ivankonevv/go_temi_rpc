<template>
  <div class="page-content-container">
    <Header />
    <SideBar />
    <SingleProduct />
  </div>
</template>

<script>
import SingleProduct from "@/components/product/SingleProduct";
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
export default {
  name: "Product",
  components: {SideBar, Header, SingleProduct},
}
</script>

<style scoped>

</style>