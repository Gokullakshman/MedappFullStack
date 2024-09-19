package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	// "time"
   "gorm.io/gorm"
)

// AddUserRequest represents the request structure for adding a user.
type AddUserRequest struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// AddUserDetails represents the details of a user.
// type AddUserDetails struct {
// 	User_id  string `json:"user_id"`
// 	Password string `json:"password"`
// 	Role     string `json:"role"`
// }

// AddUserResponse represents the response structure for adding a user.
type AddUserResponse struct {
	Status string `json:"status"`
	Errmsg string `json:"errmsg"`
}

type MEDAPP_LOGIN struct{

	gorm.Model
	UserID      string    `gorm:"column:USER_ID;uniqueIndex"`
	Password    string    `gorm:"column:PASSWORD"`
	Role        string    `gorm:"column:ROLE"`
	CreatedBy   string    `gorm:"column:CREATED_BY"`
	CreatedDate string `gorm:"column:CREATED_DATE"`
	UpdatedBy   string    `gorm:"column:UPDATED_BY"`
	UpdatedDate string `gorm:"column:UPDATED_DATE"`
}

/* This is API is used to add a New user into Database
REQUEST: 
PUT
---
"user_id":"biller1"
"password":"123"
"role":"biller"

RESPONSE:
ON SUCCESS
-----------
{
"status":"S"
"errmsg":"nil"
}

ON ERROR
---------
{
"status":"s"
"errmsg":"errmsg"
}

*/

// FetchUser is an HTTP handler for adding a user.
func FetchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	if r.Method == "PUT" {
		log.Println("aDDUSER(+)")

		var padduserRes AddUserResponse
		padduserRes.Status = "S"
		var padduserReq AddUserRequest

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			padduserRes.Status = "E"
			padduserRes.Errmsg = err.Error()+"AUAPIRB"
		} else {

			err := json.Unmarshal(body, &padduserReq)
			if err != nil {
				log.Println(err)
				padduserRes.Status = "E"
				padduserRes.Errmsg = err.Error()+"AUAPIUM"
			} else {

				log.Println("adduser", padduserReq)
				err := FetchUserMethod(padduserReq)
				if err != nil {
					if err.Error() == "no rows affected" {
						log.Println(err.Error())
						padduserRes.Errmsg = err.Error()+"AUAPINROWS"
						padduserRes.Status = "E"
					} else {
						log.Println(err.Error())
						padduserRes.Errmsg = err.Error()+"AUAPIMR"
						padduserRes.Status = "E"
					}
				}
			}
		}
		datas, err := json.Marshal(padduserRes)
		if err != nil {
			log.Println(err.Error())
			padduserRes.Errmsg = err.Error()+"AUAPIME"
			padduserRes.Status = "E"
		}
		log.Println(string(datas))
		w.Write(datas)
	}
	
}

/*Purpose: This is used to add the new user into database if not exists 

parameters:padduserReq

ON SUCCESS
--------------------
IT RETURN A NIL ON ERROR

ON ERROR
---------
IT RETURN A ERRORMESSAGE ON ERROR.

AUTHORIZARION:"GOKUL"
DATE:"25-05-24"


*/

// FetchUserMethod fetches user details from the database.
func FetchUserMethod(req AddUserRequest) error {

	db, err := DBConnect.LocalDBConnect()
	// db,err:= DBConnect.GORMDBConnect()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return err
	// }
	// var exists MEDAPP_LOGIN
	
	// result:= db.Table("MEDAPP_LOGIN").Where("user_id=?",req.User_id).First(exists)
	// if(result.Error!=nil && result.Error!=gorm.ErrRecordNotFound){
	// 	return result.Error
	// }
	// if(result.Error==nil){
	// 	return fmt.Errorf("no rows affected")
		
	// }
	// if(result.Error==gorm.ErrRecordNotFound){
	// 	new:=MEDAPP_LOGIN{
	// 		UserID: req.User_id,
	// 		Password: req.Password,
	// 		Role: req.Role,
	// 		CreatedBy: "Gokul",
	// 		CreatedDate: time.Now().Format("2006-01-02"),
	// 		UpdatedBy: "Gokul",
	// 		UpdatedDate: time.Now().Format("2006-01-02"),
	// 	}
	// 	result:=db.Create(&new) 
	// 	if(result.Error !=nil){
	// 		return result.Error
	// 	}
	// }
	
	//This query inserts a new user into the MEDAPP_LOGIN table if a user with the same USER_ID does not already exist.


	corestring := `INSERT INTO ST832_MEDAPP_LOGIN (USER_ID, PASSWORD, ROLE, CREATED_BY, CREATED_DATE, UPDATED_BY, UPDATED_DATE)
		SELECT ?,?,?,?,NOW(),?,NOW()
		FROM DUAL
		WHERE NOT EXISTS (SELECT 1 FROM ST832_MEDAPP_LOGIN WHERE USER_ID = ?)`

	result, err := db.Exec(corestring, req.User_id, req.Password, req.Role, "Gokul", "Gokul", req.User_id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	rowsaffect, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return err

	}

	log.Println("rows", rowsaffect)
	if rowsaffect == 0 {
		log.Println("no rows affected")
		return fmt.Errorf("no rows affected")

	}

	return nil
}
