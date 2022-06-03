package connectdatabase

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

//var dsn = "sql6497182:2j7eSP7MCU@tcp(sql6.freemysqlhosting.net:3306)/sql6497182?charset=utf8mb4&parseTime=True&loc=Local"

var (
	host     string = "ec2-54-86-224-85.compute-1.amazonaws.com"
	port     string = "5432"
	username string = "ubargppvqbrgnh"
	password string = "5a993e8d8add6ae7cc0b7768b903b499700e17e5e2b967d6cd50f90ee75678d4"
	database string = "ds66578msdrmn"
)

var dbConnect *gorm.DB

func DBConn() (db *gorm.DB) {
	return dbConnect
}

func InitConnect() {
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// //defer sqlDB.Close()
	// sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//sqlDB.SetConnMaxLifetime(1 * time.Second)
	dbConnect = db
}
