<template>
  <div class="about-block" id="aboutUs">
    <span class="about-block-title">О КОМПАНИИ</span>
    <div class="about-grid-container">

      <div class="about-small-box" style="grid-area: leftup">
        <span class="about-box-title">ОДО «ТЭМИ-пласт»</span>
        <div class="about-line"/>
        <span class="about-box-text">
          Занимаемся оптовой и розничной продажей металлических и межкомнатных дверей, а также окон и дверей ПВХ с 2009 года.
        </span>
      </div>
      <div class="about-small-box" v-bind:style="{gridArea: 'rightup', backgroundImage: `url(${require('@/assets/images/test-image.png')})`}"/>

      <div class="about-small-box" style="grid-area: rightdown">
        <span class="about-box-text">
          Мы дорожим каждым клиентом, поэтому активно изучаем спрос и имеющиеся предложения на рынке, стремясь постоянно обновлять ассортимент и предлагать нашим покупателям только самое лучшее.
        </span>
      </div>
      <div class="about-small-box" v-bind:style="{gridArea: 'leftdown', backgroundImage: `url(${require('@/assets/images/test-image.png')})`}"/>
      <div class="about-big-box" v-bind:style="{gridArea: 'center', backgroundImage: `url(${images.home.about[0]})`}"/>

    </div>
  </div>
</template>

<script>
import imageLinks from "@/assets/items/ImageLinks";

export default {
  name: "About",
  data () {
    return {
      images: imageLinks
    }
  }
}
</script>

<style scoped lang="scss">
  @import '../../assets/css/base/about-block';
</style>