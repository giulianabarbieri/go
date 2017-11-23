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
			_, err := service.PublishTweet(tweet)
			if err != nil {
				c.Print("An error has ocurred, tweet not published")
				return
			}

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

	shell.AddCmd(&ishell.Cmd{
		Name: "ShowAllTweets",
		Help: "te muestra TODO ",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print(service.GetTweets())

			return
		}, //publicar tweet
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "CountUserTweets",
		Help: "Cuenta los tweets de un usuario ",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()

			c.Print(service.CountTweetsByUser(username))

			return
		}, //publicar tweet
	})
	shell.Run()
}
