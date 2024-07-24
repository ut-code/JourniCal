package diary

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
out:
	for {
		events, err := calendar.GetEventsInRange(calendar.StaticService(), "primary", time.Now(), time.Now().AddDate(0, 3, 0))
		helper.PanicOn(err)
		for _, event := range events {
			lorem_ipsum[idx].eventid = event.Id
			if idx == len(lorem_ipsum)-1 {
				break out
			}
			idx++
		}
	}

	for _, journal := range lorem_ipsum {
		d := &Diary{
			CreatorID: user.StaticUser.ID,
			Date:      time.Now(),
			Title:     journal.title,
			Content:   journal.content,
			// EventID: journal.eventid,
		}
		err := CreateUnchecked(db, d)
		if err != nil {
			log.Fatalln("Failed to create diary: ", err)
		}
	}
	fmt.Println("[log] Prefilled db with", len(lorem_ipsum), "entries of diaries")
}
