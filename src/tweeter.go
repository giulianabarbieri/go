package main

import (
	"github.com/abiosoft/ishell"
	"github.com/go/src/domain"
	"github.com/go/src/service"
)

func main() {
	shell := ishell.New() //new de la libreria
	shell.SetPrompt("Tweeter >>")
	shell.Print("Type 'help' to know commands \n")
	//a la consola interactiva le agrega un comando
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "publishes  a tweet",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()

			c.Print("Write your tweet:")
			message := c.ReadLine()

			tweet := domain.NewTweet(username, message)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		}, //publicar tweet
	})
	//muestra el ultimo tweet que se guardo
	shell.AddCmd(&ishell.Cmd{
		Name: "ShowTweet",
		Help: "Shows atweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Print(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "CleanTweet",
		Help: "Clean a atweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			service.CleanTweet()

			return
		},
	})
	shell.Run()
}
