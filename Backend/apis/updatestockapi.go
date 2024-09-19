package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	"strconv"
)

// UpdateStockRequest represents the structure of the request for updating stock
type UpdateStockRequest struct {
	MedicineName string `json:"medicine_name"`
	Quantity     string `json:"quantity"`
	Unit_Price   string `json:"unit_price"`
}

// UpdateStockResponse represents the structure of the response for updating stock
type UpdateStockResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	ErrMsg string `json:"errmsg"`
}

type UpdateStock1 struct{
	Quantity int `gorm:"column:quantity"`
	Unit_Price string `gorm:"column:unit_price"`
}
/*
THIS IS API IS USED TO UPDATE THE STOCK

REQUEST
-------
{
  "medicine_name": "Aspirin",
  "quantity": "100",
  "unit_price": "0.50"
}

RESPONSE
--------
ONSUCCESS
-------
{
  "status": "S",
  "errmsg": ""
}
  ONERROR
-------
{
  "status": "E",
  "errmsg": "ERROR"
}
*/

// UpdateStock handles the HTTP request for updating stock
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")


	// Handle PUT request
	if r.Method == "POST" {
		log.Println("UpdateStock(+)")

		var pupdateReq UpdateStockRequest
		var pupdateRes UpdateStockResponse

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			pupdateRes.ErrMsg = err.Error() + "USAPIBR"
			pupdateRes.Status = "E"
		} else {
			err := json.Unmarshal(body, &pupdateReq)

			if err != nil {
				pupdateRes.ErrMsg = err.Error() + "USAPIUM"
				pupdateRes.Status = "E"

			}
			log.Println("obj", pupdateReq)

			log.Println(pupdateReq)

			pupdateRes, err = UpdateStockMethod(pupdateReq, pupdateRes)

			if err != nil {
				pupdateRes.ErrMsg = err.Error() + "USAPIMR"
				pupdateRes.Status = "E"
			} 
		}

			datas, err := json.Marshal(pupdateRes)
			if err != nil {
				pupdateRes.ErrMsg = err.Error() + "USAPIME"
				pupdateRes.Status = "E"

			} 
				fmt.Fprintf(w, string(datas))
				log.Println(string(datas))

			

		
	}
}


/*
THIS METHOD IS USED TO UPDATE THE STOCK FIRSTLY IT WILL GET MEDICINE_MASTER_ID FROM MEDAPP_MEDICINE_MASTER TABLE WITH 
THIS MEDICINE_MASTER_ID UPDATE THE STOCK IN MEDAPP_STOCK TABLE 

PARAMETERS
---------
pupdateReq,pupdateRes

ONSUCCESS
--------
IT WILL RETURN THE STATUS "S" AND ERRMSG WILL BE "NIL" IN THE pupdateRes RESPONSE STRUCTURE

ONERROR
--------
IT WILL RETURN THE STATUS "E" AND ERRMSG WILL BE "ERROR" IN THE pupdateRes RESPONSE STRUCTURE
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"


*/
// UpdateStockMethod updates the stock in the database
func UpdateStockMethod(pupdateReq UpdateStockRequest, pupdateRes UpdateStockResponse) (UpdateStockResponse, error) {
	// Connect to the local database

	db,err:= DBConnect.GORMDBConnect() 
	if err!=nil{
		pupdateRes.ErrMsg = err.Error() + "USMDBC"
		pupdateRes.Status = "E"
		return pupdateRes, err

	}
	var med_id int 

	result:=db.Table("St832_Medapp_Medicine_Master").Select("Medicine_Master_Id").Where("Medicine_name=?",pupdateReq.MedicineName).Find(&med_id)

	if result.Error!=nil{
		pupdateRes.ErrMsg =result.Error.Error() + "USMQEXEC1"
		pupdateRes.Status = "E"
		return pupdateRes ,result.Error
	}
	var defaultquant int 
	

	result2:=db.Table("St832_Medapp_Stock").Select("Quantity").Where("Medicine_Master_id =?",med_id).Find(&defaultquant)

	if result2.Error!=nil{
		pupdateRes.ErrMsg = result2.Error.Error() + "USMQEXEC2"
		pupdateRes.Status = "E"
		return pupdateRes ,result2.Error
	}

	intquantity,err:=strconv.Atoi(pupdateReq.Quantity)
	new:=defaultquant+intquantity


	result3:= db.Table("St832_Medapp_Stock").Where("medicine_master_id=?",med_id).Updates(
		UpdateStock1{
			Quantity: new,
			Unit_Price: pupdateReq.Unit_Price,
		})

		if result3.RowsAffected==0{
			log.Println("no updated stock")
		}else{
			log.Println("updated stock")
		}

		if result3.Error!=nil{
			pupdateRes.ErrMsg = result3.Error.Error() + "USMQEXEC3"
		pupdateRes.Status = "E"
		return pupdateRes ,result3.Error

		}
	// db, err := DBConnect.LocalDBConnect()

	// log.Println(pupdateReq)

	// var med_id int

	// if err != nil {
	// 	pupdateRes.ErrMsg = err.Error() + "USMDBC"
	// 	pupdateRes.Status = "E"
	// 	return pupdateRes, err

	// } else {
	// 	defer db.Close()
	// 	// Query to retrieve medicine ID
	// 	corestring := `SELECT MEDICINE_MASTER_ID
	// FROM ST832_MEDAPP_MEDICINE_MASTER
	// WHERE MEDICINE_NAME = ?`

	// 	rows, err := db.Query(corestring, &pupdateReq.MedicineName)
	// 	if err != nil {

	// 		pupdateRes.ErrMsg = err.Error() + "USMQEXEC1"
	// 		pupdateRes.Status = "E"
	// 		return pupdateRes, err

	// 	}

	// 	for rows.Next() {
	// 		err := rows.Scan(&med_id)
	// 		log.Println(med_id)

	// 		if err != nil {
	// 			pupdateRes.ErrMsg = err.Error() + "USMSCAN1"
	// 			pupdateRes.Status = "E"
	// 			return pupdateRes, err

	// 		}
	// 	}
	// 	// Query to update stock
	// 	corestring1 := `UPDATE ST832_MEDAPP_STOCK SET QUANTITY = ?,UNIT_PRICE = ? WHERE MEDICINE_MASTER_ID = ?`
	// 	log.Println(pupdateReq.Quantity)
	// 	log.Println(pupdateReq.Unit_Price)
	// 	_, err1 := db.Exec(corestring1, &pupdateReq.Quantity, &pupdateReq.Unit_Price, &med_id)
	// 	if err1 != nil {
	// 		pupdateRes.ErrMsg = err1.Error() + "USMQEXEC2"
	// 		pupdateRes.Status = "E"
	// 		return pupdateRes, err

	// 	}

	// }
	pupdateRes.Msg = "updated successfully"
	pupdateRes.Status = "S"
	return pupdateRes, nil
}
