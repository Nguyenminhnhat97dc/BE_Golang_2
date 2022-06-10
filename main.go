package main

import (
	"Nguyenminhnhat97dc/BE_Golang/controllers"
	"fmt"
	"time"

	//"log"
	"net/http"
	//"time"
	"os"

	connectDatabase "Nguyenminhnhat97dc/BE_Golang/connectDatabase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/websocket"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("PORT : ", port)
	connectDatabase.InitConnect()
	//r := gin.New()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin, Authorization, Content-Type"},
		//AllowHeaders:     []string{"Origin, Authorization, Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
	//r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Data": "hello world"})
	})

	//r.GET("/provider", controllers.FindProvider)
	r.POST("/provider/id", controllers.FindProviderID)
	r.GET("/services", controllers.FindServices)
	r.GET("/services/:count", controllers.LimitServices)
	r.POST("/requirement", controllers.AddRequirementCustomer)
	r.GET("/provider/services", controllers.ServiceProvider)
	r.GET("/servicesofprovider", controllers.AddServiceProvider)
	r.GET("/requirementcustomer", controllers.RequirementsCustomer)
	r.GET("/todolist", controllers.TodoList)
	//r.GET("history", controllers.HistoryList)
	r.POST("/loggin", controllers.Loggin)
	r.POST("/priceservices", controllers.FindPriceOfServices)
	r.POST("/addprice", controllers.AddPrice)
	r.POST("/addtodolist", controllers.AddTodoList)
	r.GET("/paginationrequirement", controllers.CountPaginationRequirement)
	r.GET("/paginationtodolist", controllers.CountPaginationToDoList)
	r.GET("/paginationhistory", controllers.CountPaginationHistory)
	r.POST("/add_historylist", controllers.UpdateTodoList)
	r.POST("/deleteservices", controllers.DeleteServicesProvider)
	r.POST("/updateprovider", controllers.UpdateInformationProvider)
	r.GET("/historyy", controllers.GetHistory)
	r.Run(":" + port)
}
