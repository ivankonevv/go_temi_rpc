<template>
  <div class="content-container">
    <span class="not-found-text">Oops... Кажется, вы обратились не по адресу. Такой страницы не существует.</span>
  </div>
</template>

<script>
export default {
  name: "NotFoundCom"
}
</script>

<style scoped>
  @import "../assets/css/base/page.css";
  span.not-found-text {
    margin-left: 30px;
    margin-top: 30px;
    font-size: 30px;
  }
</style>