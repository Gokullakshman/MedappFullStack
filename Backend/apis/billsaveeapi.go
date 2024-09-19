package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	// "time"

	_ "github.com/go-sql-driver/mysql"
)

// Define the request structure for saving a bill.
type BillSaveeRequest struct {
	BillNo    int     `json:"bill_no"`
	Amount    int     `json:"bill_amount"`
	GST       float64 `json:"bill_gst"`
	Net_Price float64 `json:"net_price"`
	Login_id  int   `json:"login_id"`
}

// Define the response structure for saving a bill.
type BillSaveeResponse struct {
	Status string `json:"status"`
	ErrMsg string `json:"err_msg"`
}


type InsertSave struct{
	BILL_NO int `gorm:"column:bill_no"`
	BILL_DATE string `gorm:"column:bill_date"`
	BILL_AMOUNT int `gorm:"column:bill_amount"`
	BILL_GST float64 `gorm:"column:bill_gst"`
	NET_PRICE float64 `gorm:"column:net_price"`
    LOGIN_ID  int  `gorm:"column:login_id"`
	CREATED_BY string  `gorm:"column:created_by"`
	CREATED_DATE string `gorm:"column:created_date"`
	UPDATED_BY string `gorm:"column:updated_by"`
	UPDATED_DATE string `gorm:"column:updated_date"`

}
/*
THIS IS API USEDTO SAVE THE BILL IN DATABASE

REQUEST
-----
{
"bill_no":"122"
"bill_amount":"10"
"bill_gst":1
"net_price":"10"
"login_id":"1"

}
RESPONSE
-------
ONSUCCESS
{
"status":"S"
"err_msg":"nil"

}
ONERROR
{
"status":"E"
"err_msg":"ERRORMSG"

}

*/

func BillSavee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	if r.Method == "POST" {
		log.Println("billsave(+)")
			// Initialize request and response structs.
		var pbillsaveReq BillSaveeRequest
		var pbillsaveRes BillSaveeResponse
		pbillsaveRes.Status="S"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			pbillsaveRes.Status = "E"
			pbillsaveRes.ErrMsg = err.Error()+"BSAPIBR"
		}else{
				// Unmarshal JSON into request struct.
		err = json.Unmarshal(body, &pbillsaveReq)

		log.Println("billsave",pbillsaveReq)

		log.Println(pbillsaveReq)
		if err != nil {
			log.Println(err)
			pbillsaveRes.Status = "E"
			pbillsaveRes.ErrMsg = err.Error()+"BSAPIUM"
		}else{
			 	// Call method to save the bill.
				 pbillsaveRes, err := BillSaveeMethod(pbillsaveReq, pbillsaveRes)

				 if err != nil {
					 pbillsaveRes.ErrMsg = err.Error()+"BSAPIMR"
					 pbillsaveRes.Status = "E"
				 }

		}
			
		}
       
        
		// Marshal response data.
		datas, err := json.Marshal(pbillsaveRes)
		if err != nil {
			pbillsaveRes.ErrMsg = err.Error()+"BSAPIME"
			pbillsaveRes.Status = "E"
		} else {
			log.Println(string(datas))
			fmt.Fprint(w, string(datas))
		}
	}
}

/*
THIS METHOD IS USED TO INSERT THE DETAILS IN MEDAP_BILL_MASTER TABLE 
PARAMETERS:
---------
pbillsaveReq ,pbillsaveRes

ONSUCCESS
--------
IT WILL RETURN THE STATUS "S " AND ERRMSG WILL BE THE NULL IN THE pbillsaveRes Response STRUCTURE

ONERROR
--------
IT WILL RETURN THE STATUS "E" AND ERRMSG WILL BE THE ERROR IN THE pbillsaveRes Response STRUCTURE
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"

*/


// Method to save a bill.
func BillSaveeMethod(pbillsaveReq BillSaveeRequest, pbillsaveRes BillSaveeResponse) (BillSaveeResponse, error) {

	// db,err:= DBConnect.GORMDBConnect()

	// if err!=nil{
	// 	pbillsaveRes.ErrMsg="database connection error"+"BSMDBC"
	// 	pbillsaveRes.Status="E"
	// 	return pbillsaveRes, err
	// }

	// IS:= InsertSave{

	// 	BILL_NO: pbillsaveReq.BillNo,
	// 	BILL_DATE: time.Now().Format("2006-01-02"),
	// 	BILL_AMOUNT: pbillsaveReq.Amount,
	// 	BILL_GST: pbillsaveReq.GST,
	// 	NET_PRICE: pbillsaveReq.Net_Price,
	// 	LOGIN_ID: pbillsaveReq.Login_id,
	// 	CREATED_BY: "Gokul",
	// 	CREATED_DATE:time.Now().Format("2006-01-02"),
	// 	UPDATED_BY: "Gokul",
	// 	UPDATED_DATE: time.Now().Format("2006-01-02"),

	// }

	// res:= db.Table("St832_Medapp_bill_master").Create(&IS)
	// if res.Error!=nil{
	// 	pbillsaveRes.ErrMsg=res.Error.Error()+"BSMQEXEC1"
	// 	pbillsaveRes.Status="E"
	// 	return pbillsaveRes,res.Error
	// }

	// pbillsaveRes.Status="S"
	// return pbillsaveRes,nil

	

	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		pbillsaveRes.ErrMsg = "database connection error"+"BSMDBC"
		pbillsaveRes.Status = "E"
		return pbillsaveRes, err
	}
	defer db.Close()
    	// SQL query to insert bill details into the MEDAPP_BILL_MASTER table.
	corestring:=` INSERT INTO ST832_MEDAPP_BILL_MASTER
				(BILL_NO,BILL_DATE,BILL_AMOUNT,BILL_GST,NET_PRICE,LOGIN_ID,CREATED_BY,CREATED_DATE,UPDATED_BY,UPDATED_DATE)
	           VALUES(?,NOW(),?,?,?,?,?,NOW(),?,NOW());`
                 	// Execute the SQL query.
			   _, err1 := db.Exec(corestring, pbillsaveReq.BillNo,pbillsaveReq.Amount, pbillsaveReq.GST ,pbillsaveReq.Net_Price,pbillsaveReq.Login_id,
				"Gokul","Gokul")
				                                                                                      
				 
		if err1 != nil {
			pbillsaveRes.ErrMsg=err1.Error()+"BSMQEXEC1"
			pbillsaveRes.Status="E"
			return pbillsaveRes,err1

		}

      // Prepare successful response.
		pbillsaveRes.ErrMsg=""
		pbillsaveRes.Status="S"
		return pbillsaveRes,nil

}
