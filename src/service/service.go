package service

import (
	"fmt"
	"strings"

	"github.com/go/src/domain"
)

//TweeterManager estruutra del TweeterManager
type TweeterManager struct {
	userFollowing        map[string][]string
	allTweeters          map[string][]*domain.Tweeter
	lastTweeter          *domain.Tweeter
	wordCounter          map[string]int
	allDirectMessages    map[string][]*domain.Tweeter
	unreadDirectMessages map[string][]*domain.Tweeter
	favTweetersByUser    map[string][]*domain.Tweeter
}

//NewTweeterManager Crea un Tweeter manager
func NewTweeterManager() TweeterManager {
	return TweeterManager{
		make(map[string][]string), //todo lo que esta luego del primer corchete es el tipo que almacena
		make(map[string][]*domain.Tweeter),
		nil,
		make(map[string]int),
		make(map[string][]*domain.Tweeter),
		make(map[string][]*domain.Tweeter),
		make(map[string][]*domain.Tweeter),
	}
}

//GetTweeters obtiene todos los Tweeters
func (manager *TweeterManager) GetTweeters() []*domain.Tweeter {
	allTweetersInSlice := make([]*domain.Tweeter, 0)
	for _, TweeterList := range manager.allTweeters {
		for _, tweeter := range TweeterList {
			allTweetersInSlice = append(allTweetersInSlice, tweeter)
		}
	}
	return allTweetersInSlice
}

