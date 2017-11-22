package service

import "github.com/go/src/domain"

var tweet *domain.Tweet

//la estructura ES el tipo
func PublishTweet(newTweet *domain.Tweet) {
	tweet = newTweet
}
func GetTweet() *domain.Tweet {
	return tweet
}
func CleanTweet() {
	tweet = nil
}
