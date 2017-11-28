package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/go/src/domain"
	"github.com/go/src/service"
)

func main() {
	shell := ishell.New() //new de la libreria
	shell.SetPrompt("Tweeter >>")
	shell.Print("Type 'help' to know commands \n")
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)
	tweetManager := service.NewTweetManager(tweetWriter)

	//PublishTextTweet
	shell.AddCmd(&ishell.Cmd{
		Name: "PublishTextTweet",
		Help: "publishes  a tweet",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()

			c.Print("Write your tweet:")
			message := c.ReadLine()

			var tweet domain.Tweet = domain.NewTextTweet(username, message)
			quit := make(chan bool)
			_, err := tweetManager.PublishTweet(tweet, quit)

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

			tweet := tweetManager.GetTweet()

			c.Print(tweet)

			return
		},
	})

	//Borra el ultimo tweet que se mando
	shell.AddCmd(&ishell.Cmd{
		Name: "CleanTweet",
		Help: "Clean a atweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweetManager.CleanTweet()

			return
		},
	})

	//Muestra todos los tweets
	shell.AddCmd(&ishell.Cmd{
		Name: "ShowAllTweets",
		Help: "te muestra TODO ",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print(tweetManager.GetTweet())

			return
		}, //publicar tweet
	})

	//Contar tweets de un usuario
	shell.AddCmd(&ishell.Cmd{
		Name: "CountUserTweets",
		Help: "Cuenta los tweets de un usuario ",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()

			c.Print(tweetManager.CountTweetsByUser(username))

			return
		}, //publicar tweet
	})

	//Obtener tweet por id
	shell.AddCmd(&ishell.Cmd{
		Name: "GetTweetById",
		Help: "Obtiene tweeter con su identificacion unica ",
		Func: func(c *ishell.Context) { //que queremos ejecutar cuando alquien escriba el comando

			defer c.ShowPrompt(true)

			c.Print("Write Id: ")
			idstr := c.ReadLine()
			id, _ := strconv.Atoi(idstr)

			c.Print(tweetManager.GetTweetById(id))

			return
		}, //publicar tweet
	})

	//Obtener tweets por usuario
	shell.AddCmd(&ishell.Cmd{
		Name: "GetTweetsByUser",
		Help: "Obtiene los tweets por usuario ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username: ")
			username := c.ReadLine()

			c.Print(tweetManager.GetTweetsByUser(username))

			return
		},
	})

	//Seguir a alguien
	shell.AddCmd(&ishell.Cmd{
		Name: "Follow",
		Help: "sigue a un usuario ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()
			c.Print("Write their username: ")
			usernameToFollow := c.ReadLine()

			tweetManager.Follow(username, usernameToFollow)
			c.Print("followed")

			return
		},
	})

	//Obtener timeline
	shell.AddCmd(&ishell.Cmd{
		Name: "GetTimeLine",
		Help: "obtiene el timeline de un usario ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()

			c.Print(tweetManager.GetTimeLine(username))

			return
		},
	})

	//Obtener trending topics
	shell.AddCmd(&ishell.Cmd{
		Name: "GetTrendingTopics",
		Help: "obtiene el trending topic ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print(tweetManager.GetTrendingTopic())

			return
		},
	})

	//Enviar mensaje directo
	shell.AddCmd(&ishell.Cmd{
		Name: "SendDirectMessage",
		Help: "envia un tweet como mensaje directo",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")
			username := c.ReadLine()

			c.Print("Write your message:")
			message := c.ReadLine()

			var tweet domain.Tweet = domain.NewTextTweet(username, message)

			c.Print("Write receiver of the message:")
			receiver := c.ReadLine()

			tweetManager.SendDirectMessage(tweet, receiver)

			c.Print("Message sent\n")

			return
		},
	})

	//Obtener mensajes directos
	shell.AddCmd(&ishell.Cmd{
		Name: "GetAllDirectMEssages",
		Help: "get all messages of user ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username:")
			username := c.ReadLine()

			c.Print(tweetManager.GetAllDirectMessages(username))

			return
		},
	})

	//Obtener mensajes directos sin leer
	shell.AddCmd(&ishell.Cmd{
		Name: "GetUnreadDirectMessages",
		Help: "get all unread messages of user ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username:")
			username := c.ReadLine()

			c.Print(tweetManager.GetUnreadDirectMessages(username))

			return
		},
	})

	//Marcar un mensaje como leido
	shell.AddCmd(&ishell.Cmd{
		Name: "ReadDirectMessage",
		Help: "get all messages of user ",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username:")
			username := c.ReadLine()

			c.Print("Write message id:")
			messageidstr := c.ReadLine()
			messageid, _ := strconv.Atoi(messageidstr)

			tweetManager.ReadDirectMessage(tweetManager.GetTweetById(messageid), username)

			return
		},
	})

	//Retweetear
	shell.AddCmd(&ishell.Cmd{
		Name: "Retweet",
		Help: "can retweet a tweetby id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username:")
			username := c.ReadLine()

			c.Print("Write message id:")
			messageidstr := c.ReadLine()
			messageid, _ := strconv.Atoi(messageidstr)

			tweetManager.ReTweet(tweetManager.GetTweetById(messageid), username)

			return
		},
	})

	//Marcar un tweet como favorito
	shell.AddCmd(&ishell.Cmd{
		Name: "Fav",
		Help: "make a tweetafav",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username:")
			username := c.ReadLine()

			c.Print("Write message id:")
			messageidstr := c.ReadLine()
			messageid, _ := strconv.Atoi(messageidstr)

			tweetManager.FavTweet(tweetManager.GetTweetById(messageid), username)
			c.Print("Tweet Faved")
			return
		},
	})

	//Obtener favoritos
	shell.AddCmd(&ishell.Cmd{
		Name: "GetFavs",
		Help: "get fav tweets of an user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write username:")
			username := c.ReadLine()

			c.Print(tweetManager.GetFavTweets(username))

			return
		},
	})

	shell.Run()
}
