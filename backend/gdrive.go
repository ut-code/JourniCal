package main

// WARNING: NOT WORKING!!
import (
	"context"

	"github.com/ut-code/JourniCal/backend/helper"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GDriveSample(ctx context.Context, config oauth2.Config, token *oauth2.Token, id string) {
	client := config.Client(ctx, token)
	dSvc, err := drive.NewService(ctx, option.WithHTTPClient(client))
	helper.ErrorLog(err)
	file, err := GetContent(*dSvc, id)
	helper.ErrorLog(err)
	link := file.WebContentLink
	helper.WriteFile("./sample.link", []byte(link))
}

func GetContent(s drive.Service, id string) (*drive.File, error) {
	b, err := s.Files.Get(id).Do()
	return b, err
}
