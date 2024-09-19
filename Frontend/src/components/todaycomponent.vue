<template>
  <div id="app">
    <v-container fluid>
      <v-row justify="center">
        <v-col cols="12" sm="6" md="4">
          <v-card   bg-color="grey-lighten-7"
      color="blue" class="elevation-15  ">
            <v-card-title class="text-h5">    
              <v-icon>mdi-chart-line</v-icon>
            Your Today Sales</v-card-title>
            <v-card-text class="text-center" style="padding: 70px">
              <h1>{{ sales }}</h1>
              <v-icon v-if="flag == 1" bg-color="green-dark" color="green">mdi-arrow-up-bold</v-icon>
              <v-icon v-else color="red">mdi-arrow-down-bold</v-icon>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      
    </v-container>
  </div>
</template>

<script>
export default {
  props: ["sales","flag"],
  data() {
    return {
      userid: "",
      usertotal: 700,
      totalNetPrice: 0,
    
    };
  },

mounted(){
  // console.log("mounted"+this.flag)
  // console.log("mounted"+this.sales)  
}

};
</script>

<style scoped>
.v-card-title {
  font-size: 1.5rem;
  
}

</style>
