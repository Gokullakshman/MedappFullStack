package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
)

// Structs for request, response, and details of sales report.
type SalesReportRequest struct {
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type SalesReportDetails struct {
	Bill_No       string `json:"bill_no" gorm:"column:BILL_NO"`
	Bill_Date     string `json:"bill_date" gorm:"column:BILL_DATE"`
	Medicine_Name string `json:"medicine_name" gorm:"column:MEDICINE_NAME"`
	Quantity      int    `json:"quantity" gorm:"column:QUANTITY"`
	Amount        int    `json:"amount" gorm:"column:AMOUNT"`

}

type SalesReportResponse struct {
	SalesArr []SalesReportDetails `json:"sales_arr"`
	ErrMsg   string               `json:"err_msg"`
	Status   string               `json:"status"`
}

/* THIS API IS USED TO GET THE SALES REPORT FOR SELECTED DATES
REQUEST
------
{
 "from_date": "2024-01-01",
  "to_date": "2024-06-30"

}
  RESPONSE
  ------
  {
  "sales_arr": [
    {
      "bill_no": "B12345",
      "bill_date": "2024-06-20",
      "medicine_name": "Paracetamol",
      "quantity": 10,
      "amount": 100
    },
    {
      "bill_no": "B12346",
      "bill_date": "2024-06-20",
      "medicine_name": "Aspirin",
      "quantity": 5,
      "amount": 50
    }
    // Additional SalesReportDetails objects can be included in the array
  ],
  "err_msg": "",
  "status": "S"

  }


*/
// Handler function to fetch sales report.
func FetchSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")


	if r.Method == "POST" {
		log.Println("fetchsales(+)")
		// Initialize request and response variables.
		var psalesReq SalesReportRequest
		var psalesRes SalesReportResponse

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			psalesRes.Status = "E"
			psalesRes.ErrMsg = err.Error() + "SRAPIBR"
		} else {
			// Unmarshal request body into psalesReq.
			err = json.Unmarshal(body, &psalesReq) 
			log.Println(psalesReq)
			if err != nil {
				psalesRes.Status = "E"
				psalesRes.ErrMsg = err.Error() + "SRAPIUM"
			} else {
				// Call FetchSalesMethod with the request data.
				psalesRes, err = FetchSalesMethod(psalesReq, psalesRes)
				if err != nil {
					psalesRes.Status = "E"
					psalesRes.ErrMsg = err.Error() + "SRAPIMR"

				}

			}

		}
		// Marshal the response data and send it.
		datas, err := json.Marshal(psalesRes)
		if err != nil {
			psalesRes.Status = "E"
			psalesRes.ErrMsg = err.Error() + "SRAPIME"
			fmt.Fprintf(w, "ERROR "+err.Error())
		} else {
			// log.Println(string(datas))
			fmt.Fprintf(w, string(datas))

		}
	}
}

/*
THIS METHOD IS USED TO GET THE SALES FROM JOINING TWO TABLES MEDAPP_BILL_DETAILS AND MEDAPP_MASTER_ID
  PARAMETERS
  --------
  psalesReq,psalesRes

  ONSUCCESS
  --------
  IT WILL RETURN  SALES ARRAY AND STATUS WILL BE "S" AND ERRMS WILL BE NIL IN THE psalesRes STRUCTURE
  ONERROR
  -------

     IT WILL RETURN SALES ARRAY WILL BE NIL AND  STATUS WILL BE "E" AND ERRMS WILL BE ERROR IN THE psalesRes STRUCTURE
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"
*/

// Function to fetch sales report from the database.
func FetchSalesMethod(psalesReq SalesReportRequest, psalesRes SalesReportResponse) (SalesReportResponse, error) {
	db, err := DBConnect.GORMDBConnect()
	if err != nil {
		psalesRes.ErrMsg = err.Error() + "SRMDBC"
		psalesRes.Status = "E"
		return psalesRes, err
	}
	result := db.Table("Medapp_bill_details MBD").
		Select("MBD.BILL_NO, MBM.BILL_DATE, MMM.MEDICINE_NAME, MBD.QUANTITY, MBD.AMOUNT").
		Joins("inner join st832_medapp_bill_master MBM on mbd.bill_no = mbm.bill_no").
		Joins("inner join st832_medapp_medicine_master mmm on mmm.medicine_master_id = mbd.medicine_master_id").
		Where("mbm.bill_date between ? and ?", psalesReq.FromDate, psalesReq.ToDate).Find(&psalesRes.SalesArr)

	if result.Error != nil {
		psalesRes.ErrMsg = result.Error.Error() + "SRMQEXEC1"
		psalesRes.Status = "E"
		return psalesRes, result.Error

	}
    
    log.Println("fetchsales(-)") 

	log.Println(psalesRes.SalesArr)
	psalesRes.Status = "S"
	return psalesRes, nil

	// 	defer db.Close()
	//    // SQL query to fetch sales report details between the specified dates.

	// 		coreString := `SELECT MBD.BILL_NO, MBM.BILL_DATE, MMM.MEDICINE_NAME, MBD.QUANTITY, MBD.AMOUNT
	// 					FROM MEDAPP_BILL_DETAILS MBD
	// 					INNER JOIN MEDAPP_BILL_MASTER MBM ON MBD.BILL_NO = MBM.BILL_NO
	// 					INNER JOIN MEDAPP_MEDICINE_MASTER MMM ON MMM.MEDICINE_MASTER_ID = MBD.MEDICINE_MASTER_ID
	// 					WHERE MBM.BILL_DATE BETWEEN ? AND ?`

	// 	rows, err := db.Query(coreString, psalesReq.FromDate, psalesReq.ToDate)
	// 	if err != nil {

	// 		psalesRes.ErrMsg=err.Error()+"SRMQEXEC1"

	// 		psalesRes.Status="E"
	// 		return psalesRes,err

	// 	}
	// 	defer rows.Close()

	// 	// var salesDetails []SalesReportDetails

	// 	// Iterate through query results and populate salesDetails.
	// 	for rows.Next() {
	// 		var salesdet SalesReportDetails
	// 		err := rows.Scan(&salesdet.BillNo,&salesdet.BillDate,&salesdet.MedicineName,&salesdet.Quantity,&salesdet.Amount)

	// 		if err != nil {

	// 			psalesRes.ErrMsg = err.Error()+"SRMSCAN1"
	// 			psalesRes.Status = "E"
	// 		} else {
	// 			// log.Println("med_name", stockdet.Medicine_name)
	// 			psalesRes.SalesArr = append(psalesRes.SalesArr, salesdet)

	// 		}

	// 	}
	// 	// log.Println(psalesRes.SalesArr)

	//    // Populate response struct and return.

	// 	psalesRes.Status="S"
	// 	return psalesRes,nil
}

