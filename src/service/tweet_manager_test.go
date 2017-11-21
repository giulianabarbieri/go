package service_test

import (
	"testing"

	"github.com/go/src/service"
)

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

}
