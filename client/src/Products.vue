<template>
  <div class="page-content-container">
    <Header />
    <SideBar />
    <ProductList v-if="this.$route.path.includes('/metal-doors')"/>
    <WoodDoorsList v-else/>
  </div>
</template>

<script>
import ProductList from "@/components/products/ProductList";
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import WoodDoorsList from "@/components/products/WoodDoorsList";
export default {
  name: "Products",
  components: {WoodDoorsList, SideBar, Header, ProductList}
}
</script>

<style scoped>

</style>