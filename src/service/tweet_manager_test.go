package service_test

import (
	"testing"

	"github.com/go/src/domain"

	"github.com/go/src/service"
)

/*
func TestPublishedTweetSaved(t *testing.T) {

	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}

}


func TestCleanTweet(t *testing.T) {

	tweet := "This is my first tweet"

	service.PublishTweet(tweet)
	service.CleanTweet()

	if service.GetTweet() != "" {
		t.Error("Expected empty string")
	}

} */

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
