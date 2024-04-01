package main

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

func GDriveSample(ctx context.Context, token oauth2.Token) {
	dSvc, err := drive.NewService(ctx)
	ErrorLog(err)
	url := "GET URL FROM SOMEWHERE"
	content, err := GetContent(*dSvc, url)
	ErrorLog(err)
	ErrorLog(writeFile("sample.jpg", content))
}

func GetContent(s drive.Service, url string) ([]byte, error) {

	return []byte{}, nil
}
