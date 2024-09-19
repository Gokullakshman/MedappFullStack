package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	"time"
	//  "gorm.io/gorm"
)

// AddStockRequest represents the request structure for adding a stock.
type AddStockRequest struct {
	MedicineName string `json:"medicine_name"`
	Brand        string `json:"brand"`
}

// AddStockResponse represents the response  structure
type AddStockResponse struct {
	Status string `json:"status"`
	ErrMsg string `json:"errmsg"`
}
type Insert struct{
    Medicine_Master_id           int      `gorm:"primaryKey;autoIncrement"`
	MedicineName string `gorm:"medicine_name"`
	Brand string `gorm:"brand"`
	CreatedBy string `gorm:"created_by"`

	CreatedDate string  `gorm:"column:CREATED_DATE"`
	UpdatedBy   string    `gorm:"column:UPDATED_BY"`
	UpdatedDate string `gorm:"column:UPDATED_DATE"`

}

type Insert1 struct{
	Medicine_Master_id int `gorm:"column:medicine_master_id"`
	Created_By string `gorm:"column:created_by"`
	Created_Date  string`gorm:"column:CREATED_DATE"`
}
/*THIS API IS USED TO ADD A NEW STOCK 
REQUEST
--------
{
  "medicine_name": "Ibuprofen",
  "brand": "HealthPlus"
}
  RESPONSE
  -------
  ONSUCCESS
  -------
  {
  "status": "S",
  "errmsg": "NIL"
}
    ONERROR
  -------
  {
  "status": "E",
  "errmsg": "ERROR"
}

*/
// AddStock is an HTTP handler for adding stock.
func AddStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")



	if r.Method == "POST" {
		log.Println("AddStock(+)")

		var paddReq AddStockRequest
		
		var paddRes AddStockResponse
		paddRes.Status="S"

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			paddRes = AddStockResponse{
				ErrMsg: err.Error()+"ASAPIRB",
				Status: "E",
			}

			return
		}

		err = json.Unmarshal(body, &paddReq)
		if err != nil {
			paddRes = AddStockResponse{
				ErrMsg: err.Error()+"ASAPIUM",
				Status: "E",
			}

			return
		}

		log.Println("Request Object:", paddReq)

		paddRes, err = AddStockMethod(paddReq)

		if err != nil {
			paddRes.Status = "E"
			paddRes.ErrMsg = err.Error()+"ASAPIMR"

		}

		datas, err := json.Marshal(paddRes)

		if err != nil {
			paddRes.ErrMsg = err.Error()+"ASAPIME"
			paddRes.Status = "E"

		}
		log.Println(string(datas))
		fmt.Fprintf(w, string(datas))

	}
}
/*
THIS METHOD IS USED TO ADD NEW STOCK IF NOT EXISTS FIRSTLY CHECK IF THE MEDICINE_NAME IS ALREADY PRESENT IN MEDAPP_MEDICINE_MASTER TABLE 
IF NOT EXISTS THEN INSERT THE MEDICINENAME,MEDICINEBRAND IN MEDICINE_MASTER_TABLE TABLE AND GET MEDICINE_MASTER_ID FROM  MEDAPP_MEDICINE_MASTER TABLE
USING MEDICINENAME AND MEDICINEBRAND AND INSERT THE MEDICINE_MASTER_ID IN MEDAPP_STOCK TABLE

PARAMETERS
--------
paddReq

ONSUCCESS
--------
IT WILL RETURN THE STATUS WILL BE "S" AND ERR_MSG WILL BE NIL IN THE paddRes RESPONSE STRUCTURE

ONERROR
-------
IT WILL RETURN THE STATUS WILL BE "E" AND ERR_MSG WILL BE ERROR IN THE paddRes RESPONSE STRUCTURE

AUTHORIZARION:"GOKUL"
DATE:"25-05-24"
*/
// AddStockMethod adds stock to the database.
func AddStockMethod(paddReq AddStockRequest) (AddStockResponse, error) {
	var paddRes AddStockResponse
	var count int64

	db,err:= DBConnect.GORMDBConnect()
	if err != nil {
			paddRes.ErrMsg = err.Error()+"ASMDBC"
			paddRes.Status = "E"
			return paddRes, err
		}
		result := db.Table("st832_medapp_medicine_master").
        Where("medicine_name = ? AND brand = ?", paddReq.MedicineName, paddReq.Brand).
        Count(&count)

    if result.Error != nil {
        return paddRes, result.Error
    }
	log.Println("count",count)
	if count>0{
		paddRes.Status="E"
		paddRes.ErrMsg="already exists"
		return paddRes,err
	}

	hg:=Insert{
		MedicineName: paddReq.MedicineName,
		Brand: paddReq.Brand,
		CreatedBy: "Gokul",
		CreatedDate:time.Now().Format("2006-01-02") ,
		UpdatedBy: "Gokul",
		UpdatedDate: time.Now().Format("2006-01-02") ,

	}
	result1:= db.Table("St832_Medapp_Medicine_Master").Create(&hg) 
	if(result1.Error!=nil){
         return paddRes,result1.Error
	}


	insertid:=hg.Medicine_Master_id
	log.Println("insert",insertid)

	hk:=Insert1{
		Medicine_Master_id: insertid,
		Created_By: "Gokul",
		Created_Date: time.Now().Format("2006-01-02"),
	}

	result2:= db.Table("St832_Medapp_Stock").Create(&hk)
	if(result2.Error!=nil){
		paddRes.Status="E"
		paddRes.ErrMsg=result2.Error.Error()
		return paddRes,result.Error

	}
	log.Println("last inserted",result2.RowsAffected)

	paddRes.Status="S"
    return paddRes,nil







		

	// db, err := DBConnect.LocalDBConnect()
	// if err != nil {
	// 	paddRes.ErrMsg = err.Error()+"ASMDBC"
	// 	paddRes.Status = "E"
	// 	return paddRes, err
	// }
	// defer db.Close()

	// // Check if the record already exists
	// existsQuery := `
    //     SELECT COUNT(1) FROM MEDAPP_MEDICINE_MASTER 
    //     WHERE MEDICINE_NAME = ? AND BRAND = ?
    // `
	// var count int
	// err = db.QueryRow(existsQuery, paddReq.MedicineName, paddReq.Brand).Scan(&count)
	// if err != nil {
	// 	paddRes.ErrMsg = err.Error()+"ASMQE"
	// 	paddRes.Status = "E"
	// 	return paddRes, err
	// }

	// if count > 0 {
	// 	paddRes.Status = "E"
	// 	return paddRes, nil
	// }

	// // Insert the record if it does not exist
	// insertQuery := `
    //     INSERT INTO MEDAPP_MEDICINE_MASTER (MEDICINE_NAME, BRAND, CREATED_BY, CREATED_DATE, UPDATED_BY, UPDATED_DATE)
    //     VALUES (?, ?, ?, NOW(), ?, NOW())
    // `
	// createdBy := "Gokul"
	// _, err = db.Exec(insertQuery, paddReq.MedicineName, paddReq.Brand, createdBy, createdBy)
	// if err != nil {
	// 	paddRes.ErrMsg = err.Error()+"ASMQEXEUTE1"
	// 	paddRes.Status = "E"
	// 	return paddRes, err
	// }

	// // Retrieve the ID of the inserted record
	// QUERY := `SELECT MEDICINE_MASTER_ID FROM MEDAPP_MEDICINE_MASTER 
	// WHERE MEDICINE_NAME = ? AND BRAND = ?`
	// var id int

	// err = db.QueryRow(QUERY, paddReq.MedicineName, paddReq.Brand).Scan(&id)

	// if err != nil {
	// 	paddRes.ErrMsg = err.Error()+"ASMQEXEUTE2"
	// 	paddRes.Status = "E"
	// 	return paddRes, err
	// }
	// // Insert stock entry
	// insertQuery1 := `
    //     INSERT INTO MEDAPP_STOCK (MEDICINE_MASTER_ID, CREATED_BY, CREATED_DATE)
    //     VALUES (?, ?, NOW())
    // `
	// createdBy1 := "Gokul"
	// _, err = db.Exec(insertQuery1, id, createdBy1)
	// if err != nil {
	// 	paddRes.ErrMsg = err.Error()+"ASMQEXEUTE3"
	// 	paddRes.Status = "E"
	// 	return paddRes, err
	// }



}
