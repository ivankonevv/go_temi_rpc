<template>
  <div class="content-container">
    <NavBar :nav="navBarLink"/>
    <div class="w-colors-warn-block">
      <div class="w-color-warn-line"/>
      <span class="w-color-warn-text">Просим обратить внимание, что фотографии цветов могут отличаться от реального восприятия.<br/>Если вы хотите посмотреть на цвет вживую - вы можете подойти в любой из наших
        <router-link to="/contacts">салонов</router-link>.</span>
    </div>
    <div class="w-colors-block" v-for="manufacturer in result" :key="'a' + manufacturer.id" v-show="manufacturer.collectionsList.length > 1">
      <span class="w-colors-b-title">{{ manufacturer.manufacturer }}</span>
      <div class="w-colors-b-content" v-for="collection in manufacturer.collectionsList" :key="'b' + collection.id" v-show="collection.colorsList.length > 0">
        <span class="w-colors-b-collection-title">Коллекция <b>{{ collection.title }}</b></span>
        <div class="w-color-collection">
          <div class="w-colors-b-c-box" v-for="color in collection.colorsList" :key="'c' + color.uuid">
            <div class="w-colors-b-c-b-image" v-bind:style="{backgroundImage: `url(${color.image})`}"/>
            <span class="w-colors-b-c-d-title">{{color.title}}</span>
            <span class="w-colors-b-c-b-unique" v-show="color.ral">RAL: <b>{{ color.ral }}</b></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import {ColorsApiClient} from "../../../proto/wood-colors_grpc_web_pb";
export default {
  name: "Colors",
  components: {NavBar},
  data () {
    return {
      navBarLink: "Выкраски",
      result: [],
    }
  },
  created() {
    this.client = new ColorsApiClient(process.env.VUE_APP_SERVER_URL, null, null)
    const {ManufacturersWColorsRequest} = require("../../../proto/wood-colors_pb");
    let request = new ManufacturersWColorsRequest();
    let stream = this.client.getMwColors(request, {});
    stream.on('data', (response) => {
      this.result.push(response.toObject());
    })
    this.loading = false
  }
}
</script>

<style scoped lang="scss">
 @import "../../assets/css/base/wood-colors";
</style>