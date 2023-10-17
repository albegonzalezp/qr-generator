package main

import "testing"

func TestNewQr(t *testing.T) {
	// qr request body
	qrReq := &QrRequest{
		QrName:    "test",
		Link:      "https://www.google.com",
		ImageType: "png",
	}

	qr, err := CreateQR(qrReq)
	if err != nil {
		t.Error(err)
	}

	if qr.QrName != qrReq.QrName {
		t.Errorf("Expected qr name to be %s, got %s", qrReq.QrName, qr.QrName)
	}

	if qr.Link != qrReq.Link {
		t.Errorf("Expected link to be %s, got %s", qrReq.Link, qr.Link)
	}

	if qr.ImageType != qrReq.ImageType {
		t.Errorf("Expected image type to be %s, got %s", qrReq.ImageType, qr.ImageType)
	}

	if qr.ImagePath != "qr_images/test.png" {
		t.Errorf("Expected image path to be qr_images/test.png, got %s", qr.ImagePath)
	}

}
