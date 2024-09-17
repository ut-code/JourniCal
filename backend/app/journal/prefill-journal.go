package journal

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ut-code/JourniCal/backend/app/calendar"
	"github.com/ut-code/JourniCal/backend/app/env/options"
	"github.com/ut-code/JourniCal/backend/app/user"
	"github.com/ut-code/JourniCal/backend/pkg/helper"
	"gorm.io/gorm"
)

func Prefill(db *gorm.DB) {
	if !options.STATIC_TOKEN {
		log.Fatalln("To do journal prefill, you must set TOKEN_SOURCE to either env or file")
	}

	rand.Shuffle(len(lorem_ipsum), func(i, j int) {
		tmp := lorem_ipsum[i]
		lorem_ipsum[i] = lorem_ipsum[j]
		lorem_ipsum[j] = tmp
	})
	var idx = 0
	var curr = time.Now().AddDate(0, -1, 0)
	var next time.Time
out:
	for {
		next = curr.AddDate(0, 1, 0)
		events, err := calendar.GetEventsInRange(calendar.StaticService(), "primary", curr, next)
		helper.PanicOn(err)
		for _, event := range events {
			lorem_ipsum[idx].eventid = event.Id
			if idx == len(lorem_ipsum)-1 {
				break out
			}
			idx++
		}
		curr = next
	}

	for _, journal := range lorem_ipsum {
		d := &Journal{
			CreatorID: user.StaticUser.ID,
			Date:      time.Now(),
			Title:     journal.title,
			Content:   journal.content,
			// EventID: journal.eventid,
		}
		err := CreateUnchecked(db, d)
		if err != nil {
			log.Fatalln("Failed to create journal: ", err)
		}
	}
	fmt.Println("[log] Prefilled db with", len(lorem_ipsum), "entries of journals")
}
