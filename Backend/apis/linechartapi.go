package apis

import (
	// "encoding/json"
	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"medapp/DBConnect"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type MonthlySales struct {
	Jan float64 `json:"jan"`
	Feb float64 `json:"feb"`
	Mar float64 `json:"mar"`
	Apr float64 `json:"apr"`
	May float64 `json:"may"`
	Jun float64 `json:"jun"`
	Jul float64 `json:"jul"`
	Aug float64 `json:"aug"`
	Sep float64 `json:"sep"`
	Oct float64 `json:"oct"`
	Nov float64 `json:"nov"`
	Dec float64 `json:"dec"`
}

type MonthResponse struct {
	Month_obj MonthlySales `json:"month_obj"`
	Status    string       `json:"status"`
	ErrMsg    string       `json:"err_msg"`
}

func FetchSaless(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	log.Println("FetchSaless(+)")

	if r.Method == "GET" {
		// Reset the response variable for each request
		var pmonthRes MonthResponse
		var pmonthDet MonthlySales
		var err1 error
		// Call method to fetch login history data.
		pmonthRes.Status = "S"
		pmonthRes, err1 = MonthMethod(pmonthDet, pmonthRes)

		if err1 != nil {
			pmonthRes.Status = "E"
			pmonthRes.ErrMsg = err1.Error()
		}
		jsonData, err := json.Marshal(pmonthRes)
		if err != nil {
			fmt.Fprintf(w, "Error encoding JSON response: %v", err)
			w.Write([]byte("Error while Marshaling : " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	}
}

func MonthMethod(pmonthDet MonthlySales, pmonthRes MonthResponse) (MonthResponse, error) {

	db, err := DBConnect.LocalDBConnect()
	if err != nil {
		pmonthRes.ErrMsg = err.Error()
		pmonthRes.Status = "E"
		return pmonthRes, err
	}
	defer db.Close()
	err1 := db.QueryRow(
		`SELECT 
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE()) AND YEAR(bill_date) = YEAR(CURDATE()) THEN net_price ELSE 0 END), 0), 2) AS current_month_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 1 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 1 MONTH) THEN net_price ELSE 0 END), 0), 2) AS previous_month_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 2 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 2 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_before_last_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 3 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 3 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_3_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 4 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 4 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_4_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 5 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 5 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_5_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 6 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 6 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_6_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 7 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 7 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_7_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 8 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 8 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_8_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 9 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 9 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_9_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 10 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 10 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_10_sales,
		ROUND(COALESCE(SUM(CASE WHEN MONTH(bill_date) = MONTH(CURDATE() - INTERVAL 11 MONTH) AND YEAR(bill_date) = YEAR(CURDATE() - INTERVAL 11 MONTH) THEN net_price ELSE 0 END), 0), 2) AS month_11_sales
	FROM 
		st832_medapp_bill_master;
	`).Scan(
		&pmonthDet.Jan,
		&pmonthDet.Feb,
		&pmonthDet.Mar,
		&pmonthDet.Apr,
		&pmonthDet.May,
		&pmonthDet.Jun,
		&pmonthDet.Jul,
		&pmonthDet.Aug,
		&pmonthDet.Sep,
		&pmonthDet.Oct,
		&pmonthDet.Nov,
		&pmonthDet.Dec,
	)

	if err1 != nil {
		pmonthRes.ErrMsg = err1.Error()
		pmonthRes.Status = "E"
		return pmonthRes, nil
	}
	pmonthDet.RoundValues()

	pmonthRes.Month_obj = pmonthDet
	pmonthRes.Status = "S"
	pmonthRes.ErrMsg = ""
	return pmonthRes, nil
}


func (ms *MonthlySales) RoundValues() {
	ms.Jan = math.Round(ms.Jan)
	ms.Feb = math.Round(ms.Feb)
	ms.Mar = math.Round(ms.Mar)
	ms.Apr = math.Round(ms.Apr)
	ms.May = math.Round(ms.May)
	ms.Jun = math.Round(ms.Jun)
	ms.Jul = math.Round(ms.Jul)
	ms.Aug = math.Round(ms.Aug)
	ms.Sep = math.Round(ms.Sep)
	ms.Oct = math.Round(ms.Oct)
	ms.Nov = math.Round(ms.Nov)
	ms.Dec = math.Round(ms.Dec)

	log.Println("obj")
	log.Println(ms)
}