//PublishTweeter publica el Tweeter
func (manager *TweeterManager) PublishTweeter(newTweeter *domain.Tweeter) (int, error) {

	if (*newTweeter).User() == "" {
		return 0, fmt.Errorf("user is required")
	}
	if (*newTweeter).Text() == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len((*newTweeter).Text()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	manager.addWordsToCounter(*newTweeter)
	manager.addTweeterToUser(newTweeter, (*newTweeter).User())
	manager.lastTweeter = newTweeter
	return (*newTweeter).Id(), nil
}

func (manager *TweeterManager) addWordsToCounter(tweet domain.Tweeter) {
	text := (tweet).Text()
	wordsList := strings.Fields(text)

	for _, word := range wordsList {
		wordCount, _ := manager.wordCounter[word]
		manager.wordCounter[word] = wordCount + 1
	}
}

//GetTweeter obtiene el ultimo Tweeter enviado. Nil si el ultimo se boroo o no hay Tweeters
func (manager *TweeterManager) GetTweeter() *domain.Tweeter {
	return manager.lastTweeter
}

//Hacer que se reduzcan las  palabras del contador para el TT
//CleanTweeter borra el ultimo Tweeter enviado
func (manager *TweeterManager) CleanTweeter() {
	//Testear bien que ser borre el Tweeter del map
	if manager.lastTweeter != nil {
		Tweeters := manager.GetTweetersByUser((*manager.lastTweeter).User())

		for idx, Tweeter := range Tweeters {
			if Tweeter == manager.lastTweeter {
				manager.allTweeters[(*manager.lastTweeter).User()] = append(manager.allTweeters[(*manager.lastTweeter).User()][:idx], manager.allTweeters[(*manager.lastTweeter).User()][idx+1:]...)
			}
		}
	}
	manager.lastTweeter = nil
}

//Hacer test de este metodo
//CleanTweeters borra todos los Tweeters
func (manager *TweeterManager) CleanTweeters() {
	manager.allTweeters = make(map[string][]*domain.Tweeter)
	manager.lastTweeter = nil
	manager.wordCounter = make(map[string]int)
	manager.userFollowing = make(map[string][]string)
}

//GetTweeterByID obtiene el Tweeter con el id. Nil si no existe
func (manager *TweeterManager) GetTweeterByID(id int) *domain.Tweeter {
	//Obtengo todos los Tweeters
	for _, userTweeterList := range manager.allTweeters {
		for _, tweet := range userTweeterList {
			if (*tweet).Id() == id {
				return tweet
			}
		}
	}

	return nil
}

//CountTweetersByUser cuenta los Tweeters del usario
func (manager *TweeterManager) CountTweetersByUser(user string) int {
	userTweeters, usuarioExiste := manager.allTweeters[user]
	if usuarioExiste {
		return len(userTweeters)
	}
	return 0
}

//GetTweetersByUser obtiene los Tweeters del usuario
func (manager *TweeterManager) GetTweetersByUser(user string) []*domain.Tweeter {
	return manager.allTweeters[user]
}

//Follow hace que un usuario siga a otro
func (manager *TweeterManager) Follow(user1, user2 string) {
	usersFollowed, esSeguidorDeUser2 := manager.userFollowing[user1]
	if !esSeguidorDeUser2 {
		usersFollowed = make([]string, 0)
	}
	manager.userFollowing[user1] = append(usersFollowed, user2)
}

//Rehacer el tests de este metodo
//GetTimeLine obtiene la timeline del user
func (manager *TweeterManager) GetTimeLine(user string) []*domain.Tweeter {
	followed := manager.userFollowing[user]
	sliceOfTweeters := make([]*domain.Tweeter, 0)

	for _, usuario := range followed {
		userTweeters := manager.GetTweetersByUser(usuario)
		sliceOfTweeters = append(sliceOfTweeters, userTweeters...)
		//copio y creo un array nuevo con lo que le agrego, LO TENGO QUE PEGAR AL VIEJO
		//lospuntos suspensivos son para decirle "como esto es una lista , quiero todos los elementos"
	}

	sliceOfTweeters = append(sliceOfTweeters, manager.GetTweetersByUser(user)...) //los mios tambien aparecen
	return sliceOfTweeters
}

//GetTrendingTopic obtiene las dos palabras mas usadas en todos los Tweeters
func (manager *TweeterManager) GetTrendingTopic() []string {
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
func (manager *TweeterManager) SendDirectMessage(tweeter *domain.Tweeter, receiver string) {
	manager.addTweeterToMapStringKey(&manager.allDirectMessages, tweeter, receiver)
	manager.addTweeterToMapStringKey(&manager.unreadDirectMessages, tweeter, receiver)
}

func (manager *TweeterManager) addTweeterToMapStringKey(mapa *map[string][]*domain.Tweeter, Tweeter *domain.Tweeter, user string) {
	TweeterList, _ := (*mapa)[user]
	(*mapa)[user] = append(TweeterList, Tweeter)
}

//GetAllDirectMessages obtiene todos los mensajes directos de un usuario
func (manager *TweeterManager) GetAllDirectMessages(user string) []*domain.Tweeter {
	return manager.allDirectMessages[user]
}

//GetUnreadDirectMessages obtiene los mensajes sin leer de un usuario
func (manager *TweeterManager) GetUnreadDirectMessages(user string) []*domain.Tweeter {
	return manager.unreadDirectMessages[user]
}

//ReadDirectMessage marca un mensaje como leido
func (manager *TweeterManager) ReadDirectMessage(tweeter *domain.Tweeter, user string) {
	unreadDirectMessages, _ := manager.unreadDirectMessages[user]
	for index, directMessage := range unreadDirectMessages {
		if directMessage == tweeter {
			manager.unreadDirectMessages[user] = append(unreadDirectMessages[:index], unreadDirectMessages[index+1:]...)
		}
	}
}

//ReTweeter hace que el Tweeter aparezca dentro de mis Tweeters
func (manager *TweeterManager) ReTweeter(tweeter *domain.Tweeter, user string) {
	manager.addTweeterToUser(tweeter, user)
}

func (manager *TweeterManager) addTweeterToUser(tweeter *domain.Tweeter, user string) {
	manager.allTweeters[user] = append(manager.allTweeters[user], tweeter)
}

func (manager *TweeterManager) FavTweeter(tweeter *domain.Tweeter, user string) {
	manager.addTweeterToMapStringKey(&manager.favTweetersByUser, tweeter, user)
}

func (manager *TweeterManager) GetFavTweeters(user string) []*domain.Tweeter {
	return manager.favTweetersByUser[user]
}
