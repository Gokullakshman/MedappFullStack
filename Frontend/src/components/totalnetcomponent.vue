<template>
    <div id="app">
   
       <v-container fluid fill-height>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="9" md="5">
    <v-window v-model="window" show-arrows>
      <v-window-item v-for="(item, index) in array" :key="index">
        <v-card class="d-flex justify-center align-center" height="180px">
          <v-card-text class="text-center">
            <v-icon size="64" >{{item.icon}}</v-icon>
            <h1 class="headline mb-4" color="primary">{{item.name}}</h1>
            <p class="text-h4">{{ item.amount }}</p>
          </v-card-text>
        </v-card>
      </v-window-item>
    </v-window>
  </v-col>
</v-row>  
  
    <!-- <v-row align="center" justify="center">
      <v-col cols="12" sm="9" md="5">
        <v-card class="elevation-12 gradient-background white--text text-center">
          <v-card-text class="text-center">
            <v-icon size="64" >mdi-cash</v-icon>
            <h1 class="headline mb-4" color="primary">Total Sales</h1>
            <p class="body-1">{{ totalNetPrice }}</p>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>

  <v-container fluid>
    <v-row justify="center">
      <v-col cols="12" sm="8" md="6">

        <v-card class="elevation-12 gradient-background white--text text-center">
          <v-card-text>
            <v-avatar size="100">
              <v-icon size="64">mdi-pill</v-icon>
            </v-avatar>
            <h1 class="headline mb-4">Unsold Medicine Amount</h1>
            <p class="body-1">Total Unsold Medicine Amount: {{ unsoldmedicineamount }}</p>
          </v-card-text>
          <v-card-actions>
           
          </v-card-actions>
        </v-card>
        
      </v-col>
    </v-row>  -->
  </v-container>
      <!-- <Unsold></Unsold> -->
      <Chart></Chart>
      <Chart1></Chart1>
      
    </div>
  </template>
  
  <script>
  // import Unsold from './unsoldcomponent.vue'
  import EventService from "../service/Eventservice.js"
  import Chart from "./chartcomponent.vue"
   import Chart1 from "./weeklysalescomponent.vue"
  
  
  export default {
    components: {
      // Unsold
      Chart,Chart1
    },
    props: ['message'],
    data() {
      return {
        totalNetPrice: 0 ,
        unsoldmedicineamount:0,
          array:[],
          length: 3,
          window: 0,
      }
    },
    
    mounted(){
   

      EventService.sendUnsold()
      .then((response) => {
        this.array.push({name:"TODAY SALES",amount:response.data.today_net,icon:"mdi-cash"})
        this.array.push({name:"UNSOLD MEDICINE AMOUNT",amount:response.data.unsold,icon:"mdi-pill"})
     
        this.unsoldmedicineamount=response.data.unsold 
        this.totalNetPrice=response.data.today_net
       
       
      })
      .catch((error) => {
        console.log(error);
      });




    }
  }
  </script>
  