package service

import "github.com/go/src/domain"
import "fmt"

var allTweets []*domain.Tweet

func InitializeService() {
	allTweets = make([]*domain.Tweet, 0)
}

func GetTweets() []*domain.Tweet {

	return allTweets
}

//la estructura ES el tipo
func PublishTweet(newTweet *domain.Tweet) (int, error) {

	if newTweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(newTweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}
	allTweets = append(allTweets, newTweet)
	return newTweet.Id, nil
}

func GetTweet() *domain.Tweet {
	if len(allTweets) == 0 {
		return nil //HACER ESTO DE UN TEST
	}
	return allTweets[len(allTweets)-1]
}

func CleanTweet() {
	//FIJARSE DE NO BORRAR CUANDO NO HAY TWEETS
	allTweets = allTweets[:len(allTweets)-1]
}

func CleanTweets() {
	allTweets = make([]*domain.Tweet, 0)
}

func GetTweetById(id int) *domain.Tweet {
	for _, element := range allTweets {
		if element.Id == id {
			return element
		}
	}
	return nil
}
