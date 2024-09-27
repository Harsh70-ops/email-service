package route

import (
	"email_service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	To      string `json:"to" binding:"required"`
	From    string `json:"from" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func SendEmail(c *gin.Context) {
	// binding the request with the struct
	var req EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//calling the smtp for sending email

	if err := services.SendEmails(req.To, req.From, req.Subject, req.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}
