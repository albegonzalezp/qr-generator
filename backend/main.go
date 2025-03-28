package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load godot env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	Config = InitEmailConfig()

	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("../frontend/templates/*")

	// Serve static files .js and .css
	r.Static("../frontend/assets", "../frontend/assets")

	// Serve images
	r.Static("./qr_images", "./qr_images")

	r.GET("/", RenderIndex)
	r.POST("/get-qr", RenderQr)
	r.POST("/send-qr", SendQR)

	r.Run(":8080")
}
