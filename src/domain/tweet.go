package domain

var nextId int = 0

//define una estructura
type Tweeter interface {
	PrintableTweet() string
	//String() string //No hace falta agregarla aca porque es una interfaz aparte.
	//Hay que implementarla igual en las estructuras donde quiero imprimir
}
