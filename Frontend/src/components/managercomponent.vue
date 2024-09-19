<template>
  <div id="app">
    <template>
      <v-app>
        <v-app-bar app dark color="primary">
          <v-toolbar-title class="d-flex align-center">
         <svg-icon  type="mdi" :path="path"></svg-icon>
         <span class="ml-2">MANAGER</span>
         </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn text></v-btn>
          <v-btn text @click="dashboardmethod">   <v-icon>mdi-view-dashboard</v-icon>dashboard</v-btn>
          <v-btn text @click="stockviewmethod"> <v-icon>mdi-view-list</v-icon>stockview</v-btn>
          <v-btn text @click="stockentrymethod">  <v-icon>mdi-plus</v-icon>stockentry</v-btn>
          <v-btn text @click="salesreportmethod">  <v-icon>mdi-file-chart</v-icon>salesreport</v-btn>
     
          <v-btn text @click="logoutDialog=true"><v-icon>mdi-logout</v-icon>Logout</v-btn>
          <v-dialog v-model="logoutDialog" max-width="500px">
      <v-card>
        <v-card-title class="headline">Confirm Logout</v-card-title>
        <v-card-text>Are you sure you want to leave?</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="logoutmethod">Yes</v-btn>
          <v-btn color="red darken-1" text @click="logoutDialog = false">No</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
        </v-app-bar>

        <v-main>
          <Stockview v-if="stockviewshow"></Stockview>
          <Stockentry v-if="stockentryshow"></Stockentry>
          <Salesreport v-if="salesreportshow"></Salesreport>
          <Total v-if="totalshow"></Total>
        </v-main>
      </v-app>
    </template>
  </div>
</template>
<script>
import Stockview from "./stockviewcomponent.vue";
import Stockentry from "./stockentrycomponent.vue";
import Salesreport from "./salesreportcomponent.vue";
import Total from "./totalnetcomponent.vue";
import EventService from "../service/Eventservice";

import SvgIcon from '@jamescoyle/vue-icon';
import { mdiAlphaMBoxOutline } from '@mdi/js';

export default {
  data() {
    return {
      stockviewshow: false,
      stockentryshow: false,
      salesreportshow: false,
      totalshow: true,
      path: mdiAlphaMBoxOutline,
      logoutDialog:false,
      resmsg:''
    };
  },
  components: {
    Stockview,
    Stockentry,
    Salesreport,
    Total,
    SvgIcon
  },

    mounted() {
      // console.log("billercomp", this.login_history_id);
    },
 
  methods: {
    stockviewmethod() {
      this.stockviewshow = true;
      this.stockentryshow = false;
      this.salesreportshow = false;
      this.totalshow = false;
    },
    stockentrymethod() {
      this.stockentryshow = true;
      this.stockviewshow = false;
      this.salesreportshow = false;
      this.totalshow = false;
    },
    salesreportmethod() {
      this.salesreportshow = true;
      this.stockentryshow = false;
      this.stockviewshow = false;
      this.totalshow = false;
    },
    dashboardmethod() {
      this.totalshow = true;
      this.salesreportshow = false;
      this.stockentryshow = false;
      this.stockviewshow = false;
    },
    logoutmethod() {
     
      const structs ={
        login_history_id:this.login_history_id
      }

      EventService.sendLogoutHistory(structs).then((response) => {
        // console.log(response)
        this.resmsg=response.data.msg

        


      }).catch((error)=>{
        console.log(error)
      })


      this.$router.push("/")
   
   
    },
    },
  props: ["login_history_id"],
};
</script>
<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.v-application {
  background-color: #f5f5f5 !important;
}
</style>
