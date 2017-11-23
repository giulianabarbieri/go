package service

import "github.com/go/src/domain"
import "fmt"

var allTweets map[string][]*domain.Tweet = make(map[string][]*domain.Tweet)
var lastTweet *domain.Tweet

func InitializeService() {
	allTweets = make(map[string][]*domain.Tweet)
}

func GetTweets() []*domain.Tweet {
	allTweetsInSlice := make([]*domain.Tweet, 0)
	for _, element := range allTweets {
		//element es una lista de tweets. _ es el usuario
		for _, tweet := range element {
			allTweetsInSlice = append(allTweetsInSlice, tweet)
		}
	}
	return allTweetsInSlice
}

//la estructura ES el tipo
func PublishTweet(newTweet *domain.Tweet) (int, error) {

	if newTweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(newTweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	allTweets[newTweet.User] = append(allTweets[newTweet.User], newTweet)
	lastTweet = newTweet
	return newTweet.Id, nil
}

func GetTweet() *domain.Tweet {
	if len(allTweets) == 0 {
		return nil //HACER ESTO DE UN TEST
	}
	return lastTweet
}

func CleanTweet() {
	//FIJARSE DE NO BORRAR CUANDO NO HAY TWEETS
	//No se si hay que borrarlo del map tambien
	lastTweet = nil
}

func CleanTweets() {
	allTweets = make(map[string][]*domain.Tweet)
}

func GetTweetById(id int) *domain.Tweet {
	//Obtengo todos los tweets
	for _, element := range allTweets {
		//element es una lista de tweets. _ es el usuario
		for _, tweet := range element {
			//Por cada tweet de la lista element
			if tweet.Id == id {
				return tweet
			}
		}
	}

	return nil
}

func CountTweetsByUser(user string) int {
	userTweets, usuarioExiste := allTweets[user]
	if usuarioExiste {
		return len(userTweets)
	}
	return 0
}

func GetTweetsByUser(user string) []*domain.Tweet {
	return allTweets[user]
}
