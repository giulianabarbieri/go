package domain

import (
	"time"
)

//define una estructura
type Tweet struct {
	User string
	Text string
	Date *time.Time //recibe un puntero
}

func NewTweet(user, text string) *Tweet { //pongo *tweet porque estoy en domain, donde quieras usar esto
	//y no estes en domain vas a tenner que importarlo
	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date, //es un puntero
	} //parece ser una variable local pero NO = magia
	return &tweet //el & para el puntero
}
