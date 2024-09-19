import axios from "axios"


const api = axios.create({
    baseURL:`http://localhost:2010`,
    withCredentials:false,
    headers:{
        Accept:"application/json",
       "content-Type":"application/json",
    }
})

export default{

    sendLogin(input){
        console.log(input)
        return api.post("/login",input)
    },
        sendLoginHistory(input){

            // console.log(input)
            return api.post("/loginhistory",input)
        },
    sendStockview(){
        // console.log("calling stock view")
        return api.get("/stockview")
    },
    sendHistory(){
        // console.log("calling history")
        return api.get("/history")
    },
    sendLogoutHistory(input){
        // console.log(" calling logout history")
        return api.post("/logouthistory",input)
    },
    sendUnsold(){
        // console.log("calling unsold ")
        return api.get("/unsoldamount")
    },

    SendAddUserDrop(){
        // console.log("calling adduser drop")
        return api.get("/adduserdropdown")
    },
    SendBillEntryDrop(){
        // console.log("calling adduser drop")
        return api.get("/billentrydropdown")
    },
    SendMedQuantity(input){
        // console.log("calling checking quantity")
        return api.put("/medquantity",input)
    },
 
    FetchUsers(input){
        console.log("calling fetch users",input)
        return api.put("/fetchusers",input)
    },
    FetchSales(input){
        // console.log("calling fetch sales")
        return api.post("/fetchsales",input)
    },
    FetchBrand(input){
        
        // console.log("calling fetch brand")
        return api.post("/fetchbrand",input)

    },
    SendUpdateStock(input){
        // console.log("calling update stock",input)
        return api.post("/updatestock",input)

    },
    SendAddStock(input){
        // console.log("calling add stock",input)
        return api.post("/addstock",input)

    },
    Stock(){
        // console.log("calling stock")
        return api.get("/stock")
    },
    SendBillDetails(input){
        // console.log("calling bill details",input)
        return api.post("/billdetails",input)

    },
    SendBillSave(input){
        // console.log("calling billsave",input)
        return api.post("/billsavee",input)
    },
    sendToday(input){
        // console.log("calling sendtoday")
        return api.post("/today",input)

    },
    FetchMonthSales(){
        return api.get("/monthsales")
    },
    FetchWeeklySales(){
        return api.get("/weeklysales")

    }
    // sendUpdateStockDropDown(){
    //     console.log("calling updatestock dropdown")
    //     return api.put("/updatestockdropdown")

    // }
    
}