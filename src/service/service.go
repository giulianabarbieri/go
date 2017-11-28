package service

import (
	"fmt"
	"strings"

	"github.com/go/src/domain"
)

//TweetManager estruutra del TweetManager
type TweetManager struct {
	userFollowing        map[string][]string
	allTweets            map[string][]domain.Tweet
	lastTweet            domain.Tweet
	wordCounter          map[string]int
	allDirectMessages    map[string][]domain.Tweet
	unreadDirectMessages map[string][]domain.Tweet
	favTweetsByUser      map[string][]domain.Tweet
	pluginList           []domain.TweetPlugin
	ChannelWrite         *ChannelTweetWriter
	lastPluginMessages   []string
}

//NewTweetManager Crea un Tweet manager
func NewTweetManager(channelWrite *ChannelTweetWriter) TweetManager {
	return TweetManager{
		make(map[string][]string), //todo lo que esta luego del primer corchete es el tipo que almacena
		make(map[string][]domain.Tweet),
		nil,
		make(map[string]int),
		make(map[string][]domain.Tweet),
		make(map[string][]domain.Tweet),
		make(map[string][]domain.Tweet),
		make([]domain.TweetPlugin, 0),
		channelWrite,
		make([]string, 0),
	}
}

//GetTweets obtiene todos los Tweets
func (manager *TweetManager) GetTweets() []domain.Tweet {
	allTweetsInSlice := make([]domain.Tweet, 0)
	for _, TweetList := range manager.allTweets {
		for _, Tweet := range TweetList {
			allTweetsInSlice = append(allTweetsInSlice, Tweet)
		}
	}
	return allTweetsInSlice
}

