package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var Config *EmailConfig

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func InitEmailConfig() *EmailConfig {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	return &EmailConfig{
		Host:     os.Getenv("HOST"),
		Port:     port,
		Username: os.Getenv("FROM"),
		Password: os.Getenv("PASSWORD"),
	}
}

func SendEmail(email string, qr *Qr) error {
	// Parse the URL to get the path
	parsedURL, err := url.Parse(qr.ImagePath)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return err
	}

	// Get the path from "/qr_images" to the end
	path := fmt.Sprintf(".%s", parsedURL.Path)

	m := gomail.NewMessage()
	m.SetHeader("From", Config.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", fmt.Sprintf("%s", qr.QrName))
	m.SetBody("text/html", "Here is the <b>QR code</b> you requested.<br>Enjoy it!<br>Regards,<br>zetacoder")
	m.Attach(path)

	d := gomail.NewDialer(Config.Host, Config.Port, Config.Username, Config.Password)

	sendCloser, err := d.Dial()
	if err != nil {
		log.Println("Could not connect to the email server: ", err)
		return fmt.Errorf("Could not connect to the email server: %v", err)
	}

	defer sendCloser.Close()

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println("Could not send email: ", err)
		return fmt.Errorf("Could not send email: %v", err)
	}

	return nil
}
