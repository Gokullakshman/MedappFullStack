package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
)

// Define the request structure for fetching the brand of a medicine.
type BrandRequest struct {
	MedicineName string `json:"medicine_name"`
}

// Define the response structure for fetching the brand of a medicine.
type BrandResponse struct {
	Brand  string `json:"brand"`
	Status string `json:"status"`
	ErrMsg string `json:"errmsg"`
}
type MedicineBrand struct{
	Brand string `gorm:"column:brand"`
}
/*
THIS IS API USEDTO FETCH THE MEDICINES BRAND 

REQUEST
-----
{
"medicine_name":"dolo"
}
RESPONSE
-------
ONSUCCESS
{
"brand":"D"
"status":"S"
"err_msg":"nil"

}
ONERROR
{
"brand":"nil"
"status":"E"
"err_msg":"ERRORMSG"

}

*/



// Handler function to fetch the brand of a medicine.
func FetchBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	if r.Method == "POST" {
		log.Println("Fetchbrand(+)")

		// Initialize request and response structs.
		var pbrandReq BrandRequest
		var pbrandRes BrandResponse
		pbrandRes.Status = "S"

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			pbrandRes.ErrMsg = err.Error() + "FBAPIBR"
			pbrandRes.Status = "E"
		} else {
			// Unmarshal JSON into request struct.
			err := json.Unmarshal(body, &pbrandReq)
			if err != nil {
				pbrandRes.ErrMsg = err.Error() + "FBAPIUM"
				pbrandRes.Status = "E"

			} else {
				// Call method to fetch the brand of the medicine.
				pbrandRes, err = FetchBrandMethod(pbrandReq, pbrandRes)

				if err != nil {
					pbrandRes.ErrMsg = err.Error() + "FBAPIUM"
					pbrandRes.Status = "E"
				}

			}
		}
		// Marshal response data and send it as a HTTP response.
		datas, err := json.Marshal(pbrandRes)
		if err != nil {
			pbrandRes.ErrMsg = err.Error() + "FBAPIME"
			pbrandRes.Status = "E"

		} else {
			fmt.Fprintf(w, string(datas))
			log.Println(string(datas))

		}
	}
}


/*
THIS METHOD IS USED TO FETCH THE MEDICINE BRAND FROM THE MEDAP_MASTER_MASTER TABLE FOR SELECTED MEDICINE NAME
PARAMETERS:
---------
pbrandReq ,pbrandRes

ONSUCCESS
--------
IT WILL RETURN THE BRAND AND STATUS "S " AND ERRMSG WILL BE THE NULL IN THE pbrandRes Response STRUCTURE

ONERROR
--------
IT WILL RETURN THE STATUS "E" AND ERRMSG WILL BE THE ERROR AND BRAND WILL BE NILL IN THE pbrandRes Response STRUCTURE
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"

*/

// Method to fetch the brand of a medicine.
func FetchBrandMethod(pbrandReq BrandRequest, pbrandRes BrandResponse) (BrandResponse, error) {
	// Connect to the local database.
	db, err := DBConnect.GORMDBConnect()
	var brand MedicineBrand

	if err != nil {
		pbrandRes.ErrMsg = err.Error() + "FBMDBC"
		pbrandRes.Status = "E"

	} else {
		

		// SQL query to fetch the brand of the medicine.
		
		result:=db.Table("st832_medapp_medicine_master").Select("brand").Where("medicine_name=?",pbrandReq.MedicineName).Find(&brand)
		if result.Error!=nil{
			pbrandRes.ErrMsg=err.Error()+"FBMQEXEC1"
			pbrandRes.Status = "E"
			return pbrandRes, err
		}
		fmt.Println(brand.Brand)

		// if err != nil {

		// 	pbrandRes.ErrMsg = err.Error() + "FBMQEXEC1"
		// 	pbrandRes.Status = "E"
		// 	return pbrandRes, err

		// }
		// Iterate through the rows and extract the brand.
		// for rows.Next() {
		// 	err := rows.Scan(&brand)
		// 	log.Println(brand)

		// 	if err != nil {
		// 		pbrandRes.ErrMsg = err.Error() + "FBMQSCAN1"
		// 		pbrandRes.Status = "E"
		// 		return pbrandRes, err

		// 	}

		// }

	} // Prepare successful response.
	pbrandRes.Status = "S"
	pbrandRes.Brand = brand.Brand
	return pbrandRes, nil
}
