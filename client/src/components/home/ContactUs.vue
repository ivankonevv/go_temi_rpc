<template>
  <div class="contact-us-block" id="contactUs">
    <span class="contact-us-block-title">СВЯЗАТЬСЯ С НАМИ</span>
    <div class="contact-us-grid-container">
      <div class="contact-us-image" v-bind:style="{backgroundImage: `url(${images.home.contactUs[0]})`}"/>
      <div class="contact-us-form">
        <span class="contact-us-form-title">Заказать звонок</span>
        <span class="contact-us-form-text">Оставьте заявку и мы свяжемся с Вами</span>
        <form class="contact-us-form" @submit.prevent @submit="submitForm">
          <input class="contact-us-input" type="text" v-model="name" required placeholder="Ваше имя"/>
          <div class="contact-us-input-underline"/>
          <input class="contact-us-input" type="tel" v-model="phone" required placeholder="Номер телефона"/>
          <div class="contact-us-input-underline"/>
          <button class="contact-us-button" type="submit">Отправить</button>
        </form>
        <div class="contact-us-line"/>
      </div>
    </div>
  </div>
</template>

<script>
import imageLinks from "@/assets/items/ImageLinks";
import {CallbackServiceClient} from "../../../proto/callback_service_grpc_web_pb";
import contactUsNotifications from "@/components/home/contactUsNotifications";

export default {
  name: "ContactUs",
  data () {
    return {
      images: imageLinks,
      phone: null,
      name: '',
    }
  },
  methods: {
    submitForm () {
      if (this.name.length > 1 && this.phone.length > 7) {
        this.client = new CallbackServiceClient(process.env.VUE_APP_SERVER_URL, null, null)
        const { CallbackRequest } = require("../../../proto/callback_service_pb")
        let request = new CallbackRequest
        request.setName(this.name)
        request.setPhone(this.phone)
        this.client.callback(request, {}, (err) => {
          if (err) {
            this.$notify(contactUsNotifications.error.internal)
          } else {
            this.$notify(contactUsNotifications.success)
          }
        })
      } else {
        this.$notify(contactUsNotifications.error.incorrect_input)
      }

    }
  }
}
</script>

<style scoped lang="scss">
  @import "../../assets/css/base/contact_us_block";
</style>