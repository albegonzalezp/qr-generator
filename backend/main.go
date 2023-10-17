package main

import (
	"os"
	"path/filepath"

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
	templatesDir := os.Getenv("TEMPLATES_DIR")
	r.LoadHTMLGlob(filepath.Join(templatesDir, "*"))

	// Serve static files .js and .css
	staticDir := os.Getenv("STATIC_DIR")
	r.Static("../frontend/assets", staticDir)

	// Serve images
	imagesDir := os.Getenv("IMAGES_DIR")
	r.Static("./qr_images", imagesDir)

	r.GET("/", RenderIndex)
	r.POST("/get-qr", RenderQr)
	r.POST("/send-qr", SendQR)

	r.Run(":8080")
}
