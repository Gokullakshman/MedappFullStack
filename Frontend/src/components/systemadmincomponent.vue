<template>
    <div id="app">
        <template>
  <v-app>


    <v-app-bar app dark color="primary">
 
      <v-toolbar-title class="d-flex align-center">
        <svg-icon type="mdi" :path="path"></svg-icon>
         <span class="ml-2">SYSTEMADMIN</span>
         </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text></v-btn>
      <v-btn text @click="dashboardmethod"> <v-icon>mdi-view-dashboard</v-icon>Dashboard</v-btn>
      <v-btn text @click="addusermethod">      <v-icon>mdi-account-plus</v-icon>adduser</v-btn>
      <v-btn text @click="loginhistorymethod"> <v-icon>mdi-history</v-icon>loginhistory</v-btn>
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

    <v-container>
        
 
     <Adduser v-if="addusershow"></Adduser>
     <Loginhistory v-if="loginhistoryshow"></Loginhistory>
     <Dashboard v-if="dashboardshow"></Dashboard>




       

        
        

    </v-container>
  </v-app>
</template>

    </div>
</template>
<script>
import Adduser from './addusercomponent.vue'
import Loginhistory from './loginhistorycomponent.vue'
import Dashboard from './welcomecomponent.vue'
import EventService from "../service/Eventservice";
import SvgIcon from '@jamescoyle/vue-icon';
import { mdiAlphaSBoxOutline } from '@mdi/js';

export default{
    data(){
        return{
            addusershow:false,
            loginhistoryshow:false,
            dashboardshow:true,
            path: mdiAlphaSBoxOutline,
            logoutDialog:false,
            resmsg:''
   

        }
    },
    mounted() {
   
  //  console.log("systemadmin", this.login_history_id)
 },
    components:{
        Adduser,Loginhistory,Dashboard,SvgIcon
       
    },
    methods:{
        addusermethod(){
            this.addusershow=true
            this.loginhistoryshow=false
            this.dashboardshow=false
        },
        loginhistorymethod(){
            this.loginhistoryshow=true
            this.addusershow=false
            this.dashboardshow=false
        },
        dashboardmethod(){
          this.dashboardshow=true
          this.loginhistoryshow=false
          this.addusershow=false
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
   
      this.$router.push("/");
    },
        
    },
    props:['login_history_id']
}
</script>
<style>

.center-content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%; 
}


