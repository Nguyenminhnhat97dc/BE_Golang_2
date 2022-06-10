package controllers

import (
	connectdatabase "Nguyenminhnhat97dc/BE_Golang/connectDatabase"
	"Nguyenminhnhat97dc/BE_Golang/create_database/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	//"gorm.io/driver/postgres"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
)

/* var (
	host     string = "ec2-54-86-224-85.compute-1.amazonaws.com"
	port     string = "5432"
	username string = "ubargppvqbrgnh"
	password string = "5a993e8d8add6ae7cc0b7768b903b499700e17e5e2b967d6cd50f90ee75678d4"
	database string = "ds66578msdrmn"
)
*/
//var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// select * from Services
func FindServices(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	var services []models.Services
	if err := dbConnect.Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})

	}
}

// SELECT * FROM users LIMIT 4;
func LimitServices(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	var services []models.Services
	count := c.Param("count")
	number, _ := strconv.Atoi(count)
	if err := dbConnect.Limit(number).Find(&services).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": services})
	}

}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

//
func AddRequirementCustomer(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type CheckCustomer struct {
		Name         string
		Address      string
		Phone        string
		NameServices string
		DayStart     string
		TimeStart    string
	}
	var checkCustomer CheckCustomer
	var Customer models.Customer
	var Requirement models.RequirementsCustomer

	// convert string
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newStr := buf.String()
	// convert Json
	res, err := PrettyString(newStr)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(res), &checkCustomer)
	abc := checkCustomer.Name
	if err := dbConnect.Where("name_customer = ? AND  address_customer = ?", abc, checkCustomer.Address).First(&Customer).Error; err != nil {
		NewCustomer := models.Customer{
			NameCustomer:    checkCustomer.Name,
			AddressCustomer: checkCustomer.Address,
			PhoneCustomer:   checkCustomer.Phone,
		}
		fmt.Println("THEM KH", NewCustomer)
		if err := dbConnect.Create(&NewCustomer).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "Không  insert Khách Hàng được"})
		} else {
			dbConnect.Where("name_customer = ? AND  address_customer = ?", abc, checkCustomer.Address).First(&Customer)
			NewRequirement := models.RequirementsCustomer{
				CustomerID:   Customer.ID,
				NameServices: checkCustomer.NameServices,
				DayStart:     checkCustomer.DayStart,
				TimeStart:    checkCustomer.TimeStart,
			}
			fmt.Println(">>requirement", NewRequirement)
			if err := dbConnect.Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
				if err := dbConnect.Create(&NewRequirement).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "create - không Insert yêu cầu khách hàng insert được"})
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
				}
			} else {
				if err := dbConnect.Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
					c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
				} else {
					c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
				}
			}
		}
	} else {
		NewRequirement := models.RequirementsCustomer{
			CustomerID:   Customer.ID,
			NameServices: checkCustomer.NameServices,
			DayStart:     checkCustomer.DayStart,
			TimeStart:    checkCustomer.TimeStart,
		}
		if err := dbConnect.Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).First(&Requirement).Error; err != nil {
			if err := dbConnect.Create(&NewRequirement).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "không Insert yêu cầu khách hàng insert được"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Insert yêu cầu khách hàng thành công"})
			}
		} else {
			if err := dbConnect.Model(&NewRequirement).Where("customer_id = ? AND day_start = ? AND time_start = ? ", NewRequirement.CustomerID, NewRequirement.DayStart, NewRequirement.TimeStart).Update("name_services", NewRequirement.NameServices).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Không Update được"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
			}
		}
	}
}

func ServiceProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	dbConnect := connectdatabase.DBConn()
	//Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}

	var data struct {
		Id string `json:"Id"`
	}

	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type GetServices struct {
		ServicesId   uint
		NameServices string
		Price        string
		ProviderId   uint
	}
	defer ws.Close()
	for {

		var getServices []GetServices
		dbConnect.Raw("SELECT services_of_providers.services_id,services.name_services, services_of_providers.price, services_of_providers.provider_id FROM"+
			" services_of_providers LEFT JOIN services on services_of_providers.services_id = services.id"+
			" WHERE services_of_providers.provider_id = ?", data.Id).Scan(&getServices)

		err = ws.WriteJSON(getServices)
		if err != nil {
			log.Println("Lỗi ở đây nè error write json: " + err.Error())
			break
		}

		time.Sleep(1 * time.Second)
	}

	/* type CheckProvider struct {
		Id string
	}
	type GetServices struct {
		ServicesId   uint
		NameServices string
		Price        string
		ProviderId   uint
	}
	var checkProvider CheckProvider
	var getServices []GetServices
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkProvider)
	if err := database.DBConn().Raw("SELECT services_of_providers.services_id,services.name_services, services_of_providers.price, services_of_providers.provider_id FROM"+
		" `services_of_providers` LEFT JOIN services on services_of_providers.services_id = services.id"+
		" WHERE services_of_providers.provider_id = ?", checkProvider.Id).Scan(&getServices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "Không tìm thấy"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"result": getServices})
		return
	} */
}

func AddServiceProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type GetServicesOfProvider struct {
		ServicesId uint
		ProviderId uint
		Price      int64
	}
	var getServicesOfProvider GetServicesOfProvider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &getServicesOfProvider)
	AddNewServiceProvider := models.ServicesOfProvider{
		ServicesId: getServicesOfProvider.ServicesId,
		ProviderID: getServicesOfProvider.ProviderId,
		Price:      getServicesOfProvider.Price,
	}
	if err := dbConnect.Create(&AddNewServiceProvider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})

	} else {
		c.JSON(http.StatusOK, gin.H{"result": "true"})
	}
}

func RequirementsCustomer(c *gin.Context) {
	dbConnect := connectdatabase.DBConn()
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	type InformationRequirementsCustomer struct {
		Id              uint
		NameServices    string
		DayStart        string
		TimeStart       string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	defer ws.Close()
	for {
		var informationRequirementsCustomer []InformationRequirementsCustomer
		if err := dbConnect.Raw(
			"SELECT requirements_customers.id,requirements_customers.name_services,requirements_customers.day_start,requirements_customers.time_start,customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM requirements_customers,customers WHERE requirements_customers.customer_id = customers.id and requirements_customers.status = ?", 0).Scan(&informationRequirementsCustomer).Error; err != nil {
			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		} else {
			if informationRequirementsCustomer != nil {
				err = ws.WriteJSON(informationRequirementsCustomer)
				if err != nil {
					log.Println("error write json: " + err.Error())
					break
				}
			} else {
				err = ws.WriteJSON("False")
				if err != nil {
					log.Println("error write json: " + err.Error())
					break
				}
			}

		}

		time.Sleep(500 * time.Millisecond)
	}

}

func TodoList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	dbConnect := connectdatabase.DBConn()
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}

	type CheckProvider struct {
		Id string
	}
	var checkProvider CheckProvider

	err = ws.ReadJSON(&checkProvider)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type TodoList struct {
		Id              uint
		NameServices    string
		Status          string
		DayStart        string
		TimeStart       string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	defer ws.Close()
	for {
		var todoList []TodoList
		if err := dbConnect.Raw(
			"SELECT requirements_customers.id,requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,"+
				" customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM to_do_lists,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
				" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = 0 and providers.id = ?", checkProvider.Id).Scan(&todoList).Error; err != nil {

			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())

				break
			}
		} else {
			if len(todoList) > 0 {

				err = ws.WriteJSON(todoList)
				if err != nil {
					log.Println("error write json: " + err.Error())

					break
				}
			} else {
				err = ws.WriteJSON("Bạn không có việc cần làm")
				if err != nil {
					log.Println("error write json: " + err.Error())

					break
				}
			}

		}
		time.Sleep(500 * time.Millisecond)

	}

	/* type CheckProvider struct {
		Id     string
		Status int
		PaginationStart uint
		PaginationEnd   uint
	}
	var checkProvider CheckProvider
	type TodoList struct {
		Id              uint
		NameServices    string
		Status          string
		DayStart        string
		TimeStart       string
		DayEnd          string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}

	var pagination Pagination
	var todoList []TodoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkProvider)
	if err := dbConnect.Raw(
		"SELECT to_do_lists.id,requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,"+
			" customers.name_customer,customers.address_customer,customers.phone_customer"+
			" FROM `to_do_lists`,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
			" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = ? and providers.id = ? LIMIT ?,?", checkProvider.Status, checkProvider.Id, checkProvider.PaginationStart, checkProvider.PaginationEnd).Scan(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		if len(todoList) > 0 {
			c.JSON(http.StatusOK, gin.H{"result": todoList})
		} else {
			if checkProvider.Status == 1 {
				c.JSON(http.StatusOK, gin.H{"result": []string{"Bạn không có lịch sử công việc"}})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"result": []string{"Bạn không có việc cần làm"}})
				return
			}
		}
	} */
}

func Loggin(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type CheckLoggin struct {
		User     string
		Password string
	}
	var checkLoggin CheckLoggin
	var informationLoggin models.User
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkLoggin)
	if err := dbConnect.Where("user_name = ? and password = ?", checkLoggin.User, checkLoggin.Password).First(&informationLoggin).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationLoggin})
	}
}

func FindProviderID(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type CheckID struct {
		Id uint
	}
	var checkID CheckID
	var informationProvider models.Provider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkID)
	if err := dbConnect.First(&informationProvider, "id=?", checkID.Id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": informationProvider})
	}
}

