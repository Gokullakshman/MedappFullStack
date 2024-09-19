package apis

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// LoginRequest represents the structure of the login request.
type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// UserDetails represents the structure of user details.
type UserDetails struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Role     string `json:"role"`
	LoginID  int    `json:"login_id"`
}

// LoginResponse represents the structure of the response for login validation.
type LoginResponse struct {
	UserDetails UserDetails `json:"user_details"`
	ErrMsg      string      `json:"err_msg"`
	Status      string      `json:"status"`
}

/*
THIS API IS USED TO VALIDATE THE USER DETAILS TO ENTER THE APP
REQUEST
------
{
 "user_id": "john_doe",
  "password": "password123"

}
  RESponse
  -------
  {
  "user_details": {
    "user_id": "john_doe",
    "password": "password123",  
    "role": "admin",
    "login_id": 12345
  },
  "err_msg": "",
  "status": "S"
}


*/

// Handler function to validate user login.
func Loginvalidation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")


	if r.Method == "POST" {
		

	

	var logReq LoginRequest
	var logRes LoginResponse 
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		logRes.Status = "E"
		logRes.ErrMsg = err.Error()+"LAPIBR"
		writeJSONResponse(w, logRes)
		return
	}

	err = json.Unmarshal(body, &logReq) 
	log.Println("user_id,",logReq.UserID)
	log.Println("pass,",logReq.Password)

	if err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		logRes.Status = "E"
		logRes.ErrMsg = err.Error()+"LAPIUM"
		writeJSONResponse(w, logRes)
		return
	}

	logRes, err = LoginValidationMethod(logReq)
	if err != nil {
		log.Printf("Error validating login: %v", err)
		logRes.Status = "E"
		logRes.ErrMsg = err.Error()+"LAPIME"
		writeJSONResponse(w, logRes)
		return
	}
	

	writeJSONResponse(w, logRes)
}
}

func writeJSONResponse(w http.ResponseWriter, response LoginResponse) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
	log.Println("login_data",string(jsonData))
	w.Write(jsonData)
}

/*THIS METHOD IS USED TO VALIDATE USING SELECT QUERY USING MEDAPP_LOGIN TABLE
PARAMETERS
----------
req

ONSUCCESS
-------
IT WILL RETURN THE USERDETAILS AND STATUS WILL BE "S " AND ERRMSG WILL BE NIL IN THE LOGINRESPONSE STRUCTURE

ONERROR
-------

IT WILL RETURN THE USERDETAILS WILL BE NIL AND STATUS WILL BE "E" AND ERRMSG WILL BE ERROR IN THE LOGINRESPONSE STRUCTURE
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"



*/

// LoginValidationMethod is the method to validate user login.
func LoginValidationMethod(req LoginRequest) (LoginResponse, error) {

	var logRes LoginResponse
	// db,err:= DBConnect.GORMDBConnect()

	// if err!=nil{
	// 	logRes.Status = "E"
	// 	logRes.ErrMsg = err.Error()+"LMDBC"
	// 	return logRes, err

	// }

	// result:= db.Table("St832_Medapp_Login").Select("user_id,password,role,login_id").
	// Where("user_id=? and password =?",req.UserID,req.Password).
	// Find(&logRes.UserDetails) 
	// if result.Error!=nil{
	// 		logRes.Status = "E"
	// 	logRes.ErrMsg = err.Error()+"LMDBC"
	// 	return logRes, err

	// }

	// logRes.Status="S"
	// return logRes ,nil

	// Find(&logRes.UserDetails.UserID,&logRes.UserDetails.Password,&logRes.UserDetails.Role,&logRes.us)
	// var logRes LoginResponse

	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		logRes.Status = "E"
		logRes.ErrMsg = err.Error()+"LMDBC"
		return logRes, err
	}
	defer db.Close()

	// SQL query to validate user login.
	corestring := `SELECT user_id, password, role, login_id FROM St832_medapp_login WHERE user_id = ? AND password = ?`
	rows, err := db.Query(corestring, req.UserID, req.Password)
	if err != nil {
		log.Printf("Query error: %v", err)
		logRes.Status = "E"
		logRes.ErrMsg = err.Error()+"LMQEXEC1"
		return logRes, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&logRes.UserDetails.UserID, &logRes.UserDetails.Password, &logRes.UserDetails.Role, &logRes.UserDetails.LoginID)
		if err != nil {
			log.Printf("Error scanning rows: %v", err)
			logRes.Status = "E"
			logRes.ErrMsg = err.Error()+"LMQSCAN1"
			return logRes, err
		}
		
	} 
	if(logRes.UserDetails.LoginID==0){
		logRes.Status="E"
		return logRes ,nil

	}
	logRes.Status="S"
	return logRes ,nil


}
