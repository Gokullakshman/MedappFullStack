<template>
  <v-container>
    <!-- Bill Entry Card -->
    <v-card>                    
      <v-card-title>Bill Entry</v-card-title>
      <v-card-text>
        <v-expansion-panels>
          <v-expansion-panel>
            <v-expansion-panel-header>
              <div class="d-flex justify-space-between align-center">
                <div>Medicine</div>
              </div>
            </v-expansion-panel-header>
            <v-expansion-panel-content>
              <v-container>
                <v-row>
                  <v-col cols="6">
                    <v-select
                      v-model="selectedMedicine"
                      :items="medicineOptions"
                      label="Medicine Name"
                    ></v-select>
                  </v-col>
                  <v-col cols="4">
                    <v-text-field
                      v-model="quantity"
                      label="Quantity"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="2">
                    <v-btn outlined color="blue" @click="addItem">Add</v-btn>
                  </v-col>
                </v-row>
              </v-container>
            </v-expansion-panel-content>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-card-text>
    </v-card>

    <!-- Display Card -->
    <v-card>
      <v-card-title>
        <div class="d-flex justify-space-between align-center">
          <div></div>
          <div>Bill No: {{ billno }}</div>
        </div>
      </v-card-title>
      <v-card-text>
        <div style="color: #3f51b5">
          Total: <span style="color: #3878ac">${{ total }}</span
          >, GST: <span style="color: #00bcd4">${{ gst }}</span
          >, Net Payable: <span style="color: #1e88e5">${{ netPrice }}</span>
        </div>
      </v-card-text>

      <v-card-actions>
        <div class="d-flex justify-space-between">
          <v-btn outlined color="blue" @click="openPreviewDialog"
            >Preview</v-btn
          >
          <v-btn outlined color="blue" @click="save">Save</v-btn>
          <v-btn outlined color="blue" blue @click="download">Download</v-btn>
          <v-snackbar v-model="snackbar">{{ text }}</v-snackbar>
        </div>
      </v-card-actions>
      <v-divider></v-divider>

      <v-data-table
        v-if="showTable"
        :items="display"
        :headers="headers"
        class="elevation-1"
      ></v-data-table>
    </v-card>

    <!-- Preview Dialog -->
    <v-dialog v-model="previewDialog" max-width="500">
      <v-card>
        <v-card-title>Preview</v-card-title>
        <v-card-text>
          <v-simple-table>
            <template v-slot:default>
              <thead>
                <tr>
                  <th class="text-left">Medicine Name</th>
                  <th class="text-left">Quantity</th>
                  <th class="text-left">Amount</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(item, index) in display" :key="index">
                  <td>{{ item.medicine_name }}</td>
                  <td>{{ item.quantity }}</td>
                  <td>{{ item.amount }}</td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
          <div class="text-right">Total: ${{ total }}</div>
          <div class="text-right">GST: ${{ gst }}</div>
          <div class="text-right">Netprice: ${{ netPrice }}</div>
        </v-card-text>
        <v-card-actions>
          <v-btn color="primary" @click="closePreviewDialog">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

   

    <v-snackbar
      v-model="snackbar1"
      :timeout="3000"
      :color="snackbarColor"
      bottom
      center
    >
      {{ snackbarMessage }}
      <v-btn color="black" text @click="snackbar1 = false">Close</v-btn>
    </v-snackbar>
  </v-container>
</template>

<script>
import EventService from "../service/Eventservice";

