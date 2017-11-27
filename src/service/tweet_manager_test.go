package service_test

import (
	"testing"

	"github.com/go/src/domain"

	"github.com/go/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet domain.Tweeter

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweeter(&tweet)

	// Validation
	publishedTweet := tweetManager.GetTweeter()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet domain.Tweeter

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweeter(&tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet domain.Tweeter

	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweeter(&tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet domain.Tweeter

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
	   all over the world. To date all community oriented activities have been organized by the community
	   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweeter(&tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet, secondTweet domain.Tweeter

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweeter(&tweet)
	secondId, _ := tweetManager.PublishTweeter(&secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweeters()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet domain.Tweeter
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweeter(&tweet)

	// Validation
	publishedTweet := tweetManager.GetTweeterByID(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet, secondTweet, thirdTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	tweetManager.PublishTweeter(&tweet)
	tweetManager.PublishTweeter(&secondTweet)
	tweetManager.PublishTweeter(&thirdTweet)

	// Operation
	count := tweetManager.CountTweetersByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweeterManager()

	var tweet, secondTweet, thirdTweet domain.Tweeter

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	firstId, _ := tweetManager.PublishTweeter(&tweet)
	secondId, _ := tweetManager.PublishTweeter(&secondTweet)
	tweetManager.PublishTweeter(&thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetersByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweeter, id int, user, text string) bool {

	if (*tweet).Id() != id {
		t.Errorf("Expected id is %v but was %v", id, (*tweet).Id())
	}

	if (*tweet).User() != user && (*tweet).Text() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, (*tweet).User(), (*tweet).Text())
		return false
	}

	if (*tweet).Date() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestIsTrendingTopic(t *testing.T) {
	manager := service.NewTweeterManager()

	userZorro1 := "Zorro1"
	userZorro2 := "Zorro2"
	textZorro1 := "Me gustan las comadrejas"
	textZorro2 := "Me encantan las comadrejas"
	text2Zorro1 := "Me casaria con una comadreja, no mentira. Las comadrejas encienden mi estomago"
	text2Zorro2 := "Que linda noche para unas comadrejas"

	var tweet1 domain.Tweeter
	var tweet2 domain.Tweeter
	var tweet3 domain.Tweeter
	var tweet4 domain.Tweeter

	tweet1 = domain.NewTextTweet(userZorro1, textZorro1)
	tweet2 = domain.NewTextTweet(userZorro2, textZorro2)
	tweet3 = domain.NewTextTweet(userZorro1, text2Zorro1)
	tweet4 = domain.NewTextTweet(userZorro2, text2Zorro2)

	manager.PublishTweeter(&tweet1)
	manager.PublishTweeter(&tweet2)
	manager.PublishTweeter(&tweet3)
	manager.PublishTweeter(&tweet4)

	theTopic := manager.GetTrendingTopic()

	if len(theTopic) != 2 {
		t.Errorf("TT no tiene dos palabras")
		return
	}
	if !(theTopic[0] == "comadrejas" && theTopic[1] == "Me") {
		t.Errorf("TT inesperado, 1° obtenido: %s, 2° obtenido: %s, cuando deberia ser comadrejas y Me ", theTopic[0], theTopic[1])
		return
	}

}

func TestCanSendDirectMessagesAndGetAllMessagesFromUser(t *testing.T) {

	//Initialization
	tweetManager := service.NewTweeterManager()

	joaquin := "joaquin"
	raul := "raul"
	msg1 := "hola man"
	msg2 := "hola!"
	msg3 := "nv"

	var tweet1 domain.Tweeter
	var tweet2 domain.Tweeter
	var tweet3 domain.Tweeter
	tweet1 = domain.NewTextTweet(joaquin, msg1)
	tweet2 = domain.NewTextTweet(raul, msg2)
	tweet3 = domain.NewTextTweet(joaquin, msg3)

	id1, _ := tweetManager.PublishTweeter(&tweet1)
	id2, _ := tweetManager.PublishTweeter(&tweet2)
	id3, _ := tweetManager.PublishTweeter(&tweet3)

	//Operation
	tweetManager.SendDirectMessage(&tweet1, raul)
	tweetManager.SendDirectMessage(&tweet2, joaquin)
	tweetManager.SendDirectMessage(&tweet3, raul)

	//Validation
	joaquinsDirectMessages := tweetManager.GetAllDirectMessages(joaquin)
	raulsDirectMessages := tweetManager.GetAllDirectMessages(raul)

	if len(joaquinsDirectMessages) != 1 {
		t.Errorf("Joaquin deberia tener 1 solo mensaje directo, tiene %v", len(joaquinsDirectMessages))
		return
	}

	if len(raulsDirectMessages) != 2 {
		t.Errorf("Raul deberia tener 2 mensajes directos, tiene %v", len(raulsDirectMessages))
	}

	isValidTweet(t, joaquinsDirectMessages[0], id2, raul, msg2)
	isValidTweet(t, raulsDirectMessages[0], id1, joaquin, msg1)
	isValidTweet(t, raulsDirectMessages[1], id3, joaquin, msg3)
}

func TestCanReadADirectMessage(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweeterManager()

	joaquin := "joaquin"
	raul := "raul"
	msg1 := "hola man"
	msg3 := "alo"

	var tweet1 domain.Tweeter
	var tweet3 domain.Tweeter
	tweet1 = domain.NewTextTweet(joaquin, msg1)
	tweet3 = domain.NewTextTweet(joaquin, msg3)

	id1, _ := tweetManager.PublishTweeter(&tweet1)
	id3, _ := tweetManager.PublishTweeter(&tweet3)

	tweetManager.SendDirectMessage(&tweet1, raul)
	tweetManager.SendDirectMessage(&tweet3, raul)

	//Operation (and validation of getunreadDM)
	raulsDirectMessages := tweetManager.GetUnreadDirectMessages(raul)

	if len(raulsDirectMessages) != 2 {
		t.Errorf("Raul deberia tener 2 mensajes directos, tiene %v", len(raulsDirectMessages))
	}

	isValidTweet(t, raulsDirectMessages[0], id1, joaquin, msg1)
	isValidTweet(t, raulsDirectMessages[1], id3, joaquin, msg3)

	tweetManager.ReadDirectMessage(&tweet3, raul)

	//Validation
	raulsDirectMessages = tweetManager.GetUnreadDirectMessages(raul)

	if len(raulsDirectMessages) != 1 {
		t.Errorf("Raul deberia tener 1 mensajes directos, tiene %v", len(raulsDirectMessages))
	}

	isValidTweet(t, raulsDirectMessages[0], id1, joaquin, msg1)
}

func TestCanRetweetSomeone(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweeterManager()

	user := "joaquin"
	userToRetweet := "raul"
	text := "This is a tweet to retweet"

	var tweet domain.Tweeter = domain.NewTextTweet(userToRetweet, text)
	id, _ := tweetManager.PublishTweeter(&tweet)

	// Operation
	tweetManager.ReTweeter(&tweet, user)

	// Validation
	myTweets := tweetManager.GetTweetersByUser(user)

	if len(myTweets) != 1 {
		t.Errorf("deberia tener 1 solo tweet, tiene %v", len(myTweets))
		return
	}

	isValidTweet(t, myTweets[0], id, userToRetweet, text)
}

func TestCanFavATweetAndGetAllFavs(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweeterManager()

	user1 := "joaquin"
	user2 := "raul"
	text1 := "this is my own tweet"
	text2 := "This is a tweet to fav"

	var tweet1 domain.Tweeter
	var tweet2 domain.Tweeter
	tweet1 = domain.NewTextTweet(user1, text1)
	tweet2 = domain.NewTextTweet(user2, text2)

	id1, _ := tweetManager.PublishTweeter(&tweet1)
	id2, _ := tweetManager.PublishTweeter(&tweet2)

	// Operation
	tweetManager.FavTweeter(&tweet1, user1)
	tweetManager.FavTweeter(&tweet2, user1)

	// Validation
	favTweets := tweetManager.GetFavTweeters(user1)

	if len(favTweets) != 2 {
		t.Errorf("deberia tener 2 tweets, tiene %v", len(favTweets))
		return
	}

	isValidTweet(t, favTweets[0], id1, user1, text1)
	isValidTweet(t, favTweets[1], id2, user2, text2)

}
