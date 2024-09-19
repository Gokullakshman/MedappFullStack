package main

import (
	"log"
	"medapp/apis"

	"net/http"
)

func main() {



	log.Println("server started")
	http.HandleFunc("/login", apis.Loginvalidation)
	http.HandleFunc("/loginhistory",apis.LoginHistory)
	http.HandleFunc("/stockview",apis.Stockview)
	http.HandleFunc("/logouthistory",apis.LogoutHistory)
	http.HandleFunc("/history",apis.History)
	http.HandleFunc("/unsoldamount",apis.UnSoldAmount)
	http.HandleFunc("/adduserdropdown",apis.AddUserDropDown)
	http.HandleFunc("/billentrydropdown",apis.BillEntryDropDown)


	http.HandleFunc("/fetchusers",apis.FetchUser)
	
	http.HandleFunc("/fetchsales",apis.FetchSales)
	http.HandleFunc("/fetchbrand",apis.FetchBrand)
	http.HandleFunc("/updatestock",apis.UpdateStock)
    http.HandleFunc("/addstock",apis.AddStock)
    http.HandleFunc("/billdetails",apis.BillDetails)
	http.HandleFunc("/billsavee",apis.BillSavee)
    http.HandleFunc("/today",apis.BillerTodaySales)
	http.HandleFunc("/monthsales",apis.FetchSaless)
	
	http.HandleFunc("/weeklysales",apis.FetchWeekSaless)

	
	// http.HandleFunc("/stock",apis.Stock)
	// http.HandleFunc("/updatestockdropdown",apis.UpdateStockDropDown)
	
	http.ListenAndServe(":2010",nil)
	
}
