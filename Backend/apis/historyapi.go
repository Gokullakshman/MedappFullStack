package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/DBConnect"
	"net/http"
)
// HistoryDetails represents the structure of login history details.
type HistoryDetails struct {
	UserID     string `json:"user_id"`
	LoginTime  string `json:"login_time"`
	LogoutTime string `json:"logout_time"`
	LoginDate  string `json:"login_date"`
}
// HistoryResponse represents the structure of the response for fetching login history.
type HistoryResponse struct {
	HistoryDetailsArr []HistoryDetails `json:"history_details_arr"`
	Status            string           `json:"status"`
	ErrMsg            string           `json:"errmsg"`
}


/*
THIS METHOD IS USED TO GET THE HISTORY OF USERS 

RESPONSE
-------
ONSUCCESS{

{
  "history_details_arr": [
    {
      "user_id": "string",
      "login_time": "string",
      "logout_time": "string",
      "login_date": "string"
    }
    // Additional HistoryDetails objects can be included in the array
  ],
  "status": "string",
  "msg": "string",
  "errmsg": "string"
}

}




*/
// Handler function to fetch login history.
func History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	log.Println("History(+)")


	if r.Method == "GET" {
		// Reset the response variable for each request
		var phistoryRes HistoryResponse
        	// Call method to fetch login history data.
			phistoryRes.Status="S"
		phistoryRes, err := HistoryMethod()
		if err != nil {
			phistoryRes.ErrMsg = err.Error()+"HAPIMR"
			phistoryRes.Status = "E"
		} 
       		// Marshal the response data to JSON format.
		response, err := json.Marshal(phistoryRes)
		if err != nil {
			log.Println("JSON marshaling error:", err)

	        phistoryRes.ErrMsg=err.Error()+"HAPIMR"
			fmt.Fprintf(w,"ERROR"+err.Error())
		}
      	// Set response headers and write the response.
		// log.Println(string(response))
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(response))
	}
}



/*
THIS METHOD IS USED TO GET THE HISTORY FROM MEAPP_LOGIN_HISTORY APPEND INTO historyDetailsArr

ONSUCCESS
---------
IT WILL THE HISTORY ARRAY , STATUS WILL BE "S" AND ERRMSG WILL BE NIL IN THE phistoryRes Structure

ONERROR
------
IT WILL THE HISTORY ARRAY WILL BE NIL , STATUS WILL BE "E" AND ERRMSG WILL BE ERROR IN THE phistoryRes Structure
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"



*/
func HistoryMethod() (HistoryResponse, error) {
	log.Println("calling history method(+)")
	var phistoryRes HistoryResponse  

	db,err:= DBConnect.GORMDBConnect()

	if err!=nil{
		phistoryRes.ErrMsg=err.Error()+"HMDBC"
		phistoryRes.Status="E"
		return phistoryRes,nil
	}

	result:= db.Table("St832_Medapp_login_history mlh").
	Select("ml.user_id,mlh.login_time,coalesce(mlh.logout_time,0) as logout_time,mlh.login_date").
	Joins("join st832_medapp_login ml on ml.login_id = mlh.login_id").Find(&phistoryRes.HistoryDetailsArr) 

	if result.Error!=nil{
		log.Println("Query execution error:", err)
			phistoryRes.ErrMsg=err.Error()+"HMQEXEC1"
			phistoryRes.Status="E"
			return phistoryRes,result.Error

	}
	phistoryRes.Status="S"

	// db, err := DBConnect.LocalDBConnect()
	// if err != nil {
	// 	log.Println("Database connection error:", err)
	// 	phistoryRes.ErrMsg=err.Error()+"HMDBC"
	// 	phistoryRes.Status="E"
		
	// 	return phistoryRes,nil
	// }
	// defer db.Close()

	// 	corestring := `SELECT 
	// 	ml.user_id,
	// 	mlh.login_time,
	// 	nvl(mlh.logout_time,0),
	// 	mlh.login_date
	// FROM 
	// 	medapp_login_history mlh
	// LEFT JOIN 
	// 	medapp_login ml 
	// ON 
	// 	mlh.login_id = ml.login_id
	// WHERE 
	// 	ml.user_id IS NOT NULL`

	// rows, err := db.Query(corestring)
	// if err != nil {
	// 	log.Println("Query execution error:", err)
	// 	phistoryRes.ErrMsg=err.Error()+"HMQEXEC1"
	// 	phistoryRes.Status="E"
	// 	return phistoryRes,nil
		
	// }
	// defer rows.Close()

	// var historyDetailsArr []HistoryDetails

	// for rows.Next() {
	// 	var historydet HistoryDetails
	// 	err := rows.Scan(&historydet.UserID, &historydet.LoginTime, &historydet.LogoutTime, &historydet.LoginDate)
	// 	if err != nil {
	// 		log.Println("Row scanning error:", err)
	// 		phistoryRes.ErrMsg=err.Error()+"HMQSCAN1"
	// 	    phistoryRes.Status="E"
	// 	    return phistoryRes,nil
	// 	}
	// 	historyDetailsArr = append(historyDetailsArr, historydet)
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Println("Row iteration error:", err)
	// 	phistoryRes.ErrMsg=err.Error()
	// 	phistoryRes.Status="E"
	// 	return phistoryRes,nil
		
	// }

	return phistoryRes , nil
}
