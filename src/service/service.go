package service

import "github.com/go/src/domain"
import "fmt"

var tweet *domain.Tweet

//la estructura ES el tipo
func PublishTweet(newTweet *domain.Tweet) error {

	if newTweet.User == "" {
		return fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(newTweet.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}

	tweet = newTweet
	return nil
}

func GetTweet() *domain.Tweet {
	return tweet
}
func CleanTweet() {
	tweet = nil
}
