<template>
  <v-card>
    <v-card-title>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :items="stock"
      :search="search"
    ></v-data-table>
  </v-card>
</template>
<script>
import EventService from "../service/Eventservice";

export default {
  name: "StockView",
  data() {
    return {
      search: "",
      dialog: false,
      items: [],
      headers: [
        { text: "Medicine Name", value: "medicine_name" },
        { text: "Brand", value: "brand" },
        { text: "Quantity", value: "quantity" },
        { text: "Unit price", value: "unit_price" },
      ],
    };
  },
  created() {
    EventService.sendStockview()
      .then((response) => {
        // console.log(response.data);

        this.items = response.data.stock_view_details;
      })
      .catch((error) => {
        console.log(error);
      });
  },
  computed: {
    stock() {
      return this.items;
    },
  },
};
</script>