func FindPriceOfServices(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type CheckID struct {
		Id string
	}
	type Price struct {
		NameServices string
		Price        string
		/* Name         string */
	}
	var checkID CheckID
	var price []Price
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkID)
	if err := dbConnect.Raw(
		"SELECT services.name_services, services_of_providers.price from"+
			" services_of_providers,services,providers WHERE services_of_providers.services_id = services.id and"+
			" services_of_providers.provider_id = providers.id and providers.id = ?", checkID.Id).Scan(&price).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": price})
	}
}

func AddPrice(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type CheckInformation struct {
		Id           uint
		NameServices string
		Price        int64
	}
	var services models.Services
	var checkInformation CheckInformation
	var servicesOfProvider models.ServicesOfProvider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &checkInformation)
	if err := dbConnect.First(&services, "name_services=?", &checkInformation.NameServices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
		return
	} else {
		fmt.Println(">>>>>>", services.ID, checkInformation.Price)
		if err := dbConnect.Where("services_id=? and provider_id=?", services.ID, checkInformation.Id).First(&servicesOfProvider).Error; err != nil {
			NewServicesOfProvider := models.ServicesOfProvider{
				ServicesId: services.ID,
				ProviderID: checkInformation.Id,
				Price:      checkInformation.Price,
			}
			if err := dbConnect.Create(&NewServicesOfProvider).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "False"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "True"})
			}
		} else {
			if err := dbConnect.Model(&servicesOfProvider).Where("services_id=?", services.ID).Update("price", checkInformation.Price).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "Update thất bại"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": "Update thành công"})
			}
		}
	}
}

func AddTodoList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	var requirementcustomer models.RequirementsCustomer
	var addTodoList models.ToDoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &addTodoList)
	if err := dbConnect.First(&requirementcustomer, "id=?", addTodoList.RequirementsCustomerID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "Không thấy"})
	} else {
		if requirementcustomer.Status == 0 {
			if err := dbConnect.Create(&addTodoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"result": "False"})
			} else {
				dbConnect.Model(&requirementcustomer).Where("id=?", addTodoList.RequirementsCustomerID).Update("status", 1)
				c.JSON(http.StatusOK, gin.H{"result": "True"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "Công việc đã được người khác nhận"})
		}
	}
}

func CountPaginationRequirement(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	dbConnect := connectdatabase.DBConn()
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	type Count struct {
		Count uint
	}
	type CheckStatus struct {
		Status uint
	}
	var checkStatus CheckStatus
	err = ws.ReadJSON(&checkStatus)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	defer ws.Close()
	for {
		var count Count
		if err := dbConnect.Raw("SELECT COUNT(requirements_customers.id) AS "+"Count"+" FROM requirements_customers WHERE requirements_customers.status = ?", checkStatus.Status).Scan(&count).Error; err != nil {
			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())

				break
			}

		} else {
			err = ws.WriteJSON(count)
			if err != nil {
				log.Println("error write json: " + err.Error())

				break
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}
func CountPaginationToDoList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	dbConnect := connectdatabase.DBConn()

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	type Count struct {
		Count uint
	}
	type CheckStatus struct {
		Status     uint
		ProviderId uint
	}

	var checkStatus CheckStatus
	err = ws.ReadJSON(&checkStatus)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	defer ws.Close()
	for {
		var count Count
		if err := dbConnect.Raw("SELECT COUNT(to_do_lists.id) AS "+"Count"+" FROM to_do_lists WHERE to_do_lists.status = ? AND to_do_lists.provider_id = ?", checkStatus.Status, checkStatus.ProviderId).Scan(&count).Error; err != nil {
			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		} else {
			err = ws.WriteJSON(count)
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func CountPaginationHistory(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	dbConnect := connectdatabase.DBConn()

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	type Count struct {
		Count uint
	}
	type CheckStatus struct {
		Status     uint
		ProviderId uint
	}

	var checkStatus CheckStatus
	err = ws.ReadJSON(&checkStatus)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	defer ws.Close()
	for {
		var count Count
		if err := dbConnect.Raw("SELECT COUNT(to_do_lists.id) AS "+"Count"+" FROM to_do_lists WHERE to_do_lists.status = ? AND to_do_lists.provider_id = ?", checkStatus.Status, checkStatus.ProviderId).Scan(&count).Error; err != nil {
			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		} else {
			err = ws.WriteJSON(count)
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func UpdateTodoList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	}
	*/
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()

	type GetInformation struct {
		ProviderId            uint
		RequirementCustomerId uint
		InformationServices   string `gorm:"type: json"`
	}

	var getInformation GetInformation
	var todoList models.ToDoList
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	/* res, err := PrettyString(newString)
	if err != nil {
		log.Fatal(err)
	} */
	fmt.Println(newString)
	json.Unmarshal([]byte(newString), &getInformation)
	addhistory := models.HistoryList{
		ProviderID:             getInformation.ProviderId,
		RequirementsCustomerID: getInformation.RequirementCustomerId,
		InformationServices:    getInformation.InformationServices,
	}
	fmt.Println("><<<<<<<<<<<<<<<<<<<<<", getInformation)
	if err := dbConnect.Select("provider_id", "requirements_customer_id", "information_services").Create(&addhistory).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		if err := dbConnect.Raw("UPDATE to_do_lists set status = 1 where requirements_customer_id = ? and provider_id = ?", getInformation.RequirementCustomerId, getInformation.ProviderId).Scan(&todoList).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "update Todolist False"})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "True"})
		}
	}
}

func DeleteServicesProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type Information struct {
		ProviderId uint
		ServicesId uint
	}
	var information Information
	var deleteservices models.ServicesOfProvider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &information)
	if err := dbConnect.Raw("DELETE FROM services_of_providers WHERE provider_id = ? and services_id = ?", information.ProviderId, information.ServicesId).Scan(&deleteservices).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "True"})
	}
}

func HistoryList(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	dbConnect := connectdatabase.DBConn()
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}

	type CheckProvider struct {
		Id string
	}
	var checkProvider CheckProvider

	err = ws.ReadJSON(&checkProvider)

	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}
	type TodoList struct {
		NameServices    string
		Status          string
		DayStart        string
		TimeStart       string
		DayEnd          string
		NameCustomer    string
		AddressCustomer string
		PhoneCustomer   string
	}
	var todoList []TodoList
	defer ws.Close()
	for {
		if err := dbConnect.Raw(
			"SELECT requirements_customers.name_services,to_do_lists.status,requirements_customers.day_start,requirements_customers.time_start,to_do_lists.day_end,"+
				" customers.name_customer,customers.address_customer,customers.phone_customer"+
				" FROM to_do_lists,requirements_customers,customers,providers WHERE to_do_lists.requirements_customer_id = requirements_customers.id and"+
				" requirements_customers.customer_id = customers.id and to_do_lists.provider_id = providers.id and to_do_lists.status = 1 and providers.id = ?", checkProvider.Id).Scan(&todoList).Error; err != nil {

			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		} else {
			if len(todoList) > 0 {

				err = ws.WriteJSON(todoList)
				if err != nil {
					log.Println("error write json: " + err.Error())
					break
				}
			} else {
				err = ws.WriteJSON("Bạn không có lịch sử công việc")
				if err != nil {
					log.Println("error write json: " + err.Error())
					break
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func UpdateInformationProvider(c *gin.Context) {
	/* dbConnect, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */

	/* dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	} else {
		fmt.Println("connect Successfull.")
	} */
	/* sqlDB, err := dbConnect.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close() */
	dbConnect := connectdatabase.DBConn()
	type GetInformation struct {
		ProviderId uint
		Name       string
		Address    string
		CCCD       string
		Phone      string
		introduce  string
	}
	var getInformation GetInformation
	var updateProvider models.Provider
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	newString := buf.String()
	json.Unmarshal([]byte(newString), &getInformation)
	if err := dbConnect.Raw("UPDATE providers set name= ?, address= ?, cccd= ?, phone= ?, introduce= ? WHERE id = ?", getInformation.Name, getInformation.Address, getInformation.Phone, getInformation.CCCD, getInformation.introduce, getInformation.ProviderId).Scan(&updateProvider).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "False"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": getInformation})
	}
}

func GetHistory(c *gin.Context) {
	dbConnect := connectdatabase.DBConn()
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}
	type Information struct {
		Id uint
	}
	type ResultInformation struct {
		NameCustomer        string
		AddressCustomer     string
		PhoneCustomer       string
		DayStart            string
		TimeStart           string
		DayEnd              string
		InformationServices string `gorm:"type: json"`
	}
	var getInformation Information
	var resultInformation []ResultInformation
	err = ws.ReadJSON(&getInformation)
	if err != nil {
		log.Println("error read Json aksjdaslkdhkaj")
		log.Fatal(err)
	}
	/* 	buf := new(bytes.Buffer)
	   	buf.ReadFrom(c.Request.Body)
	   	newString := buf.String()
	   	json.Unmarshal([]byte(newString), &getInformation) */
	for {
		if err := dbConnect.Raw("select customers.name_customer,customers.address_customer,customers.phone_customer,"+
			" requirements_customers.day_start,requirements_customers.time_start,history_lists.day_end,"+
			" history_lists.information_services from history_lists,requirements_customers,customers where provider_id = ?"+
			" and requirements_customers.customer_id = customers.id and requirements_customers.id = history_lists.requirements_customer_id", getInformation.Id).Scan(&resultInformation).Error; err != nil {
			err = ws.WriteJSON("False")
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		} else {
			err = ws.WriteJSON(resultInformation)
			if err != nil {
				log.Println("error write json: " + err.Error())
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}
