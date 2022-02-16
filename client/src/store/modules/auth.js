import axios from "axios";
import Vue from "vue";
import {UserServiceClient} from "../../../proto/user_service_grpc_web_pb";


const state = () => ({
    loginApiStatus: "",
    userProfile: {
        firstName: "",
        lastName: "",
        email: "",
        role: ""
    },
    logOut: false
});

const getters = {
    getLoginApiStatus(state) {
        return state.loginApiStatus;
    },
    getUserProfile(state) {
        return state.userProfile;
    },
    getLogout(state) {
        return state.logOut;
    }
};

const actions = {
    async loginApi({ commit }, payload) {
        const response = await axios
            .post("http://127.0.0.1:5008/api/signin",
                payload, {withCredentials: true, credentials: "include"})
            .catch((err) => {
                console.log(err);
            });

        if (response && response.data) {
            Vue.$cookies.set('jwt', response.data.access_token, '8h', '/', 'localhost', true, false)
            commit("setLoginApiStatus", "success");
        } else {
            commit("setLoginApiStatus", "failed");
        }
    },
    userProfile({ commit }) {
        this.client = new UserServiceClient(process.env.VUE_APP_SERVER_URL, null, null)
        const { GetUserRequest } = require("../../../proto/user_service_pb")
        let request = new GetUserRequest
        this.client.getUser(request, {"authorization": localStorage.getItem("authorization")}, (err, response) => {
            if (err) {
                console.log(err)
            } else {
                console.log(response.toObject())
                commit("setUserProfile", response.toObject())
            }
        })

            // const response = await axios
            //     .get("http://127.0.0.1:5008/api/user",
            //         {withCredentials: true, credentials: "include"})
            //     .catch((err) => {
            //         console.log(err);
            //     });
            // if (response && response.data) {
            //     if (response.data.result) {
            //         commit("setUserProfile", response.data.result)
            //     console.log('success!!')
            //
            //     } else {
            //         console.log('asdasdad')
            //     }
            //
            // }
    },
    async userLogout({ commit }) {
        console.log(Vue.$cookies.get('jwt'))
        Vue.$cookies.remove("jwt")
        if (!Vue.$cookies.get('jwt')) {
            commit('setLogout', true)
        } else {
            commit('setLogout', false)
        }
        // const response = await axios
        //     .post("http://127.0.0.1:5008/api/logout", {withCredentials: true, credentials: "include"})
        //     .catch((err) => {
        //         console.log(err);
        //     })
        // if (response && response.data) {
        //     commit("setLogout", true)
        // } else {
        //     commit("setLogout", false)
        // }
    }
};

const mutations = {
    setLoginApiStatus(state, data) {
        state.loginApiStatus = data;
    },
    setUserProfile(state, data) {
        const userProfile = {
            firstName: data.firstName,
            lastName: data.lastName,
            email: data.email,
            role: data.role
        }
        state.userProfile = userProfile
    },
    setLogout(state, payload) {
        state.logOut = payload;
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}