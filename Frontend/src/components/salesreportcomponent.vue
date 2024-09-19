<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="12">
          <h2>Sales Report</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-menu
            v-model="menuFrom"
            :close-on-content-click="false"
            :nudge-right="40"
            transition="scale-transition"
            offset-y
            outlined
          >
            <template v-slot:activator="{ on }">
              <v-text-field
                v-model="fromDate"
                label="From Date"
                readonly
                prepend-icon="mdi-calendar"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker
              v-model="fromDate"
              @input="menuFrom = false"
            ></v-date-picker>
          </v-menu>
        </v-col>
        <v-col cols="12" md="4">
          <v-menu
            v-model="menuTo"
            :close-on-content-click="false"
            :nudge-right="40"
            transition="scale-transition"
            offset-y
            outlined
          >
            <template v-slot:activator="{ on }">
              <v-text-field
                v-model="toDate"
                label="To Date"
                readonly
                prepend-icon="mdi-calendar"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker
              v-model="toDate"
              @input="menuTo = false"
            ></v-date-picker>
          </v-menu>
        </v-col>
        <v-col cols="12" md="4">
          <v-btn color="primary" @click="search">Search</v-btn>
          <v-btn color="primary" @click="downloadCSV">Download</v-btn>
        </v-col>
      </v-row>
      <div v-if="showTable">
        <v-row>
          <v-col cols="12">
            <v-text-field
              v-model="searchText"
              label="Search"
              clearable
              solo-inverted
              prepend-inner-icon="mdi-magnify"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-data-table
              :headers="headers"
              :items="uniqueData"
              :search="searchText"
              :items-per-page="5"
            ></v-data-table>
          </v-col>
        </v-row>
      </div>
    </v-container>

    <v-snackbar
      v-model="snackbar"
      :timeout="3000"
      :color="snackbarColor"
      bottom
      center
    >
      {{ snackbarMessage }}
      <v-btn color="black" text @click="snackbar = false">Close</v-btn>
    </v-snackbar>

    <v-dialog v-model="showDialog" max-width="280">
      <template v-slot:activator="{}"></template>
      <v-card>
        <v-card-title class="body-2 text-center">
          Select From and To Date
        </v-card-title>
        <v-card-actions>
          <v-btn color="error ml-16" @click="showDialog = false">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import EventService from "../service/Eventservice";
import Papa from "papaparse";
export default {
  data() {
    return {
      menuFrom: false,
      menuTo: false,
      fromDate: null,
      toDate: null,
      searchText: "",
      headers: [
        { text: "Bill No", value: "bill_no" },
        { text: "Bill Date", value: "bill_date" },
        { text: "Medicine Name", value: "medicine_name" },
        { text: "Quantity", value: "quantity" },
        { text: "Amount", value: "amount" },
      ],
      showTable: false,
      showDialog: false,
      uniqueData: [],
      originalData:[],
      snackbar: false,
      snackbarMessage: "",
      snackbarColor:''
    };
  },
  // computed: {
  //   billMasters() {
  //     return this.$store.state.billmaster;
  //   },
  //   billDetails() {
  //     return this.$store.state.billdetails;
  //   },
  //   filteredData() {
  //     if (!this.fromDate || !this.toDate) return [];

  //     const fromDate = new Date(this.fromDate);
  //     const toDate = new Date(this.toDate);

  //     const filteredBillMasters = this.billMasters.filter((master) => {
  //       const billDate = new Date(master.date);
  //       // Convert bill date to the same format as fromDate and toDate
  //       return billDate >= fromDate && billDate <= toDate;
  //     });

  //     const result = [];
  //     filteredBillMasters.forEach((master) => {
  //       const details = this.billDetails.filter(
  //         (detail) => detail.billno === master.billNo
  //       );
  //       details.forEach((detail) => {
  //         result.push({
  //           billNo: master.billNo,
  //           billDate: master.date,
  //           medicineName: detail.medicinename,
  //           quantity: detail.quantity,
  //           amount: detail.amount,
  //         });
  //       });
  //     });

  //     return result;
  //   },
  // },
  methods: {
    search() {
      if (this.fromDate == null || this.toDate == null) {
        this.snackbarMessage = "Please enter dates";
        this.snackbar = true;
        this.snackbarColor="red"
        return
      }

      if(this.fromDate>this.toDate){
        this.snackbarMessage = "dates are not valid";
        this.snackbar = true;
        this.snackbarColor="red"
        return

      }
       else {
        console.log(this.fromDate);
        console.log(this.toDate);

        const struct = {
          from_date: this.fromDate,
          to_date: this.toDate,
        };

        EventService.FetchSales(struct)
          .then((response) => {
            // console.log(response)

            if (
              response.data.status == "E" ||
              response.data.sales_arr === null
            ) {
              this.snackbarMessage = "no sales for selected dates";
              this.snackbarColor="red"
              this.showTable = false;
              this.snackbar = true;
              this.fromDate = null;
              this.toDate = null;
              return;
            } else{
              this.filterUniqueData(response.data.sales_arr)
              this.showTable=true
            }
          })
          .catch((error) => {
            console.log(error);
          });
      }
    },

    filterUniqueData(input) {
          const seen = new Set();
          this.uniqueData = input.filter(item => {
            const key = `${item.medicine_name}-${item.bill_no}`;
            return seen.has(key) ? false : seen.add(key);
          });

     
        },
    downloadCSV() {
      if (this.fromDate == null || this.toDate == null) {
        this.snackbarMessage = "Please enter dates for download";
        this.snackbar = true;
        this.snackbarColor="red"
        return;
      }
      const csvContent = Papa.unparse(this.uniqueData, {
        header: true,
      });

      const blob = new Blob([csvContent], { type: "text/csv;charset=utf-8;" });

      const link = document.createElement("a");
      if (link.download !== undefined) {

        const url = URL.createObjectURL(blob);
        link.setAttribute("href", url);
        link.setAttribute("download", "data.csv");

        document.body.appendChild(link);
        link.click();

        document.body.removeChild(link);
        URL.revokeObjectURL(url);
        this.snackbar = true;
        this.snackbarMessage = "download successfully";
        this.snackbarColor="green"
      } else {
        this.snackbar = true;
        this.snackbarMessage =
          "Your browser does not support downloading files";
          this.snackbarColor="red"
      }
    },
  },

  
};
</script>
