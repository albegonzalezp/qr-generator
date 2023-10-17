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
	Subject  string `json:"subject"`
	Body     string `json:"body"`
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
		Subject:  os.Getenv("SUBJECT"),
		Body:     os.Getenv("BODY"),
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
	m.SetHeader("Subject", fmt.Sprintf("%s %s", qr.QrName, Config.Subject))
	m.SetBody("text/html", Config.Body)
	m.Attach(path)

	d := gomail.NewDialer(Config.Host, Config.Port, Config.Username, Config.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println("Could not send email: ", err)
		return fmt.Errorf("Could not send email: %v", err)
	}

	return nil
}
