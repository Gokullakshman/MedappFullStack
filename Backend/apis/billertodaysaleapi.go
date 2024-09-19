package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/DBConnect"
	"net/http"
	"time"
)

// Define the request structure for retrieving today's sales by a biller.
	type BillerTodaySalesRequest struct{
		User_Id string `json:"user_id"`
	}



	// Define the response structure for today's sales.

	type BillerTodaySaleResponse struct {
		TodaySale float64   `json:"today_sales"`
		YesterdaySale float64 `json:"yesterday_sales"`
		Status string    `json:"status"`
		Errmsg string   `json:"errmsg"`
	}
/*
THIS API USED TO GET PARTICULAR THE BILLERS TODAY SALES AND YESTERDAY SALES

REQUEST
------
{
"user_id":"biller"

}
RESPONSE
--------
ONSUCCESS
{
"today_sales":"2000"
"yesterday_sales":"1200"
"status":"S"
"err_msg":"nil"
}
ONERROR
-------
{
"today_sales":"nil"
"yesterday_sales":"nil"
"status":"E"
"err_msg":"ERRORMSG"

}

*/

// Handler function to retrieve today's sales by a biller.
func BillerTodaySales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")


	if r.Method == "POST" {
        // Initialize request and response structs
		var pbilltdyReq BillerTodaySalesRequest
		 var pbilltdyRes BillerTodaySaleResponse 
		 var error1 error
		 pbilltdyRes.Status="S"


    log.Println("TodaySales(+)")  


		
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		pbilltdyRes.Status = "E"
		pbilltdyRes.Errmsg = err.Error()+"BTSAPIBR"
	} else {
			// Unmarshal JSON into request struct.
		err := json.Unmarshal(body, &pbilltdyReq) 
		log.Println("billeruser_id",pbilltdyReq.User_Id)
		if err != nil {
			log.Println(err)
			pbilltdyRes.Status = "E"
			pbilltdyRes.Errmsg = err.Error()+"BTSAPIUM"
		} else{
			 	// Call method to retrieve today's sales.
				 pbilltdyRes,error1 = BillerTodaySalesMethod(pbilltdyReq ,pbilltdyRes ) 
				 if error1!=nil{
				  pbilltdyRes.Errmsg=error1.Error()+"BTSAPIMR"
				  pbilltdyRes.Status="E"
				 }

		}
    }
	datas,err:=json.Marshal(pbilltdyRes)


	 if err!=nil{
		 pbilltdyRes.Errmsg=err.Error()+"BTSAPIME"
		 fmt.Fprintf(w,"ERROR"+err.Error())
	 }

	   log.Println(string(datas))
	   fmt.Fprintf(w,string(datas))
	  


	}

}

/*THIS METHOD IS USED TO GET THE PARTICULAR BILLER TODAY SALES AND YESTERDAY SALES
PARAMETER:pbilltdyReq,pbilltdyRes

ONSUCCESS
------
IT WILL RETURN THE TODAY SALES AND YESTERDAY IN THE pbilltdyRes RESPONSE STRUCTURE AND ERROR WILLBE NIL

ONERROR
------
IT WILL RETURN THE pilltdyRes Response will be STATUS "E" AND ERROR WILL BE ERRORMSG
AUTHORIZARION:"GOKUL"
DATE:"25-05-24"

*/