export default {
  data() {
    return {
      billNo: null,
      quantity: "",
      selectedMedicine: "",
      medicineOptions: [],
      display: [],
      total: 0,
      gst: 0,
      netPrice: 0,
      showTable: false,
      previewDialog: false,
      text: "",
      snackbar: false,
      stock: [],
      amount: 0,
      billno: 101,
      headers: [
        { text: "Medicine Name", value: "medicine_name", align: "start" },
        { text: "Brand", value: "brand", align: "start" },
        { text: "Quantity", value: "quantity", align: "end" },
        { text: "Amount", value: "amount", align: "end" },
      ],
      snackbar1: false,
      snackbarMessage: "",
      responsemsg: "",
      snackbarColor: "",
      downloaddialog: false,
      downloadflag: 0,
    };
  },

  created() {
    this.stocks();

    EventService.SendBillEntryDrop()
      .then((response) => {
        // console.log(response);
        this.medicineOptions = response.data.medicine_names;
        // console.log(this.medicineOptions);
        // console.log(response.data.medicine_names);
      })
      .catch((error) => {
        console.log(error);
      });

    this.billno = Math.floor(Math.random() * (200 - 100 + 1)) + 100;
  },

  methods: {
    openPreviewDialog() {
      // this.calculateTotal();
      this.previewDialog = true;
    },
    closePreviewDialog() {
      this.previewDialog = false;
    },

    stocks() {
      EventService.sendStockview()
        .then((response) => {
          this.stock = response.data.stock_view_details;
          // console.log(this.stock);
        })
        .catch((error) => {
          console.log(error);
        });
    },

    addItem() {
      if (!this.selectedMedicine || !this.quantity) {
        this.snackbarMessage = "Please enter a Medicine and quantity";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      }

      let selectedItem = null;
      for (const item of this.stock) {
        // console.log(item.medicine_name);
        // console.log(item.medicine_quantity);
        if (item.medicine_name === this.selectedMedicine) {
          selectedItem = item;
          break;
        }
      }
      if (!selectedItem) {
        this.snackbarMessage = "Insufficinent Quantity";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      }

      const trimmedQuantity = this.quantity.trim();

      const quantityNum = parseInt(trimmedQuantity);
      if (isNaN(quantityNum) || quantityNum <= 0) {
        this.snackbarMessage = "Please Enter a valid Positive Quantity";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      }

      if (this.quantity.includes(".")) {
        this.snackbarMessage = "decimal values are not allowed";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      }

      if (selectedItem.medicine_quantity < quantityNum) {
        this.snackbarMessage = "Insufficinent Quantity ";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      }

      this.showTable = true;

      const existingItemIndex = this.display.findIndex(
        (item) => item.medicine_name === this.selectedMedicine
      );

      if (existingItemIndex !== -1) {
        this.display[existingItemIndex].quantity += quantityNum;
        this.display[existingItemIndex].amount =
          this.display[existingItemIndex].amount +
          selectedItem.unit_price * quantityNum;
      } else {
        this.amount = selectedItem.unit_price * quantityNum;
        let obj = {
          medicine_name: this.selectedMedicine,
          brand: selectedItem.brand,
          quantity: quantityNum,
          unit_price: selectedItem.unit_price,
          amount: this.amount,
          billno: this.billno,
        };
        this.display.push(obj);
      }

      let total = 0;

      for (let i = 0; i < this.display.length; i++) {
        let item = this.display[i];

        total += item.amount;
      }
      this.total = total;

      this.gst = Math.round(total * (18 / 100));

      this.netPrice = total + this.gst;

      this.selectedMedicine = "";
      this.quantity = "";
    },
    save() {
      if (this.display.length === 0) {
        this.snackbarMessage =
          "please select the medicine and quantity before saving ";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      } else {
        let jsondata = JSON.stringify(this.display);
        EventService.SendBillDetails(jsondata)
          .then((response) => {
            this.responsemsg = response.msg;
          })
          .catch((error) => {
            console.log(error);
          });

        const struct = {
          bill_no: this.billno,
          bill_amount: this.total,
          bill_gst: this.gst,
          net_price: this.netPrice,
          login_id: this.$store.state.go_login_id,
        };
        this.billNo = this.billNo + 1;
        this.stocks();

        let jsondata1 = JSON.stringify(struct);

        EventService.SendBillSave(jsondata1)
          .then((response) => {
            if (response.data.status == "S") {
              this.snackbarMessage = "bill saved successfully";
              this.snackbar1 = true;
              this.snackbarColor = "green";
              this.selectedMedicine = " ";
            this.quantity = " ";
            }else{
              this.snackbarMessage = "error";
              this.snackbar1 = true;
              this.snackbarColor = "red";

            }

            
          })
          .catch((error) => {
            console.log(error);
          });
        if (this.downloadflag == 0) {
          this.downloaddialog = true;
        }

        this.billno++;
        this.total = 0;
        this.netPrice = 0;
        this.gst = 0;
        this.showTable = false;
        this.display = [];
        this.billno = Math.floor(Math.random() * (200 - 100 + 1)) + 100;
      }
    },

    download() {
      const array = this.display;
      if (array.length === 0) {
        this.snackbarMessage = " No Data for Download ";
        this.snackbar1 = true;
        this.snackbarColor = "red";
        return;
      }
      const csvContent = this.convertToCSV(array);
      const blob = new Blob([csvContent], { type: "text/csv;charset=utf-8;" });
      const link = document.createElement("a");
      const url = URL.createObjectURL(blob);
      link.setAttribute("href", url);
      link.setAttribute("download", "data.csv");
      link.style.visibility = "hidden";
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    },
    convertToCSV(array) {
      const header = Object.keys(array[0]).join(",") + "\n";
      const rows = array.map((obj) => Object.values(obj).join(",")).join("\n");
      this.downloadflag = 1;
      return header + rows;
    },
  },

  // calculateTotal() {

  //   this.total = this.display.reduce(
  //     (acc, curr) => acc + curr.quantity * curr.amount,
  //     0
  //   );
  //   this.gst = (this.total * 18) / 100;
  //   this.netPrice = this.total + this.gst;
  // },

  props: ["login_history_id", "user_id", "login_id"],

  watch: {
    login_history_id(newValue, oldValue) {
      console.log("login_history_id changed:", newValue);
      console.log("login_history_id changed:", oldValue);
    },
  },
};
</script>

<style scoped></style>
