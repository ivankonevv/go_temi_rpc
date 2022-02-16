<template>
  <div class="page-content-container">
    <Header/>
    <SideBar/>
    <ContactsAndAddresses />
    <iframe class="contacts-map" src="https://yandex.ru/map-widget/v1/?um=constructor%3A5bf731b2730e67bed03151ca009630afbbe8f4a9875454e13db102ad385394b4&amp;source=constructor"
            frameBorder="0"/>
  </div>
</template>

<script>
import Header from "@/components/Header";
import SideBar from "@/components/SideBar";
import ContactsAndAddresses from "@/components/ContactsAndAddresses";
export default {
  name: "Contacts",
  components: {ContactsAndAddresses, SideBar, Header}
}
</script>

<style scoped lang="scss">
  @import "./assets/css/base/contacts.scss";
</style>