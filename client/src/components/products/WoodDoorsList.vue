<template>
  <div class="content-container">
    <NavBar :nav="navBarLink"/>
    <div class="product-container">
      <template v-if="error">
        <span style="font-size: 26px; margin-top: 40px;">Произошла какая-то ошибка :(</span>
      </template>
      <loading :active.sync="loading" :height="100" :width="100" :is-full-page="true" :can-cancel="true"
               :loader="'bars'" :color="'#EE6723'"/>
      <template v-if="!loading && result">
        <WoodDoorCard v-for="item in result" :key="item.id" :item="item"/>
      </template>
    </div>
  </div>
</template>

<script>
import {WoodDoorsApiClient} from "../../../proto/wood-doors_grpc_web_pb";
import WoodDoorCard from "@/components/products/WoodDoorCard";
import NavBar from "@/components/NavBar";
import Loading from 'vue-loading-overlay';


export default {
  name: "WoodDoorsList",
  components: {NavBar, WoodDoorCard, Loading},
  data () {
    return {
      navBarLink: "Межкомнатные двери",
      result: [],
    }
  },
  created() {
      this.client = new WoodDoorsApiClient(process.env.VUE_APP_SERVER_URL, null, null)
      console.log(this.client)
      const {WoodDoorsRequest} = require("../../../proto/wood-doors_pb");

      let request = new WoodDoorsRequest();
      let stream = this.client.getWoodDoors(request, {});
      stream.on('data', (response) => {
        this.result.push(response.toObject().data);
        console.log(this.result)
      })
      this.loading = false

  }
}
</script>

<style scoped>
@import "../../assets/css/base/product_list.css";
@import "../../assets/css/base/page.css";
</style>