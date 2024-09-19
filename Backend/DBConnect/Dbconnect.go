package DBConnect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func LocalDBConnect() (*sql.DB, error) {

	log.Println("LocalDBConnect+")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "ST832", "Best@123", "192.168.2.5", 3306, "training")

	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Println("Open connection failed:", err.Error())
		return db, err
	}

	log.Println("Localdbconnect-")
	return db, nil
}
