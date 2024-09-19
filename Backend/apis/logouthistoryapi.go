package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	// "time"
)

// LogoutRequest represents the structure of the request for logout history.
type LogoutRequest struct{
	Login_History_id int `json:"login_history_id"`
}
// LogoutResponse represents the structure of the response for logout history.
type LogoutResponse struct{
	Err_msg          string `json:"err_msg`
	Status           string `json:"status"`

}

type UpdateInsert struct{
	logout_date string  `gorm:"logout_date"`

	logout_time  string  `gorm:"logout_time"`

	updated_by string  `gorm:"updated_by"`
	updated_date  string  `gorm:"updated_date"`

}

/*
THIS API IS USED TO UPDATE THE LOGOUT HISTORY 
REQUEST
-------
{
  "login_history_id": 123

}
  RESPONSE
  ------
  ONSUCCESS
  -------
  {
   "err_msg": "",
  "status": "S"
  
  }
  ONERROR
  -------
  {
   "err_msg": "ERROR",
  "status": "E"
  
  }



*/
// Global variables to hold request and response data.
var plogoutReq LogoutRequest
var plogoutRes LogoutResponse





// LogoutHistory handles the request for logout history.
func LogoutHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	
	if r.Method == "POST" {
    log.Println("logouthistory(+)")  

	body, err := ioutil.ReadAll(r.Body)

	if err!=nil{
		plogoutRes.Err_msg=err.Error()+"LOUTAPIBR"
		plogoutRes.Status="E"
		return
	}else{
		log.Println("logoutreq",plogoutReq)
		// Unmarshal request body into plogoutReq.
		err:=json.Unmarshal(body,&plogoutReq)

		if err!=nil{
			plogoutRes.Err_msg=err.Error()+"LOUTAPIUM"
			plogoutRes.Status="E"
			return
		}else{
				// Call LogoutHistoryMethod with the request data.
			data,err:=LogoutHistoryMethod(plogoutReq) 

			if err!=nil{
				plogoutRes.Err_msg=err.Error()+"LOUTAPIMR"
				plogoutRes.Status="E"
				return
			}else{
				// Marshal the response data.
				datas,err:=json.Marshal(data)

				if err!=nil{
					plogoutRes.Err_msg=err.Error()+"LOUTAPIME"
					plogoutRes.Status="E"
					return
				}else{
					fmt.Fprintf(w, string(datas))
					log.Println(string(datas))


				}


			}
			


		}
	}



	


	}

}

/*
THIS METHOD IS USED UPDATE THE LOGOUT DEAILS IN MEDAPP_LOGIN_HISTORY_TABLE


 
  PARAMETERS
  --------
  plogoutReq

  ONSUCCESS
  --------
  IT WILL RETURN  STATUS WILL BE "S" AND ERRMS WILL BE NIL IN THE plogoutRes STRUCTURE
  ONERROR
  -------

     IT WILL RETURN  STATUS WILL BE "E" AND ERRMS WILL BE ERROR IN THE plogoutRes STRUCTURE
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"
*/



// LogoutHistoryMethod updates logout history in the database and returns the response.
func LogoutHistoryMethod(plogoutReq LogoutRequest )( LogoutResponse, error) { 

	// db, err := DBConnect.GORMDBConnect()
	// if err != nil {
	// 	plogoutRes.Status = "E"
	// 	plogoutRes.Err_msg = "Failed to connect to database: " + err.Error()
	// 	return plogoutRes, err
	// }
	 // Close the database connection at the end of the function

	// Initialize UpdateInsert struct with current datetime
	

	// Perform the update operation

	// log.Println(plogoutReq.Login_History_id,"history_id_for_logout")
	// result := db.Table("st832_medapp_login_history").
	// 	Where("login_history_id = ?", plogoutReq.Login_History_id).
	// 	Updates(UpdateInsert{
	// 		logout_date :  time.Now().Format("2006-01-02"),
	// 		logout_time :  time.Now().Format("15:04:05"),
	// 		updated_by:   "Gokul",
	// 		updated_date: time.Now().Format("2006-01-02"),
	// 	})

	// if result.Error != nil {
	// 	plogoutRes.Status = "E"
	// 	plogoutRes.Err_msg = "Failed to update logout history: " + result.Error.Error()
	// 	return plogoutRes, result.Error
	// }

	// if result.RowsAffected == 0 {
	// 	plogoutRes.Status = "E"
	// 	plogoutRes.Err_msg = "No rows updated"
	// 	return plogoutRes, nil
	// }





	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		ploginhistoryRes.Err_msg = err.Error()+"LOUTMDBC"
		plogoutRes.Status="E"
	} else { //query to update logout date and time in medapp_login_history table.
		corestring := `update st832_medapp_login_history set logout_date=CURDATE(),logout_time=NOW(),updated_by=?,updated_date=CURDATE() where login_history_id=?` 
		defer db.Close()    
		
        	// Execute the update query.
		_, err := db.Query(corestring,"Gokul",plogoutReq.Login_History_id)

		if err != nil {

			log.Println(err)
			ploginhistoryRes.Err_msg = "db insert loginhistory error"+err.Error()+"LOUTMQEXEC1"
			return plogoutRes, err
		} else {
			log.Println("Inserted Successfully")
		
			plogoutRes.Status="S"
			}
			log.Println(ploginhistoryRes.Login_History_Id)

		}
		plogoutRes.Status = "S"
		return plogoutRes, nil
	}

