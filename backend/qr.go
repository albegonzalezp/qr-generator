package main

import (
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

type QrRequest struct {
	QrName    string `json:"qr_name"`
	Link      string `json:"link"`
	ImageType string `json:"image_type"`
}

type Qr struct {
	QrName    string `json:"qr_name"`
	Link      string `json:"link"`
	ImageType string `json:"image_type,omitempty"`
	ImagePath string `json:"image_path"`
}

func CreateQR(qrReq *QrRequest) (*Qr, error) {
	filename := fmt.Sprintf("./qr_images/%s.%s", qrReq.QrName, qrReq.ImageType)

	err := qrcode.WriteFile(qrReq.Link, qrcode.Medium, 256, filename)
	if err != nil {
		panic(err)
	}

	return &Qr{QrName: qrReq.QrName, Link: qrReq.Link, ImageType: qrReq.ImageType, ImagePath: filename}, nil
}

func DeleteQR(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("Could not delete qr image: %v", err)
	}
	return nil
}
