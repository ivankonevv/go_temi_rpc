<template>
  <div class="content-container">
    <NavBar :nav="navBarLink" />
    <template v-if="error">
      <span style="font-size: 26px; margin-top: 40px;">Произошла какая-то ошибка :(</span>
    </template>
    <loading :active.sync="loading" :height="100" :width="100" :is-full-page="true" :can-cancel="true" :loader="'bars'" :color="'#EE6723'"/>
    <template v-if="item">
      <div class="product-info-container">
        <div class="main-image-container" v-for="(image, index) in item.variantsList[selectedColor].inImagesList" :key="'a' + index">
          <img v-if="image === selectedImage" class="product-main-image visible-image" :src="image" alt="main"/>
          <img v-else class="product-main-image" :src="image" alt="main"/>
        </div>
        <div class="main-image-container" v-for="(image, index) in item.variantsList[selectedColor].outImagesList" :key="'b' + index">
          <img v-if="image === selectedImage" class="product-main-image visible-image" :src="image" alt="main"/>
          <img v-else class="product-main-image" :src="image" alt="main"/>
        </div>
        <span class="product-info-title">{{ item.title }}</span>
        <span class="product-info-price">{{ item.price }} руб</span>
        <span class="product-color-block-title">Цвет:</span>
        <div class="product-colors-block">
          <button class="product-color-item" v-for="(color, index) in item.variantsList" :key="'c' + index" :title="color.inColor" :style="{backgroundColor: color.hex}" v-on:click="handleColor(index)"/>
        </div>
        <div class="product-info-spec-container">
          <div class="product-info-spec-row" v-for="(spec, index) in specItems" :key="'e' + index">
            <span class="spec-row-title">{{ spec.spec }}</span>
            <span class="spec-row-value">{{ spec.value }}</span>
          </div>
        </div>
        <div class="product-images-gallery absolute">
          <div v-for="(image, index) in item.variantsList[selectedColor].outImagesList" :key="'g' + index">
            <img v-if="image === selectedImage" class="product-gallery-image active-image active-image" alt="gallery" v-on:click="selectedImage=image" :src="image"/>
            <img v-else class="product-gallery-image" alt="gallery" v-on:click="selectedImage=image" :src="image"/>
          </div>
          <div v-for="(image, index) in item.variantsList[selectedColor].inImagesList" :key="'f' + index">
            <img v-if="image === selectedImage" class="product-gallery-image active-image active-image" alt="gallery" v-on:click="selectedImage=image" :src="image"/>
            <img v-else class="product-gallery-image" alt="gallery" v-on:click="selectedImage=image" :src="image"/>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/vue-loading.css';
import NavBar from "@/components/NavBar";
// import axios from "axios";
import ProductItems from "@/assets/items/ProductItems";
import {MetalDoorsApiClient} from "../../../proto/metal-doors_grpc_web_pb";

export default {
  name: "SingleProduct",
  components: {NavBar, Loading},
  created() {
    this.client = new MetalDoorsApiClient(process.env.VUE_APP_SERVER_URL, null, null)
    console.log(this.client)

    const { SingleMetalDoorRequest } = require("../../../proto/metal-doors_pb")
    let request = new SingleMetalDoorRequest
    request.setId(this.$route.params.id)
    this.client.getSingleMetalDoor(request, {}, (err, response) => {
      if (err) {
        this.error = true
      } else {
        this.item = response.toObject()
        this.navBarLink = response.toObject().title
        this.specItems = ProductItems(this.item, this.selectedColor);
        this.loading = false

        this.selectedImage = this.item.variantsList[this.selectedColor].outImagesList[0]


        console.log(this.item)
      }
    })
  },
  data: function () {
    return {
      item: null,
      navBarLink: "",
      selectedColor: 0,
      selectedImage: "",
      specItems: null,
      error: false,
      loading: true,
    }
  },

  mounted () {

  //   axios
  //       .get(`http://localhost:5008/api/doors/${this.$route.params.id}`)
  //       // .get(`http://143.198.109.211:3000/api/doors/${this.$route.params.id}`)
  //       .then(response => (this.item = response.data.result))
  //       .then(response => {this.specItems = ProductItems(response, this.selectedColor); this.selectedImage=this.item.variants[this.selectedColor].in_images[0].image_url})
  //       .catch(err => {console.log(err); this.error = true})
  //       .finally(() => this.loading = false)
  //
  },
  methods: {
    handleColor: function (index) {
      this.selectedColor = index
      this.specItems = ProductItems(this.item, index)
      this.selectedImage = this.item.variantsList[index].inImagesList[0]
    }
  },

}
</script>

<style scoped>
  @import "../../assets/css/base/page.css";
  @import "../../assets/css/base/product.css";
</style>