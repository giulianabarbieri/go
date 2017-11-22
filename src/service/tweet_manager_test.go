package service_test

import (
	"testing"

	"github.com/go/src/domain"

	"github.com/go/src/service"
)

func TestCleanTweet(t *testing.T) {

	var tweet *domain.Tweet //Tweet es el nombre de la estructura!!

	user := "grupoesfera"
	text := "hello"
	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)
	service.CleanTweet()

	if service.GetTweet() != nil {
		t.Error("Expected empty string")
	}

}

func TestPublishedTweetIsSaved(t *testing.T) {
	//inicialization
	var tweet *domain.Tweet //Tweet es el nombre de la estructura!!

	user := "grupoesfera"
	text := "hello"

	tweet = domain.NewTweet(user, text)

	//operation
	service.PublishTweet(tweet)

	//validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Error("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)

	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}
func TestTweetWhithoutUserIsNotPublished(t *testing.T) {
	//inicialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//operation
	var err error
	err = service.PublishTweet(tweet)

	//validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}
func TestTweetWhithoutTextIsNotPublished(t *testing.T) {
	//inicialization
	var tweet *domain.Tweet

	user := "El grupo del canal cultural"
	var text string

	tweet = domain.NewTweet(user, text)

	//operation
	var err error
	err = service.PublishTweet(tweet)

	//validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is requered")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community
	somethingthatcannotbereadonscreen all over the world. To date all communityyyyyyyy
	yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy
	yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy`

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
		return
	}
}
