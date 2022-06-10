package connectdatabase

import (
	"fmt"

	"gorm.io/driver/postgres"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

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

	/* db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/*
		type InformationRequirementsCustomer struct {
			Id              uint
			NameServices    string
			DayStart        string
			TimeStart       string
			NameCustomer    string
			AddressCustomer string
			PhoneCustomer   string
		}
		var informationRequirementsCustomer []InformationRequirementsCustomer
		if err := db.Raw(
			"SELECT requirements_customers.id,requirements_customers.name_services,requirements_customers.day_start,requirements_customers.time_start,customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM requirements_customers,customers WHERE requirements_customers.customer_id = customers.id and requirements_customers.status = ?", 0).Scan(&informationRequirementsCustomer).Error; err != nil {
			fmt.Println("Thất bại")
		} else {
			fmt.Println(informationRequirementsCustomer)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalln(err)
		} */
	// //defer sqlDB.Close()
	// sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	/* 	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(18) */

	//sqlDB.SetConnMaxLifetime(1 * time.Second)
	dbConnect = db
}
