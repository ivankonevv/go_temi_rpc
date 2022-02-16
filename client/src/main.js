import Vue from 'vue'
import App from "./App";
import VueRouter from "vue-router";
import VueCookies from 'vue-cookies'
import Notifications from 'vue-notification'


import VueGlide from 'vue-glide-js'
import InlineSvg from 'vue-inline-svg';
import store from "./store/index";



import Home from "@/Home";
import Products from "@/Products";
import Product from "@/Product";
import Blog from "@/Blog";
import Contacts from "@/Contacts";

import Login from "@/Login";
import DashboardPage from "@/components/auth/dashboard/DashboardPage";
import AddChoicePage from "@/components/auth/add/AddChoicePage";
import AddPage from "@/components/auth/add/AddPage";
import ShippingAndPayment from "@/ShippingAndPayment";
import Installment from "@/Installment";

import Pagination from 'vue-pagination-2';
import WoodColors from "@/WoodColors";
import NotFound from "@/NotFound";



// const NotFound = { template: '<p>Page not found</p>'}

Vue.use(VueRouter)
Vue.use(VueCookies)
Vue.use(VueGlide)
Vue.use(Notifications)
Vue.component('inline-svg', InlineSvg);
Vue.component('pagination', Pagination);



const routes = [
  { path: '/', component: Home },
  { path: '/blog', component: Blog },
  { path: '/contacts', component: Contacts },
  { path: '/metal-doors', component: Products },
  { path: '/wood-doors', component: Products },
  { path: '/wood-doors/colors', component: WoodColors },
  { path: '/shipping-and-payment', component: ShippingAndPayment },
  { path: '/installment', component: Installment },
  { path: '/product/:id', component: Product, name: 'product' },

  { path: '/console/login', component: Login},
  { path: '/console/dashboard', component: DashboardPage, meta: { requiredAuth: true }},
  { path: '/console/add', component: AddChoicePage, meta: { requiredAuth: true }},
  { path: '/console/add/metal', component: AddPage, meta: { requiredAuth: true }},
  { path: '/console/add/wood', component: AddPage, meta: { requiredAuth: true }},
  { path: '/console/add/inventory', component: AddPage, meta: { requiredAuth: true }},
  // { path: '/console/dashboard', component: DashboardPage},
  // { path: '/console/add', component: AddChoicePage},
  // { path: '/console/add/metal', component: AddPage},
  // { path: '/console/add/wood', component: AddPage},
  // { path: '/console/add/inventory', component: AddPage},
  { path: "*", component: NotFound},
]

const router = new VueRouter({
  routes: routes,
  mode: 'history'
})

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiredAuth) {
    let cookie = Vue.$cookies.get('jwt')
    if (!cookie) {
      return next({ path: '/console/login'});
    } else {
      await store.dispatch('auth/userProfile');
      let userProfile = store.getters["auth/getUserProfile"]
      if (userProfile.email === '') {
        return next({ path: '/console/login'} )
      }
      return next()
    }


    // if (userProfile.id === "") {
    //   await store.dispatch("auth/userProfile");
    //   userProfile = store.getters["auth/getUserProfile"];
    //   if (userProfile.id === "") {
    //     return next({ path: "/console/login" });
    //   } else {
    //     return next();
    //   }
    // }
  }
  return next();
});



Vue.config.productionTip = false

new Vue({
  // store,
  router,
  store,
  render: (h) => h(App),
}).$mount('#app')
