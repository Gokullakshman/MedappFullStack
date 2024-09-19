<template>
    <v-container fluid>
      <v-row justify="center">
        <v-col cols="12" sm="8" md="6">
          <v-card class="elevation-12">
            <v-card-text class="text-center">
              <h1>Unsold Medicine Amount</h1>
              <p>Total Unsold Medicine Amount: {{ totalUnsoldAmount }}</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </template>
  
  <script>
  export default {
    data() {
      return {
        totalUnsoldAmount: 0
      };
    },
    created() {
      this.calculateUnsoldAmount();
    },
    methods: {
      calculateUnsoldAmount() {

        
        
  }}};
  </script>
  
  <style scoped>
</style>
  