package main

import (
	"github.com/andikanugraha11/go-gin-gorm-restapi/database"
	"github.com/andikanugraha11/go-gin-gorm-restapi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitMysqlDB()
	DBConn := &controllers.DBConn{DB: db}

	router := gin.Default()

	// router.GET("/hello", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello World!")
	// })
	router.GET("/person", DBConn.GetPersons)
	router.GET("/person/:id", DBConn.GetPerson)
	router.POST("/person", DBConn.CreatePerson)
	router.PUT("/person", DBConn.UpdatePerson)
	router.DELETE("/person/:id", DBConn.DeletePerson)
	
	router.Run(":8080")
}
