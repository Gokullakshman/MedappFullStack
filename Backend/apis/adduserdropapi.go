package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/DBConnect"
	"net/http"
)
// AddUserDropDownResponse represents the response structure for the AddUserDropDown API.
type AddUserDropDownResponse struct{
	Role  []string `json:"role"`
	Status            string           `json:"status"`
	ErrMsg            string           `json:"errmsg"`

}

/*
THIS IS API IS USED TO GET THE USERROLES AND SHOWED IT ON DROPDOWN
RESPONSE:
---------

ONSUCCESS
--------
{
"role":["biller","manager","systemadmin","inventry"]
"status":"S"
"ERRMSG":"NIL"
}

ONERROR
---------
{
"ROLE":"NIL"
"STATUS":"E"
"ERRMSG":"ERROR MSG"
}


*/

func AddUserDropDown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization") 

	
	if r.Method=="GET"{
		var padduserdrop AddUserDropDownResponse 
		padduserdrop.Status="S"

		data,err:=AddUserDropDownMethod(padduserdrop)

		if err!=nil{
			padduserdrop.ErrMsg=err.Error()+"AUDROPAPIMR"
			padduserdrop.Status="E"
		}

			datas,err:=json.Marshal(data)

			if err!=nil{
				padduserdrop.ErrMsg=err.Error()+"AUDROPAPIME"
				fmt.Fprintf(w,"ERROR TAKING DATA"+err.Error())
			    padduserdrop.Status="E"

			}

			log.Println(string(datas))
            fmt.Fprintf(w,string(datas))
		}

		

	}


/*

PURPOSE: THIS METHOD TO FETCH THE DISTINCT ROLES FROM THE MEDAPP_LOGIN TABLE 
PARAMETERS:padduserdrop
ONSUCCESS
-------
{
it return padduserdrop response structure with status = "S" and errormsg = "nil"
}
ONERROR
------
{
it retur the padduserdrop response structure with status = "E" and errormsg ="Errormsg"

}
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"


*/

// AddUserDropDownMethod fetches dropdown data for adding a user.
func AddUserDropDownMethod(padduserdrop AddUserDropDownResponse)(AddUserDropDownResponse,error){


		db, err := DBConnect.LocalDBConnect()
		if err != nil {
			padduserdrop.ErrMsg = err.Error()+"AUDROPMDBC"
			padduserdrop.Status = "E"
			return padduserdrop, err
		}
		defer db.Close()
	   //Select Unique Role from medapp_login
		corestring := `SELECT DISTINCT role FROM st832_medapp_login`
	
		rows, err := db.Query(corestring)
		if err != nil {
			padduserdrop.ErrMsg = err.Error()+"AUDROPMQEXEC1"
			padduserdrop.Status = "E"
			return padduserdrop, err
		}
		defer rows.Close()
	
		for rows.Next() {
			var role string
			err := rows.Scan(&role)
			if err != nil {
				padduserdrop.ErrMsg = err.Error()+"AUDROPMQSCAN1"
				padduserdrop.Status = "E"
				return padduserdrop, err
			}
			padduserdrop.Role = append(padduserdrop.Role, role)
		}
	
		return padduserdrop, nil
	}
	
	
  
  
	   
	 
	

  
  

	

	


