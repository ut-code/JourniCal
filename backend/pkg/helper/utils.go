package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// meta functions for programmers, not for program
func UseLater(_ ...any) {
}

func TODO(text ...string) {
	log.Fatal("TODO() was called: " + strings.Join(text, ", "))
}

// log.Fatalln's if err != nil. give err and describe what is expected to have happened. no ":" needed.
func ErrorLog(err error, msgs ...string) {
	if err != nil {
		log.Fatalln(": "+strings.Join(msgs, " "), err)
	}
}

func PanicIf(err error, msg ...string) {
	if err != nil {
		message := "helper.PanicIf called on error: " + fmt.Sprint(err)
		message += strings.Join(msg, " ")
		log.Fatalln(message)
	}
}

func WriteFile(path string, b []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}

func ReadFile(path string) (content string, err error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ToJSON[T any](v T) string {
	b, err := json.Marshal(v)
	ErrorLog(err)
	return string(b)
}
