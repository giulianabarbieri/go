package domain

import (
	"time"
)

//define una estructura
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *Tweet {
	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date,
	}
	return &tweet //el & para el puntero
}
