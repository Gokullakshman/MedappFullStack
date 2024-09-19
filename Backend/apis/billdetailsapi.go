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

// BillDetailsRequest represents the structure for the request to the BillDetails
type BillDetailsRequest struct {
	Amount        int    `json:"amount"`
	BillNo        int    `json:"billno"`
	Brand         string `json:"brand"`
	Medicine_Name string `json:"medicine_name"`
	Quantity      int    `json:"quantity"`
	Unit_Price    int    `json:"unit_price"`
}

var arr []BillDetailsRequest

type BillDetailsResponse struct {
	Status string `json:"status"`
	ErrMsg string `json:"err_msg"`
}

// `INSERT INTO MEDAPP_BILL_DETAILS (BILL_NO,MEDICINE_MASTER_ID, QUANTITY, UNIT_PRICE, AMOUNT,
// 	// 		CREATED_BY, CREATED_DATE, UPDATED_BY, UPDATED_DATE)
// 	// 	 VALUES (?, ?, ?, ?, ?,?, NOW(), ?, NOW())`

type InsertDet struct {
	Bill_No            int `gorm:"column:"bill_no"`
	MEDICINE_MASTER_ID int `gorm:"column:"medicine_master_id"`
	QUANTITY           int `gorm:"column:"quantity"`
	Unit_Price         int `gorm:"column:"unit_price"`
	Amount             int `gorm:"column:amount"` 
	Created_By         string `gorm:"column:created_by"`
	Created_Date       string `gorm:"column:created_date"`
	Updated_By         string `gorm:"column:updated_by"`
	Updated_Date      string `gorm:"column:updated_date"`
}

type UpdateDet struct {
	Quantity int `gorm:"column:quantity"`
	Updated_By string `gorm:"column:updated_by"`
	Updated_Date string `gorm:'column:updated_date"`
}

/*THIS API IS USED TO ADD THE BILL DETAILS IN DATABASE

REQUEST
-------
[
		{
			"amount": 150,
			"billno": 123,
			"brand": "BrandA",
			"medicine_name": "MedicineA",
			"quantity": 2,
			"unit_price": 75
		},
		{
			"amount": 200,
			"billno": 124,
			"brand": "BrandB",
			"medicine_name": "MedicineB",
			"quantity": 4,
			"unit_price": 50
		}
	]`

	RESPONSE
	-------
	ONSUCCESS
	--------
	{
	"status":"S"
	"err_msg":"nil"
	}
	ONERROR
	-------
	{
	"status":"E"
	"err_msg":"error msg"
	}

*/

// BillDetails is an HTTP handler for adding bill details.
func BillDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	if r.Method == "POST" {
		log.Println("BillDetails(+)")

		var pbilldetailsRes BillDetailsResponse
		var pbilldetailsReq BillDetailsRequest

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			pbilldetailsRes.Status = "E"
			pbilldetailsRes.ErrMsg = err.Error() + "BDAPIBR"
		} else {
			err = json.Unmarshal(body, &arr) 
			log.Println(arr)
			if err != nil {
				log.Println(err)
				pbilldetailsRes.Status = "E"
				pbilldetailsRes.ErrMsg = err.Error() + "BDAPIUM"
			} else {
				pbilldetailsRes, err = BillDetailsMethod(arr, pbilldetailsReq, pbilldetailsRes)

				if err != nil {
					log.Println(err)
					pbilldetailsRes.ErrMsg = err.Error() + "BDAPIMR"
					pbilldetailsRes.Status = "E"
				}

			}

		}

		datas, err := json.Marshal(pbilldetailsRes)
		if err != nil {
			log.Println(err)
			pbilldetailsRes.ErrMsg = "marshal error"
			pbilldetailsRes.Status = "E"
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			log.Println("suucess",string(datas))
			fmt.Fprintf(w, string(datas))

		}

		// response, _ := json.Marshal(pbilldetailsRes)
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(response)
	}
}

