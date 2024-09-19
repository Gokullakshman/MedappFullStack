import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    login:[
      {userid:"B123",password:"123",role:"Biller"  },
  {userid:"M123",password:"456",role:"Manager"},
{userid:"SA123",password:"789",role:"Systemadmin"},
{userid:"I123",password:"012",role:"Inventry"}],
currentuser:null,
loginhistory:[],
go_login_id:0,
go_user_id:0,
totalamount:0,
stock: [
  { medicinename: "Paracetamol", quantity: 10, amount: 10 },
  { medicinename: "Fentanyl", quantity: 50, amount: 5 },
  { medicinename: "Folicacid", quantity: 75, amount: 8 }
],
medicinemaster:[{medicinename:"Paracetamol",Brand:"p"},
{medicinename:"Fentanyl",Brand:"F"},
{medicinename:"Folicacid",Brand:"Fo"}
],
billmaster:[
  {billNo:11,  // Generate a random bill number
  date:"26/11/2023",
     total: 100,
  gst: 26,
  netPrice: 126,
  userid:"sac"},

  { billno:15,date:"2023-11-21",total:260,gst:15,netPrice:275,userid:"cas"}
],

// medicinename: item.medicinename,
// quantity: item.quantity,
// amount: item.amt ,
// unitprice: item.net,
// billno:
billdetails:[{billno:12,quantity:2,amount:40,unitprice:45,medicinename:"Paracetamol"}

]

},

  mutations:{
    setcurrentuser(state,user){
      state.currentuser=user;
    },
    setloginhistory(state,HistoryEntry){
      state.loginhistory.push(HistoryEntry)
    },
    updatelogin(state,newuser){
      console.log("HELLO")
        state.login.push(newuser)
    },
    updatestock(state,obj){
        state.stock[obj.i].quantity=state.stock[obj.i].quantity+parseInt(obj.q);
        state.stock[obj.i].amount=obj.p;
        console.log(obj.q);
        console.log(obj.i);
        console.log(obj.p);
    },
    addnewmaster(state,{n}){
        state.medicinemaster.push(n)
        console.log(state.medicinemaster[3].medicinename)
    },
    addmedicine(state,n){
     
      state.medicinemaster.push(n);
    },
    newupdatestock(state,n){
      console.log("hello");
      state.stock.push(n)
    },
    billmaster(state,{bill}){
      console.log("HELLO")
       state.billmaster.push(bill)
       console.log(state.billmaster[0].netPrice)
       console.log(state.billmaster[2].date)
       console.log(state.billmaster[2].userid)

      
    },
    billdetails(state, { bill }) {
      // Check if the bill with the same medicinename already exists in the array
      const existingBillIndex = state.billdetails.findIndex(
        existingBill => existingBill.medicinename === bill.medicinename
      );
    
      if (existingBillIndex === -1) {
        // If the bill doesn't exist, push it to the array
        state.billdetails.push(bill);
      } else {
        // If the bill exists, update its quantity and amount
        state.billdetails[existingBillIndex].quantity += bill.quantity;
        state.billdetails[existingBillIndex].amount += bill.amount;
      }
    },
  
    updateStock(state, updatedStock) {
      // Update the stock in the store
      state.stock = updatedStock;
    }


   
  },

  actions: {
    updateStock({ commit, state }, medicine) {
      // Find the medicine in the stock array
      const index = state.stock.findIndex(item => item.medicinename === medicine.medicinename);

      
      if (state.stock[index].quantity === 0 || state.stock[index].quantity<0) {
        return; // Stop the action if the quantity is already zero
      }
  
      if (index !== -1) {
        // Update the quantity of the medicine in the stock
        const updatedStock = [...state.stock];
        updatedStock[index].quantity -= medicine.quantity;
        
        // Commit mutation to update the stock in the store
        commit('updateStock', updatedStock);
      }
    }
  },
  modules: {},
});