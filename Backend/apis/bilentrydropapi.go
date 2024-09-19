package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/DBConnect"
	"net/http"
)

// BillEntryDropDownResponse represents the response structure for the BillEntryDropDown API.
type BillEntryDropDownResponse struct{
	Medicinenames  []string `json:"medicine_names"`
	Status            string           `json:"status"`
	ErrMsg            string           `json:"errmsg"`

}
/* THIS API IS USED TO FETCH THE MEDICINENAMES 

ONSUCCESS
--------
{
"medicine_names":"["paracetomel","dolo",.....]"
"status":"S"
"errmsg":"nil"

}
ONERROR
------
{
"medicine_names":"nil",
"status":"E"
"err_msg":"errormessage"

}
*/

// BillEntryDropDown is an HTTP handler for fetching dropdown data for bill entry.
func BillEntryDropDown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization") 


	if r.Method=="GET"{
		var pbillentryRep BillEntryDropDownResponse 
		var err error
		pbillentryRep.Status="S"
		pbillentryRep,err =BillEntryDropDownMethod(pbillentryRep)

		if err!=nil{
			pbillentryRep.ErrMsg=err.Error()+"BEDROPAPIMR"
			pbillentryRep.Status="E"
		}

			datas,err:=json.Marshal(pbillentryRep)

			if err!=nil{
				pbillentryRep.ErrMsg=err.Error()+"BEDROPAPIME"
				pbillentryRep.Status="E"
				fmt.Fprintf(w,"ERROR TAKING DATA",err.Error())


			}else{
				log.Println(string(datas))
                fmt.Fprintf(w,string(datas))
			}


		}

	}




	/* This Method is used to return the pbillentryRes structure in medicinenames from using medapp_medicine_master table
	PARAMETERS : pbillentryRep

	ONSUCCESS
	--------
	{
	IT RETURN THE MEDICINENAMES FROM FETCHING TABLE MEDAPP_NEDICINE_MASTER_TABLE IN THE pbillentryRes Response structure and error will be nill
	}
	ONERROR
	-------
	{
	 IT RETURN THE pbillentryRes Response structure with Status "E" AND ERROR WILL BE ERROR MESSAGE
	}
	 AUTHORIZARION:"GOKUL"
DATE:"25-05-24"

	
	*/
// BillEntryDropDownMethod fetches dropdown data for bill entry.
func BillEntryDropDownMethod(pbillentryRep BillEntryDropDownResponse)(BillEntryDropDownResponse,error){

	db,err:= DBConnect.GORMDBConnect()

	if err!=nil{
		pbillentryRep.ErrMsg = err.Error()+"BEDROPMDBC"
			pbillentryRep.Status = "E"
			return pbillentryRep, err

	}

	result:= db.Table("St832_Medapp_Medicine_Master").Select("Medicine_name").Find(&pbillentryRep.Medicinenames)
	if result.Error!=nil{
		pbillentryRep.ErrMsg = result.Error.Error()+"BEDROPMDBC"
			pbillentryRep.Status = "E"
			return pbillentryRep, err

	}

	return pbillentryRep,nil


	// 	db, err := DBConnect.LocalDBConnect()
	// 	if err != nil {
	// 		pbillentryRep.ErrMsg = err.Error()+"BEDROPMDBC"
	// 		pbillentryRep.Status = "E"
	// 		return pbillentryRep, err
	// 	}
	// 	defer db.Close()
	// 	var medicine_name string
	//    //Select Medicinename from Medicine_Master
	// 	corestring := `select medicine_name from medapp_medicine_master`
	
	// 	rows, err := db.Query(corestring)
	// 	if err != nil {
	// 		pbillentryRep.ErrMsg = err.Error()+"BEDROPMQEXEC1"
	// 		pbillentryRep.Status = "E"
	// 		return pbillentryRep, err
	// 	}
	// 	defer rows.Close()
	
	// 	for rows.Next() {
			
	// 		err := rows.Scan(&medicine_name)
	// 		if err != nil {
	// 			pbillentryRep.ErrMsg = err.Error()+"BEDROPMSCAN1"
	// 			pbillentryRep.Status = "E"
	// 			return pbillentryRep, err
	// 		}
	// 		pbillentryRep.Medicinenames=append(pbillentryRep.Medicinenames, medicine_name)
			
	// 	}
	
	}
	
	
  
  
	   