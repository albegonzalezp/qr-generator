package main

import (
	"github.com/gin-gonic/gin"
)

func RenderIndex(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func RenderQr(c *gin.Context) {
	// takes the data recolected from the form and creates a QR code
	var request struct {
		Email     string `json:"email"`
		QrName    string `json:"qr_name"`
		Link      string `json:"link"`
		ImageType string `json:"image_type"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"status":  "failed",
			"message": "Invalid parameters to create QR code",
		})
		return
	}

	qr, err := CreateQR(&QrRequest{QrName: request.QrName, Link: request.Link, ImageType: request.ImageType})
	if err != nil {
		c.JSON(500, gin.H{
			"error":   err.Error(),
			"status":  "failed",
			"message": "Could not create QR code",
		})
		return
	}

	// send the response to client with the necessary information to render the QR code
	c.JSON(200, gin.H{
		"status":     "success",
		"image_path": qr.ImagePath,
		"qr_name":    qr.QrName,
		"link":       qr.Link,
		"email":      request.Email,
	})
}

func SendQR(c *gin.Context) {
	// get the data necessary to send the email
	var requestData struct {
		Email     string `json:"email"`
		QrName    string `json:"qr_name"`
		Link      string `json:"link"`
		ImagePath string `json:"image_path"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := SendEmail(requestData.Email, &Qr{QrName: requestData.QrName, Link: requestData.Link, ImagePath: requestData.ImagePath}); err != nil {
		c.JSON(500, gin.H{
			"error":  err.Error(),
			"status": "failed",
		})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
