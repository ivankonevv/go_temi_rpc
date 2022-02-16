<template>
  <div class="content-container">
    <NavBar :nav="navBarLink"/>
    <div class="blog-container">
      <BlogCard v-for="(item, index) in items" :key="index" :item="item"/>
    </div>
    <b-pagination-nav
        v-model="currenPage"
        :number-of-pages="pages"
        base-url="#"
        first-number
        last-number
    />
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import BlogCard from "@/components/blog/BlogCard";
import blogTestItems from "@/assets/items/BlogTestItems";
export default {
  name: "BlogList",
  components: {BlogCard, NavBar},
  data: function () {
    return {
      navBarLink: "Статьи",
      items: blogTestItems,
      pages: 20,
      currenPage: 4,
    }
  },
}
</script>

<style scoped>
  @import "../../assets/css/base/page.css";
  @import "../../assets/css/base/blog_list.css";
</style>