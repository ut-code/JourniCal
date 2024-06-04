package main

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GDriveSample(ctx context.Context, config oauth2.Config, token *oauth2.Token, id string) {
	client := config.Client(ctx, token)
	dSvc, err := drive.NewService(ctx, option.WithHTTPClient(client))
	ErrorLog(err)
	file, err := GetContent(*dSvc, id)
	ErrorLog(err)
	link := file.WebContentLink
	writeFile("./sample.link", []byte(link))
}

func GetContent(s drive.Service, id string) (*drive.File, error) {
	b, err := s.Files.Get(id).Do()
	return b, err
}
