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
        <!--        <div v-for="item in result" :key="item.id">{{ item }}</div>-->
        <!--        <ProductCard :item="this.result[0]"/>-->
        <ProductCard v-for="item in paginated" :key="item.id" :item="item"/>
        <!--          <nav role="navigation" class="pagination-block"-->
        <!--              aria-label="pagination"-->
        <!--          >-->
        <!--            <ul class="pagination-container">-->
        <!--              <li class="pagination-link"><a @click="prev"></a></li>-->
        <!--              <li class="pagination-link"><span class="pagination-link go-to has-text-orange" aria-label="Goto page 1">{{ current }}</span></li>-->
        <!--              <li class="pagination-link"><a @click="next()">J</a></li>-->
        <!--&lt;!&ndash;              <li>{{ result.length / 3 }}</li>&ndash;&gt;-->
        <!--            </ul>-->
        <!--          </nav>-->
      </template>

    </div>
    <pagination :options="paginationOptions" :records="result.length" v-model="page" :per-page="3"></pagination>

  </div>
</template>

<script>
import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/vue-loading.css';
import NavBar from "@/components/NavBar";
import ProductCard from "@/components/products/ProductCard";
// import { credentials } from "@grpc/grpc-js"
// import { grpc } from "grpc-web";

// import Pagination from "@/components/Pagination";

export default {
  name: "ProductList",
  components: {ProductCard, NavBar, Loading},
  created: function () {
    let service = require('../../../proto/metal-doors_grpc_web_pb')

    const client = new service.MetalDoorsApiClient(process.env.VUE_APP_SERVER_URL, null, null);
    console.log(client)
    const {PostsRequest} = require("../../../proto/metal-doors_pb");

    let request = new PostsRequest();
    let stream = client.getPosts(request, {});
    stream.on('data', (response, err) => {
      console.log(err)
      this.result.push(response.toObject().data);
      this.loading = false
      console.log(this.result)
    })
    const page = this.$route.query.page || 1;

    this.$nextTick(() => {
      this.currentPage = page
    })

  },
  watch: {
    $route(to) {
      this.$refs.console.value += `\nRoute Change detected: page: ${to.query.page} / tweetPages: ${to.query.tweetPages}`
    }
  },

  data: function () {
    return {
      navBarLink: "Входные двери",
      result: [],

      loading: true,
      error: false,
      posts: null,

      page: 1,
      current: 1,
      pageSize: 9,

      paginationOptions: {
        chunk: 4,
        // chunksNavigation: 'scroll',
        edgeNavigation: true,
      },
    }
  },
  methods: {
    prev() {
      this.current--;
    },
    next() {
      this.current++;
    },
    creds() {


    }
  },
  computed: {
    indexStart() {
      return (this.page - 1) * this.pageSize;
    },
    indexEnd() {
      return this.indexStart + this.pageSize;
    },
    paginated() {
      return this.result.slice(this.indexStart, this.indexEnd);
    }
    // currentPage: this.$router.push({query: {...this.$route.query, page: newPage}}).catch(() => {})

  },

}
</script>

<style scoped>
@import "../../assets/css/base/product_list.css";
@import "../../assets/css/base/page.css";
@import "../../assets/css/base/pagination.scss";
</style>