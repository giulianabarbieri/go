package service

import (
	"fmt"
	"strings"

	"github.com/go/src/domain"
)

type tweetManager struct {
	userFollowing map[string][]string
	allTweets     map[string][]*domain.Tweet
	lastTweet     *domain.Tweet
	wordCounter   map[string]int
}

func (manager *tweetManager) GetTweets() []*domain.Tweet {
	allTweetsInSlice := make([]*domain.Tweet, 0)
	for _, element := range manager.allTweets {
		//element es una lista de tweets. _ es el usuario
		for _, tweet := range element {
			allTweetsInSlice = append(allTweetsInSlice, tweet)
		}
	}
	return allTweetsInSlice
}

//la estructura ES el tipo
func (manager *tweetManager) PublishTweet(newTweet *domain.Tweet) (int, error) {

	if newTweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(newTweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	manager.allTweets[newTweet.User] = append(manager.allTweets[newTweet.User], newTweet)
	manager.lastTweet = newTweet
	return newTweet.Id, nil
}

func (manager *tweetManager) CountTweetWords(tweet *domain.Tweet) {
	text := tweet.Text
	wordsList := strings.Fields(text) //te las guarda en una lista de strings PREZIOZO
	for index, element := range wordsList {
		value, ok := manager.CountTweetWords[element]
	}
}

func (manager *tweetManager) GetTweet() *domain.Tweet {
	if len(manager.allTweets) == 0 {
		return nil //HACER ESTO DE UN TEST
	}
	return manager.lastTweet
}

func (manager *tweetManager) CleanTweet() {
	//TODO Borrar el tweet del map
	if manager.lastTweet != nil {
		tweets := manager.GetTweetsByUser(manager.lastTweet.User)
		for idx, tweet := range tweets {
			if tweet == manager.lastTweet {
				manager.allTweets[manager.lastTweet.User] = append(manager.allTweets[manager.lastTweet.User][:idx], manager.allTweets[manager.lastTweet.User][idx+1:]...)
			}
		}
	}
	manager.lastTweet = nil
}

func (manager *tweetManager) CleanTweets() {
	manager.allTweets = make(map[string][]*domain.Tweet)
}

func (manager *tweetManager) GetTweetById(id int) *domain.Tweet {
	//Obtengo todos los tweets
	for _, element := range manager.allTweets {
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

func (manager *tweetManager) CountTweetsByUser(user string) int {
	userTweets, usuarioExiste := manager.allTweets[user]
	if usuarioExiste {
		return len(userTweets)
	}
	return 0
}

func (manager *tweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	return manager.allTweets[user]
}

func (manager *tweetManager) Follow(user1, user2 string) {
	userFollowed, ok := manager.userFollowing[user1]
	if !ok {
		userFollowed = make([]string, 0)
	}
	manager.userFollowing[user1] = append(userFollowed, user2)
}

func (manager *tweetManager) GetTimeLine(user string) []*domain.Tweet {
	followed := manager.userFollowing[user]
	sliceOfTweets := make([]*domain.Tweet, 0)
	for _, usuario := range followed {
		userTweets := manager.GetTweetsByUser(usuario)
		sliceOfTweets = append(sliceOfTweets, userTweets...) //copio y creo un array nuevo con lo que le agrego, LO TENGO QUE PEGAR AL VIEJO
		//lospuntos suspensivos son para decirle "como esto es una lista , quiero todos los elementos"
	}
	sliceOfTweets = append(sliceOfTweets, manager.GetTweetsByUser(user)...) //los mios tambien aparecen
	return sliceOfTweets
}

func NewTweetManager() tweetManager {
	return tweetManager{
		make(map[string][]string), //todo lo que esta luego del primer corchete es el tipo que almacena
		make(map[string][]*domain.Tweet),
		nil,
		make(map[string]int),
	}
}

func (manager *tweetManager) GetTrendingTopic() []string {

}
