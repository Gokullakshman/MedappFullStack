<template>
  <div id="app">
    <template>
      <v-app>
        <v-app-bar app dark color="primary">
          <v-toolbar-title class="d-flex align-center">
            <svg-icon type="mdi" :path="path"></svg-icon>
      <span class="ml-2">BILLER</span>
    </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn text></v-btn>
          <v-btn text @click="dashboardmethod"><v-icon>mdi-view-dashboard</v-icon>dashboard</v-btn>
          <v-btn text @click="stockviewmethod">  <v-icon>mdi-pill</v-icon>stockview</v-btn>
          <v-btn text @click="billentrymethod" :login_history_id="login_history_id" :user_id="user_id" :login_id="login_id">  <v-btn icon>
              <v-icon>mdi-cash-register</v-icon>
            </v-btn>Billentry</v-btn>
          <v-btn text @click="logoutDialog = true"><v-icon>mdi-logout</v-icon>Logout</v-btn>

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
          <Billentry v-if="billentryshow" ></Billentry>
          <Dashboard v-if="dashboardshow" :sales="sales"    :flag="flag" ></Dashboard>
        </v-main>
      </v-app>
    </template>
  </div>
</template>
<script>
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiAlphaBBox } from '@mdi/js';
import Billentry from "./billentrycomponent.vue";
import Stockview from "./stockviewcomponent.vue";
import Dashboard from "./todaycomponent.vue";
import EventService from "../service/Eventservice";
export default {
  data() {
    return {
      stockviewshow: false,
      billentryshow: false,
      dashboardshow: true, 
      loginn_id:0,
      path: mdiAlphaBBox,
      logoutDialog:false,
      resmsg:''
    
    };
  },

  
  components: {
    Billentry,
    Stockview,
    Dashboard,
    SvgIcon
  },
  methods: {
    stockviewmethod() {
      this.stockviewshow = true;
      this.billentryshow = false;
      this.dashboardshow = false;
      console.log(this.userId);
    },
    billentrymethod() {
      this.billentryshow = true;
      this.stockviewshow = false;
      this.dashboardshow = false;
    },
    dashboardmethod() {
      (this.billentryshow = false),
        (this.dashboardshow = true),
        (this.stockviewshow = false);
    },
    logoutmethod() {
      const structs ={
        login_history_id:this.login_history_id
      }
      console.log(this.login_history_id)

      EventService.sendLogoutHistory(structs).then((response) => {
       this.resmsg=response.data.msg

        


      }).catch((error)=>{
        console.log(error)
      })
     
      this.$router.push("/")
   
    },
  },

 props:['login_history_id','login_id','sales','user_id','flag']
  

};
</script>
