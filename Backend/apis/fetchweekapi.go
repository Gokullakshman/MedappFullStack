package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"medapp/DBConnect"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type WeeklySales struct {
	Mon float64 `json:"mon"`
	Tue float64 `json:"tue"`
	Wed float64 `json:"wed"`
	Thu float64 `json:"thu"`
	Fri float64 `json:"fri"`
	Sat float64 `json:"sat"`
	Sun float64 `json:"sun"`
}

type SalesManSales struct {
	Userid  string  `json:"user_id"`
	Loginid int     `json:"login_id"`
	Net     float64 `json:"net"`
}

type ThisMonthSales struct {
	Userid string  `json:"user_id`
	Net    float64 `json:"net"`
}

type WeekResponse struct {
	WeekArr      WeeklySales `json:"week_obj"`
	DailyArr     []SalesManSales `json:"daily_arr"`
	ThisMonthArr []ThisMonthSales `json:"this_month_arr"`
	Status       string           `json:"status"`
	ErrMsg       string           `json:"err_msg"`
}



func FetchWeekSaless(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	log.Println("FetchWeekSaless(+)")

	if r.Method == "GET" {
		// Reset the response variable for each request
		var pweekRes WeekResponse
		var pweekDet WeeklySales
		var pdailydet SalesManSales
		var pthismonthdet ThisMonthSales
		var err1 error
		// Call method to fetch login history data.
		pweekRes.Status = "S"
		pweekDet, pweekRes, err1 = WeekMethod(pweekDet, pweekRes)

		if err1 != nil {
			pweekRes.Status = "E"
			pweekRes.ErrMsg = err1.Error()
		} else {
			pdailydet, pweekRes, err1 = DailyMethod(pdailydet, pweekRes)
			if err1 != nil {
				pweekRes.Status = "E"
				pweekRes.ErrMsg = err1.Error()
			} else {
				pthismonthdet, pweekRes, err1 = ThisMonthSalesMethod(pthismonthdet, pweekRes) 
				if err1!=nil{
					pweekRes.Status = "E"
					pweekRes.ErrMsg = err1.Error()

				}
			}
		}

		datas, err := json.Marshal(pweekRes)
		if err != nil {
			pweekRes.ErrMsg = err.Error() + "FBAPIME"
			pweekRes.Status = "E"
		} else {
			fmt.Fprintf(w, string(datas))
			log.Println(string(datas))
		}
	}
}

func WeekMethod(pweekDet WeeklySales, pweekRes WeekResponse) (WeeklySales, WeekResponse, error) {
	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		pweekRes.ErrMsg = err.Error()
		pweekRes.Status = "E"
		return pweekDet, pweekRes, err
	}

	err1 := db.QueryRow(
		`SELECT 
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() THEN net_price ELSE 0 END), 0) AS today_sales,
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() - INTERVAL 1 DAY THEN net_price ELSE 0 END), 0) AS yesterday_sales,
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() - INTERVAL 2 DAY THEN net_price ELSE 0 END), 0) as thirdday,
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() - INTERVAL 3 DAY THEN net_price ELSE 0 END), 0) AS fourthday,
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() - INTERVAL 4 DAY THEN net_price ELSE 0 END), 0) AS fifthday_sales,
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() - INTERVAL 5 DAY THEN net_price ELSE 0 END), 0) AS sixth_sales,
			COALESCE(SUM(CASE WHEN DATE(bill_date) = CURDATE() - INTERVAL 6 DAY THEN net_price ELSE 0 END), 0) AS seventh_sales
		FROM st832_medapp_bill_master`).Scan(&pweekDet.Mon, &pweekDet.Tue, &pweekDet.Wed, &pweekDet.Thu, &pweekDet.Fri, &pweekDet.Sat, &pweekDet.Sun)

	if err1 != nil {
		pweekRes.ErrMsg = err1.Error()
		pweekRes.Status = "E"
		return pweekDet, pweekRes, err1
	}

	log.Println(pweekDet,"weekly log")

    pweekDet.RoundValues()
	pweekRes.WeekArr=pweekDet 
	pweekRes.Status = "S"
	pweekRes.ErrMsg = ""
	return pweekDet, pweekRes, nil
}