//PublishTweet publica el Tweet
func (manager *TweetManager) PublishTweet(newTweet domain.Tweet, quit chan bool) (int, error) {

	if (newTweet).User() == "" {
		return 0, fmt.Errorf("user is required")
	}
	if (newTweet).Text() == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len((newTweet).Text()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	manager.addWordsToCounter(newTweet)
	manager.addTweetToUser(newTweet, (newTweet).User())
	manager.lastTweet = newTweet

	tweetChannel := make(chan domain.Tweet)
	go manager.ChannelWrite.WriteTweet(tweetChannel, quit)
	tweetChannel <- newTweet
	manager.runPlugins()
	close(tweetChannel)

	return (newTweet).Id(), nil
}

func (manager *TweetManager) runPlugins() {
	manager.lastPluginMessages = make([]string, 0)

	for _, plugin := range manager.pluginList {
		manager.lastPluginMessages = append(manager.lastPluginMessages, plugin.RunPlugin())
	}

}

func (manager *TweetManager) addWordsToCounter(tweet domain.Tweet) {
	text := (tweet).Text()
	wordsList := strings.Fields(text)

	for _, word := range wordsList {
		wordCount, _ := manager.wordCounter[word]
		manager.wordCounter[word] = wordCount + 1
	}
}

//GetTweet obtiene el ultimo Tweet enviado. Nil si el ultimo se boroo o no hay Tweets
func (manager *TweetManager) GetTweet() domain.Tweet {
	return manager.lastTweet
}

//Hacer que se reduzcan las  palabras del contador para el TT
//CleanTweet borra el ultimo Tweet enviado
func (manager *TweetManager) CleanTweet() {
	//Testear bien que ser borre el Tweet del map
	if manager.lastTweet != nil {
		Tweets := manager.GetTweetsByUser((manager.lastTweet).User())

		for idx, Tweet := range Tweets {
			if Tweet == manager.lastTweet {
				manager.allTweets[(manager.lastTweet).User()] = append(manager.allTweets[(manager.lastTweet).User()][:idx], manager.allTweets[(manager.lastTweet).User()][idx+1:]...)
			}
		}
	}
	manager.lastTweet = nil
}

//Hacer test de este metodo
//CleanTweets borra todos los Tweets
func (manager *TweetManager) CleanTweets() {
	manager.allTweets = make(map[string][]domain.Tweet)
	manager.lastTweet = nil
	manager.wordCounter = make(map[string]int)
	manager.userFollowing = make(map[string][]string)
}

//GetTweetByID obtiene el Tweet con el id. Nil si no existe
func (manager *TweetManager) GetTweetById(id int) domain.Tweet {
	//Obtengo todos los Tweets
	for _, userTweetList := range manager.allTweets {
		for _, tweet := range userTweetList {
			if (tweet).Id() == id {
				return tweet
			}
		}
	}

	return nil
}

//CountTweetsByUser cuenta los Tweets del usario
func (manager *TweetManager) CountTweetsByUser(user string) int {
	userTweets, usuarioExiste := manager.allTweets[user]
	if usuarioExiste {
		return len(userTweets)
	}
	return 0
}

//GetTweetsByUser obtiene los Tweets del usuario
func (manager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
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

//Rehacer el tests de este metodo
//GetTimeLine obtiene la timeline del user
func (manager *TweetManager) GetTimeLine(user string) []domain.Tweet {
	followed := manager.userFollowing[user]
	sliceOfTweets := make([]domain.Tweet, 0)

	for _, usuario := range followed {
		userTweets := manager.GetTweetsByUser(usuario)
		sliceOfTweets = append(sliceOfTweets, userTweets...)
		//copio y creo un array nuevo con lo que le agrego, LO TENGO QUE PEGAR AL VIEJO
		//lospuntos suspensivos son para decirle "como esto es una lista , quiero todos los elementos"
	}

	sliceOfTweets = append(sliceOfTweets, manager.GetTweetsByUser(user)...) //los mios tambien aparecen
	return sliceOfTweets
}

//GetTrendingTopic obtiene las dos palabras mas usadas en todos los Tweets
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

//SendDirectMessage Envia un mensaje directo al usuario receptor
func (manager *TweetManager) SendDirectMessage(Tweet domain.Tweet, receiver string) {
	manager.addTweetToMapStringKey(&manager.allDirectMessages, Tweet, receiver)
	manager.addTweetToMapStringKey(&manager.unreadDirectMessages, Tweet, receiver)
}

func (manager *TweetManager) addTweetToMapStringKey(mapa *map[string][]domain.Tweet, Tweet domain.Tweet, user string) {
	TweetList, _ := (*mapa)[user]
	(*mapa)[user] = append(TweetList, Tweet)
}

//GetAllDirectMessages obtiene todos los mensajes directos de un usuario
func (manager *TweetManager) GetAllDirectMessages(user string) []domain.Tweet {
	return manager.allDirectMessages[user]
}

//GetUnreadDirectMessages obtiene los mensajes sin leer de un usuario
func (manager *TweetManager) GetUnreadDirectMessages(user string) []domain.Tweet {
	return manager.unreadDirectMessages[user]
}

//ReadDirectMessage marca un mensaje como leido
func (manager *TweetManager) ReadDirectMessage(Tweet domain.Tweet, user string) {
	unreadDirectMessages, _ := manager.unreadDirectMessages[user]
	for index, directMessage := range unreadDirectMessages {
		if directMessage == Tweet {
			manager.unreadDirectMessages[user] = append(unreadDirectMessages[:index], unreadDirectMessages[index+1:]...)
		}
	}
}

//ReTweet hace que el Tweet aparezca dentro de mis Tweets
func (manager *TweetManager) ReTweet(Tweet domain.Tweet, user string) {
	manager.addTweetToUser(Tweet, user)
}

func (manager *TweetManager) addTweetToUser(Tweet domain.Tweet, user string) {
	manager.allTweets[user] = append(manager.allTweets[user], Tweet)
}

func (manager *TweetManager) FavTweet(Tweet domain.Tweet, user string) {
	manager.addTweetToMapStringKey(&manager.favTweetsByUser, Tweet, user)
}

func (manager *TweetManager) GetFavTweets(user string) []domain.Tweet {
	return manager.favTweetsByUser[user]
}

func (manager *TweetManager) AddPlugin(plugin domain.TweetPlugin) {
	manager.pluginList = append(manager.pluginList, plugin)
}

func (manager *TweetManager) GetPlugins() []domain.TweetPlugin {
	return manager.pluginList
}

func (manager *TweetManager) PluginMessages() []string {
	return manager.lastPluginMessages
}
