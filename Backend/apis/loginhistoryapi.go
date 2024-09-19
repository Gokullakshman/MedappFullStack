package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)
// LoginHistoryRequest represents the structure of the request for login history.
type LoginHistoryRequest struct {
	Login_id int `json:"login_id"`
}
// LoginHistoryResponse represents the structure of the response for login history.
type LoginHistoryResponse struct {
	Login_History_Id int    `json:"login_history_id"`
	Err_msg          string `json:"err_msg`
	Status           string `json:"status"`
}

type Insert2 struct{
	Login_History_id int  `gorm:"primaryKey;autoIncrement"`
	Login_id int `gorm:"login_id"`
	Login_Date string `gorm:"login_date"`
	Login_Time  string  `gorm:"login_time"`
	Created_By string `gorm:"created_by"`
	Created_Date string `gorm:"created_date"`

}
/*THIS API IS USED TO INSERT LOGIN HISTORY FOR USER

REQUEST
-------
{
  "login_id": 12345
}
RESPONSE
--------
ONSUCCESS{
"login_history_id": "123",
  "err_msg": "",
  "status": "S"

}
  ONERRor{
"login_history_id": "nil",
  "err_msg": "error",
  "status": "E"

}



*/


var ploginhistoryReq LoginHistoryRequest

var ploginhistoryRes LoginHistoryResponse
// LoginHistory handles the request for login history.
func LoginHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			ploginhistoryRes.Err_msg = err.Error()+"LHAPIBR"
			ploginhistoryRes.Status="E"
		} else {

			err := json.Unmarshal(body, &ploginhistoryReq)

			if err != nil {
				ploginhistoryRes.Err_msg = err.Error()+"LHAPIUM"
				ploginhistoryRes.Status="E"
			} else {
				data, err := LoginHistorymethod(ploginhistoryReq)
				if err != nil {
					ploginhistoryRes.Err_msg = err.Error()+"LHAPIMR"
					ploginhistoryRes.Status ="E"
				} else {
					datas, err := json.Marshal(data)

					if err != nil {
						log.Println(err)
						ploginhistoryRes.Err_msg = err.Error()+"LHAPIME"
		                ploginhistoryRes.Status ="E"

					} else {
						fmt.Fprintf(w, string(datas))
						log.Println(string(datas))
					}
				}
			}
		}
	}

}

/*
THIS METHOD IS USED TO INSERT THE LOGIN HISTORY DETAILS IN MEDAPP_LOGIN_HISTORY TABLE
  AND GET LAST INSERTED LOGIN_HISTORY_ID IN MEDAPP_LOGIN_HISTORY TABLE

  PARAMETERS
  --------
  ploginhistoryReq

  ONSUCCESS
  --------
  IT WILL RETURN THE LOGIN HISTORY ID AND STATUS WILL BE "S" AND ERRMS WILL BE NIL IN THE ploginhistoryReS STRUCTURE
  ONERROR
  -------

    IT WILL RETURN THE LOGIN HISTORY ID WILL BE NIL  AND STATUS WILL BE "E" AND ERRMS WILL BE ERROR IN THE ploginhistoryReS STRUCTURE
  AUTHORIZARION:"GOKUL"
DATE:"25-05-24"
*/
// LoginHistorymethod inserts login history into the database and returns the response.
func LoginHistorymethod(ploginhistoryReq LoginHistoryRequest) (LoginHistoryResponse, error) {

	db,err:= DBConnect.GORMDBConnect() 

	if err!=nil{
		ploginhistoryRes.Err_msg = err.Error()+"LHMDBC"
		ploginhistoryRes.Status ="E"
		return ploginhistoryRes,err

	}
 
	ins:=Insert2{
		Login_id: ploginhistoryReq.Login_id,
		Login_Date: time.Now().Format("2006-01-02"),
		Login_Time: time.Now().Format("15:04:05"),
		Created_By: "Gokul",
		Created_Date: time.Now().Format("2006-01-02"),
	}

	result:= db.Table("St832_Medapp_login_history").Create(&ins) 

	if result.Error!=nil{
		log.Println(err)
		ploginhistoryRes.Err_msg = result.Error.Error()+"LHMQEXEC1"
		ploginhistoryRes.Status ="E"
		return ploginhistoryRes,result.Error

	}

	ploginhistoryRes.Login_History_Id=ins.Login_History_id

	log.Println("history-id",ploginhistoryRes.Login_History_Id)
	ploginhistoryRes.Status="S"
	return ploginhistoryRes,nil

	

	// db, err := DBConnect.LocalDBConnect()
	// if err != nil {
	// 	ploginhistoryRes.Err_msg = err.Error()+"LHMDBC"
	// 	ploginhistoryRes.Status ="E"
	// } else {
	// 	defer db.Close()
	// 	//  query to insert login history into the medapp_login_history table.
	// 	corestring := `insert into medapp_login_history(login_id,login_date,login_time,
	// 		created_by,created_date) values(?,CURDATE(),CURTIME(),?,CURDATE())`

	// 	_, err := db.Query(corestring, ploginhistoryReq.Login_id, "Gokul")

	// 	if err != nil {

	// 		log.Println(err)
	// 		ploginhistoryRes.Err_msg = err.Error()+"LHMQEXEC1"
	// 		ploginhistoryRes.Status ="E"

	// 		return ploginhistoryRes, nil
	// 	} else {
	// 		log.Println("Inserted Successfully")
	// 		ploginhistoryRes.Err_msg = " "
	
	// 		ploginhistoryRes.Status ="S"

	// 		corestring:=`select login_history_id from medapp_login_history order by login_history_id desc limit 1`

	// 		rows,err:=db.Query(corestring) 


	// 		if err!=nil{

	// 			ploginhistoryRes.Err_msg=err.Error()+"LHMQEXEC2"
	// 			ploginhistoryRes.Status="E"
	// 			return ploginhistoryRes,nil
	// 		}else{
	// 			for rows.Next(){
	// 				err:=rows.Scan(&ploginhistoryRes.Login_History_Id)

	// 				if err!=nil{
	// 					ploginhistoryRes.Err_msg=err.Error()+"LHMQSCAN2"
	// 					ploginhistoryRes.Status="E"
	// 					return ploginhistoryRes,err
	// 				}
	// 			}
	// 		}
		
	// 		}
	// 		log.Println(ploginhistoryRes.Login_History_Id)

	// 	}
	
	}



