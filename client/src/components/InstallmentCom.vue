<template>
  <div class="content-container">
    <NavBar :nav="navBarLink"/>
    <div class="installment-main-container">
      <div class="installment-upper-block">
        <div class="installment-upper-text-container">
          <div class="installment-upper-text">
            Входные двери, межкомнатные двери, а так же окна и двери ПВХ вы можете приобрести в РАССРОЧКУ.
            <div class="installment-upper-text-down">
              Мы предоставляем банковскую рассрочку от ОАО "Белагропромбанк"
              сроком<br/> от <b>1</b> до <b>36</b> месяцев.
              Рассрочка "Легкие покупки" сроком до <b>4</b> месяцев предоставляется бесплатно.
            </div>
          </div>
        </div>
        <img src="../assets/images/bank.png" alt="bank"/>
      </div>
      <div class="installment-lower-block">
        <div class="installment-lower-container">
          <div class="installment-lower-box" v-for="(item, index) in items" :key="index">
            <span class="installment-lower-box-number">{{ index+1 }}</span>
            <span class="installment-lower-box-text" v-html="item.text"></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import installmentItems from "@/assets/items/InstallmentItems";

export default {
  name: "InstallmentCom",
  components: {NavBar},
  data() {
    return {
      items: installmentItems,
      navBarLink: "Рассрочка",
    }
  }
}
</script>

<style scoped lang="scss">
  @import "../assets/css/base/page.css";
  @import "../assets/css/base/installment";
</style>