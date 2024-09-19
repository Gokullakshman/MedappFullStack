package apis

import (
	"encoding/json"
	"time"

	"log"
	"medapp/DBConnect"
	"net/http"
)

// UnSoldAmountResponse represents the structure of the response for unsold amount
type UnSoldAmountResponse struct {
	TodayNet float64 `json:"today_net"`
	UnSold   float64 `json:"unsold"`
	Status   string  `json:"status"`
	Errmsg   string  `json:"err_msg`
}

/*
THIS API IS USED TO GET THE TODAY TOTAL  SALES AND UNSOLD MEDICINE AMOUNT
RESPONSE
-------

ONSUCCESS
-------
{
  "today_net": "1234.56",
  "unsold": "789.01",
  "status": "S",
  "err_msg": "NIL",
}
  ONERROR
-------
{
  "today_net": "NIL",
  "unsold": 789.01,
  "status": "E",
  "err_msg": "ERROR",
}


*/
// UnSoldAmount handles the HTTP request for retrieving unsold amount
func UnSoldAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	log.Println("UnSoldAmount (+)")

	if r.Method == "GET" {
		var punsoldRes UnSoldAmountResponse 
		punsoldRes.Status="S"

		amt1, amt2, err := UnsoldAmountMethod() 
		log.Println("err",err)
		if err != nil {
			log.Println(err)
			punsoldRes.Errmsg = err.Error()+"USMAPIMR"
			punsoldRes.Status = "E"
		}else{}
		punsoldRes.TodayNet = amt2
		punsoldRes.UnSold = amt1 

		log.Println("amount",punsoldRes.TodayNet)
        
		// Marshal response into JSON
		datas, err := json.Marshal(punsoldRes)
		if err != nil {
			log.Println(err)
			punsoldRes.Errmsg = err.Error()+"USMAPIME"
			punsoldRes.Status = "E"

		}
       // Write response data
	   log.Println(string(datas))
		w.Write(datas)
		log.Println("UnSoldAmount (-)")
	}
}
/*THIS METHOD IS USED TO GET THE TOTAL SALES FROM MEDAPP_BILL_MASTER TABLE  AND UNSOLD MEDICINE AMOUNT FROM MEDAPP_STOCK

ONSUCCESS
---------
IT WILL RETURN THE TODAY TOTAL SALES , UNSOLDMEDICINE AMOUNT AND ERR WILL BE "NIL"


ONERROR
---------
IT WILL RETURN THE TODAY TOTAL SALES WILL BE 0  , UNSOLDMEDICINE AMOUNT WILL BE 0  AND ERR WILL BE "ERRMSG"
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"
*/
// UnsoldAmountMethod retrieves the unsold amount from the database
func UnsoldAmountMethod() (float64, float64, error) {
	log.Println("UnsoldAmountMethod (+)")
     var inventry float64
	var sales float64
   
   
	db,err:= DBConnect.GORMDBConnect()

	if err!=nil{
		return 0, 0, err

	}

	

	result:= db.Table("st832_medapp_stock").Select("COALESCE(SUM(QUANTITY*UNIT_PRICE),0) as amount").Find(&inventry)

	if result.Error!=nil{
		return 0,0,result.Error
	}


	result1:= db.Table("st832_medapp_bill_master").Select("COALESCE(SUM(NET_PRICE),0) as tdy").
	Where("bill_date =?",time.Now().Format("2006-01-02")).Find(&sales)

	if result1.Error!=nil{
		
		return 0,0,result1.Error

	}
   
    

	return inventry,sales,nil



	// SELECT NVL(SUM(NET_PRICE),0)
	// FROM MEDAPP_BILL_MASTER
	// WHERE BILL_DATE = CURRENT_DATE();`
	// Connect to the local database
	// db, err := DBConnect.LocalDBConnect()
	// if err != nil {
		

	// 	return 0, 0, err
	// }
	// defer db.Close()

	// var inventry float64
	// var sales float64
  	// // Query to retrieve total sales value for today
	// corestring := `SELECT NVL(SUM(QUANTITY * UNIT_PRICE),0)
	// 	FROM MEDAPP_STOCK;`

	// rows, err := db.Query(corestring)
	// if err != nil {
	// 	log.Println(err)
	// 	return 0, 0, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&inventry)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return 0, 0, err
	// 	}
	// }
	// corestring1 := `
   	// SELECT NVL(SUM(NET_PRICE),0)
	// FROM MEDAPP_BILL_MASTER
	// WHERE BILL_DATE = CURRENT_DATE();`
	// rows1, err := db.Query(corestring1)
	// if err != nil {
	// 	log.Println(err)
	// 	return 0, 0, err

	// }

	// for rows1.Next() {
	// 	err := rows1.Scan(&sales)

	// 	if err != nil {
	// 		log.Println(err)
	// 		return 0, 0, err

	// 	}
	// }

	// log.Println("UnsoldAmountMethod (-)")
	// return inventry, sales, nil
}
