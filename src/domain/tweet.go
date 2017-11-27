package domain

import (
	"time"
)

var nextId int = 0

//define una estructura
type Tweeter interface {
	PrintableTweet() string
	User() string
	Text() string
	Id() int
	Date() *time.Time
	//String() string //No hace falta agregarla aca porque es una interfaz aparte.
	//Hay que implementarla igual en las estructuras donde quiero imprimir
}
