package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/DBConnect"
	"net/http"
)

// StockViewDetails represents the structure of each stock view detail.
type StockViewDetails struct {
	Medicine_name     string  `json:"medicine_name"`
    Brand     string  `json:"brand"`
    Quantity int     `json:"quantity"`
	Unit_Price        float32 `json:"unit_price"`
}

// StockviewResponse represents the structure of the response for stock view.
type StockviewResponse struct {
	Stocksarr []StockViewDetails `json:"stock_view_details"`
	Status    string             `json:"status"`
	Errmsg    string             `json:"errmsg"`
}

/*
THIS API IS USED TO GET THE STOCKS

RESPONSE
------
ONSUCCESS
--------
{
  "stock_view_details": [
    {
      "medicine_name": "Paracetamol",
      "brand": "BrandX",
      "medicine_quantity": 100,
      "unit_price": 1.5
    },
    {
      "medicine_name": "Aspirin",
      "brand": "BrandY",
      "medicine_quantity": 50,
      "unit_price": 2.0
    }
    // Additional StockViewDetails objects can be included in the array
  ],
  "status": "S",
  "errmsg": "",
}

ONERROR
--------
{
  "stock_view_details": "NIL"
  "status": "E",
  "errmsg": "ERROR",
}



*/

// Stockview handles the request for fetching stock view.
func Stockview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	log.Println("Stockview(+)")
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method == "GET" {

		// Initialize response struct.
		var pstockRes StockviewResponse
		pstockRes.Status="S"
		// Call StockViewMethod to fetch stock view data.
		pstockRes, err := StockViewMethod(pstockRes)
		// Process response data and send the response.
		if err != nil {

			pstockRes.Errmsg = err.Error()+"SVAPIMR"
			pstockRes.Status = "E"
		}

		datas, err := json.Marshal(pstockRes)

			if err != nil {
				pstockRes.Errmsg = err.Error()+"SVAPIME"
				pstockRes.Status = "E"
				fmt.Fprintf(w,"ERROR"+err.Error())

			} else {
				// log.Println(string(datas))
				fmt.Fprintf(w, string(datas))

			}
			log.Println("Stockview(-)")

	}

}
/*
THIS METHOD IS USED TO GET THE STOCKS FROM MEDAPP_STOCK TABLE

PARAMETERS
---------
pstockRes

ONSUCCESS
--------
IT WILL RETURN THE STOCKARR AND STATUS WILL BE S AND ERR MSG WILL BE NIL IN THE pstockRes STRUCTURE


ONERROR
--------
IT WILL RETURN THE STOCKARR WILL BE NIL  AND STATUS WILL BE E AND ERR MSG WILL BE ERROR IN THE pstockRes STRUCTURE

AUTHORIZARION:"GOKUL"
DATE:"25-05-24"
*/

// StockViewMethod fetches stock view data from the database.
func StockViewMethod(pstockRes StockviewResponse) (StockviewResponse, error) {

	db,err:= DBConnect.GORMDBConnect() 
	if err!=nil{
			pstockRes.Errmsg = err.Error()+"SRMDBC"
		pstockRes.Status = "E"
		return pstockRes,err

	}

	result := db.Table("st832_medapp_medicine_master mm").
    Select("mm.medicine_name, mm.brand, COALESCE(ms.quantity, 0) as quantity, COALESCE(ms.unit_price, 0) as unit_price").
    Joins(" INNER JOIN st832_medapp_stock ms ON ms.medicine_master_id = mm.medicine_master_id").
    Find(&pstockRes.Stocksarr)


	if result.Error!=nil{
		pstockRes.Errmsg = result.Error.Error()+"SRMQexec1"
		pstockRes.Status = "E"
		return pstockRes,result.Error

	}

	pstockRes.Status="S"
	return pstockRes ,nil

	// db, err := DBConnect.LocalDBConnect()
	// var stockdet StockViewDetails
	// if err != nil {
	// 	pstockRes.Errmsg = err.Error()+"SRMDBC"
	// 	pstockRes.Status = "E"
	// 	return pstockRes,err
	// } else {
	// 	defer db.Close()

	// 	// SQL query to fetch stock view details.
	// 	corestring := `select mm.medicine_name,mm.brand,nvl(ms.quantity,0),nvl(ms.unit_price ,0)
    //                from medapp_medicine_master mm , medapp_stock ms 
    //                where mm.medicine_master_id = ms.medicine_master_id `

	// 	rows, err := db.Query(corestring)

	// 	// Process query results
	// 	if err != nil {
	// 		pstockRes.Errmsg = err.Error()+"SRMQEXEC1"
	// 		pstockRes.Status = "E"
	// 		return pstockRes,err
	// 	} else {
	// 		for rows.Next() {
	// 			err := rows.Scan(&stockdet.Medicine_name, &stockdet.MedicineBrand, &stockdet.Medicine_quantity, &stockdet.Unit_Price)

	// 			if err != nil {

	// 				pstockRes.Errmsg = err.Error()+"SRMQSCAN1"
	// 				pstockRes.Status = "E"
	// 				return pstockRes,err

	// 			} else {
	// 				log.Println("med_name", stockdet.Medicine_name)
	// 				pstockRes.Stocksarr = append(pstockRes.Stocksarr, stockdet)

	// 			}

	// 		}

	// 	}

	// }


}
