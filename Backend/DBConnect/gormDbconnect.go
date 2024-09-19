package DBConnect

import(
	"gorm.io/driver/mysql"
   "gorm.io/gorm"
    "fmt"
	"log"


)
func GORMDBConnect()( *gorm.DB,error){
	log.Println("GORMDBConnect+")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "ST832", "Best@123", "192.168.2.5", 3306, "training")

	db,err:= gorm.Open(mysql.Open(connString),&gorm.Config{})

	if err!=nil{
		return db,err

	}
	log.Println("GORMDBConnect-")
	
	return db,nil


}
