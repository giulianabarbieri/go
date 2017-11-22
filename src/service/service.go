package service

import "github.com/go/src/domain"
import "fmt"

var tweet *domain.Tweet

//la estructura ES el tipo
func PublishTweet(newTweet *domain.Tweet) error {
	var err error = nil
	if newTweet.User == "" {
		err = fmt.Errorf("user is required")
	}
	tweet = newTweet
	return err
}

func GetTweet() *domain.Tweet {
	return tweet
}
func CleanTweet() {
	tweet = nil
}
