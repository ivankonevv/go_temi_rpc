<template>
  <div class="main-console-content-container">
    <div class="main-add-container">
      <span class="add-container-title">Создание новой двери</span>
      <div class="add-form-container">
        <form class="add-form" @submit.prevent>
          <div class="add-parent-form-row" v-for="(item, index) in items" :key="index" v-show="item.child">
            <span class="add-form-row-title">{{ item.title }}:</span>
            <div class="add-child-form-container">
              <div class="add-child-form-row" v-for="(child, i) in item.child" :key="i">
                <span class="add-form-row-title">{{ child.title }}</span>
                <input class="add-form-row-input-container add-child-input" v-model.number="fullSpec.price[child.state_name]" :type="child.type"/>
              </div>
            </div>
          </div>
          <div class="add-form-row" v-for="(item, index) in items" :key="index+20"
               v-show="item.type === 'checkbox' && !item.child">
            <span class="add-form-row-title">{{ item.title }}</span>
            <input class="add-form-row-input-container" :type="item.type" v-model="fullSpec.for_catalog"/>
          </div>
          <div class="add-form-row" v-for="(item, index) in items" :key="index+40"
               v-show="item.type !== 'checkbox' && !item.child && !item.spec">
            <span class="add-form-row-title">{{ item.title }}</span>
            <input class="add-form-row-input-container" v-model="fullSpec[item.state_name]" :type="item.type"/>
          </div>
          <div class="add-form-row" v-for="(item, index) in items" :key="index+60"
               v-show="item.spec">
            <span class="add-form-row-title">{{ item.title }}</span>
            <input class="add-form-row-input-container" v-model.number="fullSpec.specification[item.state_name]" :type="item.type"/>
          </div>
          <button class="add-submit-main-form absolute" type="submit">ГОТОВО</button>
        </form>
      </div>
      <div class="form-delim-line"/>
      <div class="add-color-container">

        <div class="popup-box" v-for="(cur, index) in variantBlocks" :key="index">
          <div class="popup-h-color" v-bind:class="{hidden: cur.hidden}"/>
<!--          <span class="popup-h-title" v-show="cur.hidden" v-bind:class="{hidden: cur.hidden}">Какой-то крутой цвет</span>-->
        <div class="popup-images-upload-container" v-bind:class="{hidden: cur.hidden}">
          <div class="popup-images-upload-block" v-bind:class="{hidden: cur.hidden}">
<!--            <div class="popup-h-color" v-show="cur.hidden===true" v-bind:/>-->
            <div class="popup-images-upload-up-row">
              <span class="popup-image-upload-title">Внешний цвет:</span>
              <input class="popup-image-upload-input popup-upload-row-focus" v-model="cur.data.out_color"/>
            </div>
            <div class="popup-images-upload-middle-row">
              <div class="popup-image-upload-button">
                <label class="image-input-label">
                  <inline-svg class="popup-upload-icon"
                              :src="require('../../../assets/icons/console_add_upload_image.svg')"/>
                  <input class="image-input" type="file" id="image" name="img" accept="image/*" multiple
                         @change="(e) => setOutImages(e, index)"/>

                </label>
              </div>
              <div class="popup-image-uploaded-container">
                <div class="popup-image-uploaded-wrapper">
                  <div class="popup-uploaded-image" v-for="(image, i) in outImages[index]" :key="i"
                       :style="{backgroundImage: `url('${image}')`}"/>
                </div>
                <div class="popup-delim-line-horizontal"/>
              </div>
            </div>
