<template>
  <div class="page-content-container">
    <Header/>
    <SideBar/>
    <BlogList />
  </div>
</template>

<script>
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import BlogList from "@/components/blog/BlogList";
export default {
  name: "Blog",
  components: {BlogList, SideBar, Header}
}
</script>

<style scoped>

</style>