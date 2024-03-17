package main

import (
	"log"
	"strings"
)

// I'm telling the compiler that its fine to not use it for now, by "Using" it.
func Use(_ ...any) {
}

// log.Fatalln's if err != nil. give err and describe what is expected to have happened. no ":" needed.
func ErrorLog(err error, msgs ...string) {
	if err != nil {
		log.Fatalln(err, strings.Join(msgs, " ")+": ")
	}
}
