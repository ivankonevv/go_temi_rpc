<template>
  <div class="page-content-container">
    <Header/>
    <SideBar/>
    <About/>
    <ContactUs />
  </div>
</template>

<script>
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import About from "@/components/home/About";
import ContactUs from "@/components/home/ContactUs";
export default {
  name: "Home",
  // components: {Header}
  components: {ContactUs, About, SideBar, Header},

}
</script>

<style scoped>

</style>