// Method to retrieve today's sales by a biller.
func BillerTodaySalesMethod(pbilltdyReq BillerTodaySalesRequest,pbilltdyRes BillerTodaySaleResponse ) (BillerTodaySaleResponse ,error) {
	// var sales float64 
	// var yesterdaysales float64 
	var user int 
	 time1:=time.Now().Format("2006-12-01")
	 

	db,err:=DBConnect.GORMDBConnect()

	if err!=nil{
		pbilltdyRes.Errmsg=err.Error()
			pbilltdyRes.Status="E"
			return pbilltdyRes,err

	}

	subquery:= db.Table("ST832_MEDAPP_LOGIN").Select("LOGIN_ID").Where("USER_ID=?",pbilltdyReq.User_Id).Find(&user)

	if subquery.Error!=nil{
			pbilltdyRes.Errmsg=subquery.Error.Error()
		pbilltdyRes.Status="E"
		return pbilltdyRes,subquery.Error

	}

	result:= db.Table("St832_Medapp_Bill_Master").Select("coalesce(sum(net_price),0) as tdy").Where("login_id=? and bill_date=?",user,time.Now().AddDate(0, 0, 0).Format("2006-01-02")).Find(&pbilltdyRes.TodaySale)

    log.Println("loginid",user)
	log.Println("time",time1)


   log.Println("tdy",pbilltdyRes.TodaySale)

	if result.Error!=nil{
		pbilltdyRes.Errmsg=result.Error.Error()
		pbilltdyRes.Status="E"
		return pbilltdyRes,result.Error

	}

	result1:= db.Table("St832_Medapp_Bill_Master").Select("coalesce(sum(net_price),0) as yes").
	Where("login_id=? and bill_date=?",user,time.Now().AddDate(0, 0, -1).Format("2006-01-02")).Find(&pbilltdyRes.YesterdaySale)

	if result1.Error!=nil{
		pbilltdyRes.Errmsg=result1.Error.Error()
		pbilltdyRes.Status="E"
		return pbilltdyRes,result1.Error

	}
	pbilltdyRes.Status="S"

	return pbilltdyRes ,nil



	// db.Table("Medapp_Bill_Master").Select("collasce(sum(net_price),0) as ")

	// db,err:= DBConnect.GORMDBConnect()

	// if err!=nil{
	// 	pbilltdyRes.Errmsg=err.Error()
	// 	pbilltdyRes.Status="E"
	// 	return pbilltdyRes,err
	// }

	



// 	  db , err:= DBConnect.LocalDBConnect()

//   if err!=nil{
//    pbilltdyRes.Errmsg=err.Error()
//    pbilltdyRes.Status="E"
//   }else{
	
//     defer db.Close()

// 		// Query to calculate the sum of sales for the given biller's user ID.
// 	corestring:=`SELECT NVL(SUM(NET_PRICE),0)
// 	FROM MEDAPP_BILL_MASTER
// 	 WHERE LOGIN_ID =(SELECT LOGIN_ID FROM MEDAPP_LOGIN WHERE USER_ID = ?) AND BILL_DATE=CURDATE()`


// 	rows,err:=db.Query(corestring,&pbilltdyReq.User_Id)
// 	if err!=nil{

// 		pbilltdyRes.Errmsg=err.Error()
// 	    pbilltdyRes.Status="E"
// 		return pbilltdyRes,err

// 	}

// 	for rows.Next(){
// 		err:=rows.Scan(&sales)
		

// 		if err!=nil{
// 			pbilltdyRes.Errmsg=err.Error()
// 	       pbilltdyRes.Status="E"
// 		   return pbilltdyRes,err
			
// 		}
//   }
//   corestring1:=`SELECT NVL(SUM(NET_PRICE),0)
// 	FROM MEDAPP_BILL_MASTER
// 	 WHERE LOGIN_ID =(SELECT LOGIN_ID FROM MEDAPP_LOGIN WHERE USER_ID = ?) AND BILL_DATE=CURDATE()-INTERVAL 1 DAY`

// 	 rows1,err1:=db.Query(corestring1,&pbilltdyReq.User_Id)

// 	 if err1!=nil{

// 		pbilltdyRes.Errmsg=err1.Error()
// 	    pbilltdyRes.Status="E"
// 		return pbilltdyRes,err

// 	}

// 	for rows1.Next(){
// 		err:=rows1.Scan(&yesterdaysales)
		

// 		if err!=nil{
// 			pbilltdyRes.Errmsg=err.Error()
// 	       pbilltdyRes.Status="E"
// 		   return pbilltdyRes,err
			
// 		}
//   }



//   log.Println(sales)
 
// }
// pbilltdyRes.Status="S"
// pbilltdyRes.TodaySale=sales
// pbilltdyRes.YesterdaySale=yesterdaysales
// return pbilltdyRes,nil
}