func DailyMethod(pdailydet SalesManSales, pweekRes WeekResponse) (SalesManSales, WeekResponse, error) {
	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		pweekRes.ErrMsg = err.Error()
		pweekRes.Status = "E"
		return pdailydet, pweekRes, err
	}

	rows, err := db.Query(
		`SELECT 
			COALESCE(SUM(CASE WHEN curdate() = DATE(bill_date) THEN net_price ELSE 0 END), 0) AS net_price,
			mbm.login_id,
			(SELECT user_id FROM st832_medapp_login ml WHERE ml.login_id = mbm.login_id) AS user_id
		FROM 
			st832_medapp_bill_master mbm
		GROUP BY 
			mbm.login_id`)
	if err != nil {
		pweekRes.Status = "E"
		pweekRes.ErrMsg = err.Error()
		return pdailydet, pweekRes, err
	}

	// var d []SalesManSales
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pdailydet.Net, &pdailydet.Loginid, &pdailydet.Userid)
		if err != nil {
			pweekRes.Status = "E"
			pweekRes.ErrMsg = err.Error()
			return pdailydet, pweekRes, err
		}
		pweekRes.DailyArr = append(pweekRes.DailyArr, pdailydet)
	}
    
	pweekRes.ErrMsg = ""
	pweekRes.Status = "S"
	// pweekRes.DailyArr = d
	return pdailydet, pweekRes, nil
}

func ThisMonthSalesMethod(pthismonthdet ThisMonthSales, pweekRes WeekResponse) (ThisMonthSales, WeekResponse, error) {

	// db,err:= DBConnect.GORMDBConnect() 

	// if err!=nil{
	// 	pweekRes.ErrMsg = err.Error()
	// 	pweekRes.Status = "E"
	// 	return pthismonthdet, pweekRes, err

	// }

	// db.Table("medapp_login ml ").Select("ml.user_id,sum(mbm.net_price)").
	// Joins("medapp_bill_master mbm").Where("month(mbm.bill_date)=month() and year(mbm.bill_date)=year(curdate())").Group("ml.user_id").

	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		pweekRes.ErrMsg = err.Error()
		pweekRes.Status = "E"
		return pthismonthdet, pweekRes, err
	}

	rows, err := db.Query(
		`select ml.user_id  ,sum(mbm.net_price)  from 
st832_medapp_login ml  join
st832_medapp_bill_master mbm on ml.login_id = mbm.login_id 
WHERE month(mbm.bill_date) = MONTH(CURDATE())
  AND YEAR(mbm.bill_date) = YEAR(CURDATE())
  group  by ml.user_id ;`)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pthismonthdet.Userid,&pthismonthdet.Net)
		
		if err != nil {
			pweekRes.Status = "E"
			pweekRes.ErrMsg = err.Error()
			return pthismonthdet, pweekRes, err
		}
		pthismonthdet.RoundValues()
		 pweekRes.ThisMonthArr=append(pweekRes.ThisMonthArr, pthismonthdet)
	}
    log.Println(&pthismonthdet)
	return pthismonthdet,pweekRes,nil
}


func (ws *WeeklySales) RoundValues() {
	ws.Mon = math.Round(ws.Mon)
	ws.Tue = math.Round(ws.Tue)
	ws.Wed = math.Round(ws.Wed)
	ws.Thu = math.Round(ws.Thu)
	ws.Fri = math.Round(ws.Fri)
	ws.Sat = math.Round(ws.Sat)
	ws.Sun = math.Round(ws.Sun)
}


func(md *ThisMonthSales)RoundValues(){
	md.Net=math.Round(md.Net)
}