<!--            <input class="popup-color-picker" type="color" v-model="cur.data.color_hex"/>-->
          </div>
          <div class="popup-images-upload-block"  v-bind:class="{hidden: cur.hidden}">
            <div class="popup-images-upload-up-row">
              <span class="popup-image-upload-title">Внутренний цвет:</span>
              <input class="popup-image-upload-input popup-upload-row-focus" v-model="cur.data.in_color"/>
              <label>
                <input type="color">
                <div class="circle"/>
              </label>

            </div>
            <div class="popup-images-upload-middle-row">
              <div class="popup-image-upload-button">
                <label class="image-input-label">
                  <inline-svg class="popup-upload-icon"
                              :src="require('../../../assets/icons/console_add_upload_image.svg')"/>
                  <input class="image-input" type="file" id="img" name="img" accept="image/*" multiple
                         @change="(e) => setInImages(e, index)"
                  />
                </label>
              </div>
              <div class="popup-image-uploaded-container">
                <div class="popup-image-uploaded-wrapper">
                  <div class="popup-uploaded-image" v-for="(image, i) in inImages[index]" :key="i"
                       :style="{backgroundImage: `url('${image}')`}"/>
                </div>
                <div class="popup-delim-line-horizontal"/>
              </div>
            </div>
          </div>

        </div>
                  <div class="popup-delim-line left"/>
        </div>


        <div class="popup-delim-line center"/>
        <div class="popup-images-block-width-container">
          <span class="popup-width-title">
            Размеры
          </span>
          <div class="popup-width-input-block" v-for="(count, index) in width" :key="index">
            <span class="popup-input-row-title">Ширина: </span>
            <input class="popup-width-input popup-upload-row-focus" :name="index" type="number"/>

            <select class="popup-input-select popup-upload-row-focus">
              <option value="L">L</option>
              <option value="R">R</option>
            </select>
          </div>
          <button class="popup-width-button" @click="width < 4 ? width = width + 1 : null">ЕЩЕ</button>
        </div>
        <div class="popup-delim-line right"/>
        <button class="popup-add-one-more-color-button " @click="logg">ASfdfasfasf</button>
        <button class="popup-add-one-more-color-button " @click="submit">Отправить</button>
        <button class="popup-add-one-more-color-button absolute" @click="addVariant">Добавить</button>
      </div>
    </div>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: "AddNewProduct",
  data() {
    return {
      width: 1,
      inImages: [],
      outImages: [],
      variantBlocks: [
        {
          hidden: false,
          data: {
            in_color: '',
            out_color: '',
            in_images: [],
            color_hex: '',
            out_images: [],
            sizes: [],
          }
        }
      ],
      fd: new FormData,
      // variant: {
      //   in_color: '',
      //   out_color: '',
      //   in_images: [],
      //   out_images: [],
      //   sizes: [],
      // },

      fullSpec: {
        variants:[],
        for_catalog: false,
        number: "",
        title: "",
        price: {
          vip: 0,
          opt: 0,
          k_opt: 0,
          roz: 0
        },
        specification: {
          height: 2050,
          metal: 0,
          leaf: 0,
          upper_lock: "",
          lower_lock: "",
          valve: "",
          armor: "",
          cylinder: "",
          peephole: "",
          insulation: "",
          sealer: "",
          handle: ""
        }
      }
    }
  },
  methods: {
    // submitImage () {
    //   if (this.inI) {
    //     let image_links = []
    //     for (let i = 0; i < this.inI; i++) {
    //       this.toUploadData.append('color')
    //     }
    //   }
    // },
    addVariant () {
      this.variantBlocks.push({
        hidden: false,
        data: {
          in_color: '',
          out_color: '',
          in_images: [],
          out_images: [],
          sizes: [],
        }
      })
      let len = this.variantBlocks.length - 2;
      this.variantBlocks[len].hidden = true

    },
    logg() {
      console.log(this.fullSpec)
      console.log(this.variantBlocks)
    },
    async submit() {
      let item = this.fullSpec

      for (let i = 0; this.variantBlocks.length > i; i++) {
        item.variants.push(this.variantBlocks[i].data)
      }

      axios({
        method: 'post',
        url: 'http://localhost:5008/api/doors',
        data: item
      })
      .then(res => console.log(res.data))
    },
    setInImages (event, index) {
      let picURL = []
      for (let i = 0; i < event.target.files.length; i++) {
        picURL.push(URL.createObjectURL(event.target.files[i]));
        this.fd.append('color', this.variantBlocks[index].data.in_color)
        this.fd.append('image', event.target.files[i])
        axios({
          method: 'post',
          url: 'http://localhost:5008/upl',
          data: this.fd
        })
        .then(res => {
          this.variantBlocks[index].data.in_images.push({image_url: res.data.image_url})
        })
        .catch(err => {
          console.log(err)
        })
        this.fd.delete('color')
        this.fd.delete('image')
      }
      this.inImages.push(picURL)
    },
    setOutImages(event, index) {
      let picURL = []
      for (let i = 0; i < event.target.files.length; i++) {
        picURL.push(URL.createObjectURL(event.target.files[i]));
        this.fd.append('color', this.variantBlocks[index].data.out_color)
        this.fd.append('image', event.target.files[i])
        axios({
          method: 'post',
          url: 'http://localhost:5008/upl',
          data: this.fd
        })
            .then(res => {
              this.variantBlocks[index].data.out_images.push({image_url: res.data.image_url})
            })
            .catch(err => {
              console.log(err)
            })
        this.fd.delete('color')
        this.fd.delete('image')
      }
      this.outImages.push(picURL)
    }
  },
  props: ['items']
}
</script>

<style scoped>
@import "../../../assets/css/auth/console_add.css";
</style>