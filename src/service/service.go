package service

import (
	"fmt"
	"strings"

	"github.com/go/src/domain"
)

//TweetManager estruutra del tweetManager
type TweetManager struct {
	userFollowing map[string][]string
	allTweets     map[string][]*domain.Tweet
	lastTweet     *domain.Tweet
	wordCounter   map[string]int
}

//GetTweets obtiene todos los tweets
func (manager *TweetManager) GetTweets() []*domain.Tweet {
	allTweetsInSlice := make([]*domain.Tweet, 0)
	for _, tweetList := range manager.allTweets {
		for _, tweet := range tweetList {
			allTweetsInSlice = append(allTweetsInSlice, tweet)
		}
	}
	return allTweetsInSlice
}

//PublishTweet publica el tweet
func (manager *TweetManager) PublishTweet(newTweet *domain.Tweet) (int, error) {

	if newTweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(newTweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	manager.addWordsToCounter(newTweet)
	manager.allTweets[newTweet.User] = append(manager.allTweets[newTweet.User], newTweet)
	manager.lastTweet = newTweet
	return newTweet.Id, nil
}

func (manager *TweetManager) addWordsToCounter(tweet *domain.Tweet) {
	text := tweet.Text
	wordsList := strings.Fields(text)

	for _, word := range wordsList {
		wordCount, _ := manager.wordCounter[word]
		manager.wordCounter[word] = wordCount + 1
	}
}

//GetTweet obtiene el ultimo tweet enviado. Nil si el ultimo se boroo o no hay tweets
func (manager *TweetManager) GetTweet() *domain.Tweet {
	return manager.lastTweet
}

//CleanTweet borra el ultimo tweet enviado
func (manager *TweetManager) CleanTweet() {
	//Testear bien que ser borre el tweet del map
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

//CleanTweets borra todos los tweets
func (manager *TweetManager) CleanTweets() {
	manager.allTweets = make(map[string][]*domain.Tweet)
}

//GetTweetByID obtiene el tweet con el id. Nil si no existe
func (manager *TweetManager) GetTweetByID(id int) *domain.Tweet {
	//Obtengo todos los tweets
	for _, userTweetList := range manager.allTweets {
		for _, tweet := range userTweetList {
			if tweet.Id == id {
				return tweet
			}
		}
	}

	return nil
}

//CountTweetsByUser cuenta los tweets del usario
func (manager *TweetManager) CountTweetsByUser(user string) int {
	userTweets, usuarioExiste := manager.allTweets[user]
	if usuarioExiste {
		return len(userTweets)
	}
	return 0
}

//GetTweetsByUser obtiene los tweets del usuario
func (manager *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	return manager.allTweets[user]
}

//Follow hace que un usuario siga a otro
func (manager *TweetManager) Follow(user1, user2 string) {
	usersFollowed, esSeguidorDeUser2 := manager.userFollowing[user1]
	if !esSeguidorDeUser2 {
		usersFollowed = make([]string, 0)
	}
	manager.userFollowing[user1] = append(usersFollowed, user2)
}

//GetTimeLine obtiene la timeline del user
func (manager *TweetManager) GetTimeLine(user string) []*domain.Tweet {
	followed := manager.userFollowing[user]
	sliceOfTweets := make([]*domain.Tweet, 0)

	for _, usuario := range followed {
		userTweets := manager.GetTweetsByUser(usuario)
		sliceOfTweets = append(sliceOfTweets, userTweets...)
		//copio y creo un array nuevo con lo que le agrego, LO TENGO QUE PEGAR AL VIEJO
		//lospuntos suspensivos son para decirle "como esto es una lista , quiero todos los elementos"
	}

	sliceOfTweets = append(sliceOfTweets, manager.GetTweetsByUser(user)...) //los mios tambien aparecen
	return sliceOfTweets
}

//NewTweetManager Crea un tweet manager
func NewTweetManager() TweetManager {
	return TweetManager{
		make(map[string][]string), //todo lo que esta luego del primer corchete es el tipo que almacena
		make(map[string][]*domain.Tweet),
		nil,
		make(map[string]int),
	}
}

//GetTrendingTopic obtiene las dos palabras mas usadas en todos los tweets
func (manager *TweetManager) GetTrendingTopic() []string {
	firstTopic := ""
	firstTopicCounter := 0
	secondTopic := ""
	secondTopicCounter := 0

	for word, wordCount := range manager.wordCounter {
		if wordCount >= firstTopicCounter {
			secondTopic = firstTopic
			secondTopicCounter = firstTopicCounter
			firstTopic = word
			firstTopicCounter = wordCount
		} else if wordCount >= secondTopicCounter {
			secondTopic = word
			secondTopicCounter = wordCount
		}
	}
	return []string{firstTopic, secondTopic}
}
