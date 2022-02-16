<template>
  <div class="login-container">
    <!--      <h4>Login form</h4>-->
    <div class="login-form-block">
      <form class="login-form" @submit.prevent @submit="submit">
        <div class="login-form-rows">
          <input
              type="email"
              class="login-form-row"
              id="txtEmail"
              aria-describedby="emailHelp"
              v-model="user.email"
              placeholder="E-mail"
          />

          <input
              type="password"
              class="login-form-row"
              id="txtPassword"
              v-model="user.password"
              placeholder="Password"
          />
          <button type="submit" class="login-form-button">Войти</button>
        </div>

      </form>
    </div>

  </div>
</template>

<script>
import {AuthServiceClient} from "../../../proto/auth_service_grpc_web_pb";
import {mapGetters, mapActions} from "vuex";

export default {
  name: "LoginForm",
  data() {
    return {
      user: {
        email: "",
        password: ""
      }
    }
  },
  computed: {
    ...mapGetters("auth", {
      getLoginApiStatus: "getLoginApiStatus",
    }),
  },
  methods: {
    ...mapActions("auth", {
      actionLoginApi: "loginApi",
    }),
    submit() {
      this.client = new AuthServiceClient(process.env.VUE_APP_SERVER_URL, null, null)
      const {LoginRequest} = require("../../../proto/auth_service_pb");
      let request = new LoginRequest;
      request.setEmail(this.user.email)
      request.setPassword(this.user.password)
      this.client.login(request, {}, (err, response) => {
        if (err) {
          alert(err)
        } else {
          localStorage.setItem("authorization", response.toObject().accessToken)
          this.$router.push("/console/dashboard")
        }
      })
    }
  }
}
</script>

<style scoped>
@import "../../assets/css/auth/login.css";
</style>