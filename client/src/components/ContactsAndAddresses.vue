<template>
  <div class="content-container">
    <NavBar :nav="navBarLink"/>
    
    <span class="contacts-block-title">Контактная информация</span>
    <div class="contacts-container">
      <div class="contacts-box">
        <span class="contacts-box-title">Адреса:</span>
        <div class="contacts-box-content">
<!--          {addresses.map((item) => (-->
<!--          <>-->
<!--          {item.full}-->
<!--          <br/><br/>-->
<!--        </>-->
<!--        ))}-->
          <span v-for="(item, index) in addresses" :key="index">
            {{ item.full }}
            <br/><br/>
          </span>
        </div>

      </div>
      <div class="contacts-box">
        <span class="contacts-box-title">Телефоны:</span>
        <div class="contacts-box-content">
          <div class="contacts-phones-container">
            <a class="contacts-phone" v-for="(item, index) in phones" :href="item.href" :key="index">{{ item.full }}</a>
          </div>
        </div>

        <span class="contacts-box-title" :style="{marginTop: '0'}">Email:</span>
        <div class="contacts-box-content">
          <a href="mailto:temi-plast@mail.ru">temi-plast@mail.ru</a>
        </div>
      </div>
      <div class="contacts-box">
        <span class="contacts-box-title">Юридическая информация:</span>
        <div class="contacts-box-content">
          Общество с дополнительной ответственностью "ТЭМИ-пласт"
          <br/><br/>
          290499234
          <br/><br/>
          р/с BY84 AKBB 3012 2091 5000 2600 0000
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import blogTestItems from "@/assets/items/BlogTestItems";
import phones from "@/assets/items/Contacts";
import addresses from "@/assets/items/Addresses";
export default {
  name: "ContactsAndAddresses",
  components: {NavBar},
  data: function () {
    return {
      navBarLink: "Контакты",
      items: blogTestItems,
      phones: phones,
      addresses: addresses,
    }
  },
}
</script>

<style scoped lang="scss">
  @import "../assets/css/base/page.css";
  @import "../assets/css/base/contacts.scss";
</style>