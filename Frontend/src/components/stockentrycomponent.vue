
<template>
  <v-container max-width="200">
    <v-expansion-panels>
      <v-expansion-panel >
        <v-expansion-panel-header>
          <div class="font-weight-bold">Stock Entry</div>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-row>
            <v-col cols="12">
              <v-select
                v-model="selectedMedicine"
                :items="medicineNames"
                label="Medicine Name"
             
              ></v-select>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="selectedBrand"
                label="Brand"
                readonly
                dense
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="quantity"
                label="Quantity"
                type="number"
                dense
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model="unitPrice"
                label="Unit Price"
                type="number"
                dense
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="UpdateStock" color="primary">Update</v-btn>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn @click="openDialog" color="primary">Add</v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>

    <v-snackbar v-model="snackbar" :timeout="3000" :color="snackbarColor" bottom center>
        {{ snackbarMessage }}
        <v-btn color="black" text @click="snackbar = false">Close</v-btn>
      </v-snackbar>

    <v-dialog v-model="dialog" max-width="300">
      <v-card>
        <v-card-title>Add Stock Entry</v-card-title>
        <v-card-text>
          <v-text-field
            v-model="newMedicineName"
            label="Medicine Name"
            dense
          ></v-text-field>
          <v-text-field
            v-model="newMedicineBrand"
            label="Medicine Brand"
            dense
          ></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-btn color="blue darken-1" @click="dialog = false">Close</v-btn>
          <v-btn color="blue darken-1" @click="addstock">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>

  
</template>


  
  <script>
  import EventService from "../service/Eventservice";
  export default {
    data() {
      return {
        selectedMedicine: '',
        quantity: 0,
        unitPrice: 0,
        dialog: false,
        newMedicineName: '',
        newMedicineBrand: '',
        selectedBrand:'',
        medicineNames:[],
        snackbar: false,
        snackbarMessage: '',
        snackbarColor: 'brown',
        res:''

    
      };
    },
    created(){
      EventService.SendBillEntryDrop()
      .then((response) => {
        // console.log(response.data);
        this.medicineNames=response.data.medicine_names

       
      })
      .catch((error) => {
        console.log(error);
      });

    },
  //   updated(){
     
  //   EventService.SendAddUserDrop()
  //     .then((response) => {
  //       // console.log(response);
  //       this.apiroles = response.data.role;
  //       // console.log(this.apiroles);
  //     })
  //     .catch((error) => {
  //       console.log(error);
  //     });

  //     EventService.sendStockview()
  //     .then((response) => {
  //       // console.log(response.data);

  //       this.items = response.data.stock_view_details;
  //     })
  //     .catch((error) => {
  //       console.log(error);
  //     });
  // },
    

    watch: {
    selectedMedicine(newValue, oldValue) {
      if (newValue !== oldValue && newValue) {
        // console.log(newValue)
        
        this.fetchBrand(newValue);
      }
    },
  },
    // computed: {
    //   medicineNames() {
    //     return this.$store.state.medicinemaster.map(item => item.medicinename);
    //   },
    //   stock() {
    //     return this.$store.state.stock;
    //   },
    //   medicinemaster(){
    //     return this.$store.state.medicinemaster
    //   }
    // },
    methods: {

      medicines(){
        EventService.SendBillEntryDrop()
      .then((response) => {
        // console.log(response.data);
        this.medicineNames=response.data.medicine_names

       
      })
      .catch((error) => {
        console.log(error);
      });

      },
      
 
      fetchBrand(input){
        console.log(input)
        const struct={
          medicine_name:input
        }
        // console.log(struct)

        EventService.FetchBrand(struct)
      .then((response) => {
        if(response.data.status=="S"){
          this.selectedBrand=response.data.brand

        }
      })
      .catch((error) => {
        console.log(error);
      });



      },
      UpdateStock(){


        if (!this.selectedMedicine || !this.quantity || !this.unitPrice) {
          this.snackbarMessage = 'Please fill all fields';
          this.snackbar = true;
          this.snackbarColor = 'red'; 
          return;
        
       
      }
      
      if (this.quantity <= 0 || this.unitPrice <= 0) {
        this.snackbarMessage = 'Quantity and Unit Price cannot be negative';
        this.snackbar = true;
        this.snackbarColor = 'brown'; 
        return;
      }
        
        

        const struct={
          medicine_name:this.selectedMedicine,
          quantity:this.quantity,
          unit_price:this.unitPrice
        }
        EventService.SendUpdateStock(struct)
      .then((response) => {
        if(response.data.status=="S"){
          console.log(response) 
          this.res=response.data.err_msg
          

        }
      })
      .catch((error) => {
        console.log(error);
      });

      this.selectedMedicine='',
        this.selectedBrand='',
        this.unitPrice=0,
        this.quantity=0

        this.snackbarMessage = 'Stock Updated';
        this.snackbar = true;
        this.snackbarColor = 'green'; 

      },
      openDialog(){
        this.dialog=true
      },
      addstock(){

        if (!this.newMedicineName|| ! this.newMedicineBrand) {
          this.snackbarMessage = 'Please fill all fields';
          this.snackbar = true;
          this.snackbarColor = 'red'; 
          return;
        
       
      }
      
      if (!/^[a-zA-Z]+$/.test(this.newMedicineName) || !/^[a-zA-Z]+$/.test(this.newMedicineBrand)) {
  this.snackbarMessage = 'Please enter letters only';
  this.snackbar = true;
  this.snackbarColor = 'red';
  return;
}

        const struct={
          medicine_name:this.newMedicineName,
          brand:this.newMedicineBrand
        }
        EventService.SendAddStock(struct)
      .then((response) => {
        if(response.data.status=="S"){
          this.res=response.data.err_msg

        }else{

      this.snackbarMessage = 'medicine already exists';
      this.snackbar = true;
      this.snackbarColor = 'red '; 
      return
        }
      })
      .catch((error) => {
        console.log(error);
      });

      this.snackbarMessage = 'Stock Added';
      this.snackbar = true;
      this.snackbarColor = 'green'; 
      this.newMedicineName=" "
      this.newMedicineBrand=" "
      this.dialog=false
      this.medicines()


      }
      
      
    },

  

  };
  </script>
  