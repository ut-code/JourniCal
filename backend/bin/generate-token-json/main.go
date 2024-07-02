package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ut-code/JourniCal/backend/app/calendar"
	"golang.org/x/oauth2"
)

func main() {
	token, err := obtainTestingToken()
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create("./token.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(f).Encode(token)
	fmt.Println("successfully generated token.json.")
}

func obtainTestingToken() (*oauth2.Token, error) {
	config := calendar.ReadCredentials()
	authURL := config.AuthCodeURL("state-string", oauth2.AccessTypeOffline)
	fmt.Println("Go to this link and click ok: ", authURL)
	handler := handler{ch: make(chan string)}
	go http.ListenAndServe(":3000", handler)

	code := <-handler.ch
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

type handler struct{ ch chan string }

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ch <- r.URL.Query().Get("code")
	fmt.Fprintf(w, "accepted")
	close(h.ch)
}
