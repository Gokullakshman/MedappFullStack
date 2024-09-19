<template>
  <div id="app">
    <template>
      <v-app>
        <v-app-bar app dark color="blue">
          <v-toolbar-title class="d-flex align-center">
            <svg-icon type="mdi" :path="path"></svg-icon>
            <span class="ml-2">INVENTRY</span>
          </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn text></v-btn>
          <v-btn text @click="dashboardmethod"> <v-icon>mdi-view-dashboard</v-icon>Dashboard</v-btn>
          <v-btn text @click="stockviewmethod"><v-icon>mdi-view-list</v-icon>stockview</v-btn>
          <v-btn text @click="stockentrymethod"><v-icon>mdi-plus</v-icon>stockentry</v-btn>
          <v-btn text @click="logoutDialog = true"><v-icon>mdi-logout</v-icon>Logout</v-btn>

          <v-dialog v-model="logoutDialog" max-width="500px">
            <v-card>
              <v-card-title class="headline">Confirm Logout</v-card-title>
              <v-card-text>Are you sure you want to leave?</v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="green darken-1" text @click="logoutmethod"
                  >Yes</v-btn
                >
                <v-btn color="red darken-1" text @click="logoutDialog = false"
                  >No</v-btn
                >
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-app-bar>

        <v-container>
          <Stockview v-if="stockviewshow"></Stockview>
          <Stockentry v-if="stockentryshow"></Stockentry>
          <Dashboard v-if="dashboardshow"></Dashboard>
        </v-container>
      </v-app>
    </template>
  </div>
</template>
<script>
import Stockview from "./stockviewcomponent.vue";
import Stockentry from "./stockentrycomponent.vue";
import Dashboard from "./welcomecomponent.vue";
import EventService from "../service/Eventservice";
import SvgIcon from "@jamescoyle/vue-icon";
import { mdiAlphaIBox } from "@mdi/js";
export default {
  data() {
    return {
      stockviewshow: false,
      stockentryshow: false,
      dashboardshow: true,
      path: mdiAlphaIBox,
      logoutDialog:false,
      resmsg:''
    };
  },
  components: {
    Stockview,
    Stockentry,
    Dashboard,
    SvgIcon,
  },
  mounted() {
    // console.log("inventry ", this.login_history_id);
  },
  methods: {
    stockviewmethod() {
      this.stockviewshow = true;
      this.stockentryshow = false;
      this.dashboardshow = false;
    },
    stockentrymethod() {
      this.stockentryshow = true;
      this.stockviewshow = false;
      this.dashboardshow = false;
    },
    dashboardmethod() {
      this.dashboardshow = true;
      this.stockentryshow = false;
      this.stockviewshow = false;
    },
    logoutmethod() {
      const structs = {
        login_history_id: this.login_history_id,
      };

      EventService.sendLogoutHistory(structs)
        .then((response) => {
          this.resmsg=response.data.msg
        })
        .catch((error) => {
          console.log(error);
        });

      this.$router.push("/");
    },
  },
  props: ["login_history_id"],
};
</script>
