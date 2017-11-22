package service_test

import (
	"testing"

	"github.com/go/src/domain"

	"github.com/go/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var id int
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

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
	service.InitializeService()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

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
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	var id1 int
	id1, _ = service.PublishTweet(tweet)
	var id2 int
	id2, _ = service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, id1, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, id2, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v",
			id, tweet.Id)
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

func TestLastTweetDeleted(t *testing.T) {
	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)
	service.CleanTweet()

	// Validation
	if len(service.GetTweets()) != 0 {
		t.Errorf("no se borro el ultimo tweet")
	}
}

func TestAllTweetsDeleteds(t *testing.T) {
	//inicialization
	service.InitializeService()
	var tweet *domain.Tweet
	user := "giuli"
	text := "primer tweet"
	segundotext := "segundo tweet"

	tweet = domain.NewTweet(user, text)
	tweet2 := domain.NewTweet(user, segundotext)
	//operation
	service.PublishTweet(tweet)
	service.PublishTweet(tweet2)
	service.CleanTweets()
	//validation
	if len(service.GetTweets()) != 0 {
		t.Errorf("no se borraron todos los tweets")
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}
