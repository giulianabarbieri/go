package service_test

import (
	"testing"

	"github.com/go/src/domain"

	"github.com/go/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

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
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
	   all over the world. To date all community oriented activities have been organized by the community
	   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

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
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

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
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetByID(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

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

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.Id)
	}

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestIsTrendingTopic(t *testing.T) {
	manager := service.NewTweetManager()

	userZorro1 := "Zorro1"
	userZorro2 := "Zorro2"
	textZorro1 := "Me gustan las comadrejas"
	textZorro2 := "Me encantan las comadrejas"
	text2Zorro1 := "Me casaria con una comadreja, no mentira. Las comadrejas encienden mi estomago"
	text2Zorro2 := "Que linda noche para unas comadrejas"

	tweet1 := domain.NewTweet(userZorro1, textZorro1)
	tweet2 := domain.NewTweet(userZorro2, textZorro2)
	tweet3 := domain.NewTweet(userZorro1, text2Zorro1)
	tweet4 := domain.NewTweet(userZorro2, text2Zorro2)

	manager.PublishTweet(tweet1)
	manager.PublishTweet(tweet2)
	manager.PublishTweet(tweet3)
	manager.PublishTweet(tweet4)

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
	tweetManager := service.NewTweetManager()

	joaquin := "joaquin"
	raul := "raul"
	msg1 := "hola man"
	msg2 := "hola!"
	msg3 := "nv"

	tweet1 := domain.NewTweet(joaquin, msg1)
	tweet2 := domain.NewTweet(raul, msg2)
	tweet3 := domain.NewTweet(joaquin, msg3)

	id1, _ := tweetManager.PublishTweet(tweet1)
	id2, _ := tweetManager.PublishTweet(tweet2)
	id3, _ := tweetManager.PublishTweet(tweet3)

	//Operation
	tweetManager.SendDirectMessage(tweet1, raul)
	tweetManager.SendDirectMessage(tweet2, joaquin)
	tweetManager.SendDirectMessage(tweet3, raul)

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
	tweetManager := service.NewTweetManager()

	joaquin := "joaquin"
	raul := "raul"
	msg1 := "hola man"
	msg3 := "alo"

	tweet1 := domain.NewTweet(joaquin, msg1)
	tweet3 := domain.NewTweet(joaquin, msg3)

	id1, _ := tweetManager.PublishTweet(tweet1)
	id3, _ := tweetManager.PublishTweet(tweet3)

	tweetManager.SendDirectMessage(tweet1, raul)
	tweetManager.SendDirectMessage(tweet3, raul)

	//Operation (and validation of getunreadDM)
	raulsDirectMessages := tweetManager.GetUnreadDirectMessages(raul)

	if len(raulsDirectMessages) != 2 {
		t.Errorf("Raul deberia tener 2 mensajes directos, tiene %v", len(raulsDirectMessages))
	}

	isValidTweet(t, raulsDirectMessages[0], id1, joaquin, msg1)
	isValidTweet(t, raulsDirectMessages[1], id3, joaquin, msg3)

	tweetManager.ReadDirectMessage(tweet3, raul)

	//Validation
	raulsDirectMessages = tweetManager.GetUnreadDirectMessages(raul)

	if len(raulsDirectMessages) != 1 {
		t.Errorf("Raul deberia tener 1 mensajes directos, tiene %v", len(raulsDirectMessages))
	}

	isValidTweet(t, raulsDirectMessages[0], id1, joaquin, msg1)
}
