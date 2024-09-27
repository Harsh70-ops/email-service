package main

import (
	"email_service/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//running and testing the port 
	router.GET("/",func (c *gin.Context){
        c.JSON(http.StatusOK,"server is running")

	})

    router.POST("/email",route.SendEmail)
	router.Run(":8080")

}