/*
PURPOSE:FIRST THIS METHOD IS USED TO GET THE MEDICINE_MASTER_ID FROM MEDAPP_MEDICINE_MASTERTABLE WITH THIS ID IT WILL INSERT THE
BILLN0,MEDICINE_MASTER_ID,QUANTITY,UNIT_PRICE,AMOUNT AND AFTER INSERTING IT WILL THE STOCKS IN MEDAPP_STOCK TABLE


PARAMETERS:arr ,pbilldetailsRes

ONSUCCESS
------
 IT WILL RETURN THE STATUS "S " AND ERRMSG WILL BE "NIL IN THE pbilldetailsRes RESPONSE STRUCTURE"

ONERROR
-----
IT WILL RETURN THE STATUS "E"AND ERRMSG WILL BE ERROR IN THE pbilldetailsRes RESPONSE STRUCTURE

AUTHORIZARION:"GOKUL"
DATE:"25-05-24"



*/
// BillDetailsMethod adds bill details to the database.
func BillDetailsMethod(arr []BillDetailsRequest, pbilldetailsReq BillDetailsRequest, pbilldetailsRes BillDetailsResponse) (BillDetailsResponse, error) {
	// db, err := DBConnect.GORMDBConnect()
	// if err != nil {
	// 	log.Println(err)
	// 	pbilldetailsRes.ErrMsg = "database connection error"
	// 	pbilldetailsRes.Status = "E"
	// 	return pbilldetailsRes, err
	// }
	// var med_id int
	// for _, medicine := range arr {
	// 	result := db.Table("st832_medapp_medicine_master").Select("medicine_master_id").Where("medicine_name=?", medicine.Medicine_Name).Find(&med_id)

	// 	if result.Error != nil {
	// 		pbilldetailsRes.Status = "E"
	// 		pbilldetailsRes.ErrMsg = result.Error.Error() + "BDMQEXEC1"
	// 		return pbilldetailsRes, result.Error

	// 	}

	// 	IB := InsertDet{
	// 		Bill_No:            medicine.BillNo,
	// 		MEDICINE_MASTER_ID: med_id,
	// 		QUANTITY:           medicine.Quantity,
	// 		Amount:             medicine.Amount,
	// 		Unit_Price:         medicine.Unit_Price,
	// 		Created_By: "GOKUL",
	// 		Created_Date: time.Now().Format("2006-01-02"),
	// 		Updated_By: "Gokul",
	// 		Updated_Date:time.Now().Format("2006-01-02") ,
	// 	}
	// 	result2 := db.Table("medapp_bill_details").Create(&IB)

	// 	if result2.Error != nil {
	// 		pbilldetailsRes.Status = "E"
	// 		pbilldetailsRes.ErrMsg = result2.Error.Error() + "BDMQEXEC2"
	// 		return pbilldetailsRes, result2.Error

	// 	}

	// 	var quant int 
	// 	var newquant int 

	// 	result3 := db.Table("st832_medapp_stock").Select("quantity").Where("medicine_master_id=?", med_id).Find(&quant)

	// 	if result3.Error != nil {
	// 		pbilldetailsRes.Status = "E"
	// 		pbilldetailsRes.ErrMsg = result3.Error.Error() + "BDMQEXEC2"
	// 		return pbilldetailsRes, result3.Error

	// 	} 
	// 	newquant = quant - medicine.Quantity 

	// 	log.Println("new quan",newquant)
	// 	result4 := db.Table("st832_medapp_stock").Where("medicine_master_id=?", med_id).Updates(UpdateDet{
	// 		Quantity: newquant,
	// 		Updated_By: "Gokul",
	// 		Updated_Date: time.Now().Format("2006-01-02"),

	// 	}) 

	// 	fmt.Println("Rows affected:", result.RowsAffected)
    //    if result.RowsAffected > 0 {
    //     fmt.Println("Update successful!")
    //    } else {
    //     fmt.Println("Update did not affect any rows.")
    //      }


	//    if result4.Error != nil {
	// 	pbilldetailsRes.Status = "E"
	// 	pbilldetailsRes.ErrMsg = result4.Error.Error() + "BDMQEXEC2"
	// 	return pbilldetailsRes, result4.Error

	// }

		



	// }




	// lerr := db.Table("medapp_login_history").Where("login_history_id = ?", historyId).Updates(MedappLoginHistory{
	//         LogoutDate:  logoutDate,
	//         LogoutTime:  logoutTime,
	//         UpdatedBy:   "elizabeth",
	//         UpdatedDate: currentTime,
	//     }).Error
 db,err:= DBConnect.LocalDBConnect()
  if err != nil {
		log.Println(err)
		pbilldetailsRes.ErrMsg = "database connection error"
		pbilldetailsRes.Status = "E"
		return pbilldetailsRes, err
	}


	defer db.Close()

	for _, medicine := range arr {
		var med_id int
		// Query to retrieve the MEDICINE_MASTER_ID from the MEDAPP_MEDICINE_MASTER table based on MEDICINE_NAME.
		corestring := `SELECT MEDICINE_MASTER_ID FROM ST832_MEDAPP_MEDICINE_MASTER WHERE MEDICINE_NAME = ?`
		err := db.QueryRow(corestring, medicine.Medicine_Name).Scan(&med_id)
		log.Println(medicine.Medicine_Name)
		log.Println(med_id)
		if err != nil {
			log.Println(err)
			log.Println("errrrrr1")
			pbilldetailsRes.Status = "E"
			pbilldetailsRes.ErrMsg = err.Error()+"BDMQEXEC1"
			return pbilldetailsRes, err
		}
		// Insertion query to add a new record into the MEDAPP_BILL_DETAILS table.
		corestring1 := `INSERT INTO MEDAPP_BILL_DETAILS (BILL_NO,MEDICINE_MASTER_ID, QUANTITY, UNIT_PRICE, AMOUNT,
			CREATED_BY, CREATED_DATE, UPDATED_BY, UPDATED_DATE)
		 VALUES (?, ?, ?, ?, ?,?, NOW(), ?, NOW())`
		_, err1 := db.Exec(corestring1, medicine.BillNo, med_id, medicine.Quantity, medicine.Unit_Price, medicine.Amount, "Gokul", "gokul")
		if err1 != nil {
			log.Println(err)
			log.Println("errrrrr2")
			pbilldetailsRes.Status = "E"
			pbilldetailsRes.ErrMsg = err1.Error()+"BDMQEXEC2"
			return pbilldetailsRes, err1
		}
		// Update query to decrement the QUANTITY in the MEDAPP_STOCK table based on MEDICINE_MASTER_ID.
		corestring2 := `UPDATE ST832_MEDAPP_STOCK SET QUANTITY=QUANTITY-? WHERE MEDICINE_MASTER_ID=?`
		_, err2 := db.Exec(corestring2, medicine.Quantity, med_id)

		if err2 != nil {
			log.Println(err)
			pbilldetailsRes.Status = "E"
			pbilldetailsRes.ErrMsg = err2.Error()+"BDMQEXEC2"
			return pbilldetailsRes, err2
		}

	}
	pbilldetailsRes.Status = "S"

	return pbilldetailsRes, nil
}
