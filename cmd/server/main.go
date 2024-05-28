package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nirmalkatiyar/bitespeed/config"
	"github.com/nirmalkatiyar/bitespeed/controllers"
	"github.com/nirmalkatiyar/bitespeed/database"
)

func main() {
	fmt.Println("Hello World")
	g := gin.Default()
	db := database.InitDB()
	g.POST("/identify", controllers.IdentifyContact(db))
	port := config.GetPort()
	g.Run(":" + port)

}
