package service_test

import (
	"fmt"
	"testing"

	"github.com/go/src/domain"

	"github.com/go/src/service"
)

func TestCanWriteATweet(t *testing.T) {

	// Initialization
	var tweet domain.Tweet = domain.NewTextTweet("grupoesfera", "Async tweet")
	var tweet2 domain.Tweet = domain.NewTextTweet("grupoesfera", "Async tweet2")

	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetsToWrite := make(chan domain.Tweet)
	quit := make(chan bool)

	go tweetWriter.WriteTweet(tweetsToWrite, quit)

	// Operation
	tweetsToWrite <- tweet
	tweetsToWrite <- tweet2
	close(tweetsToWrite)

	<-quit

	fmt.Println(len(memoryTweetWriter.Tweets))

	// Validation
	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected, it was %v and it should %v", memoryTweetWriter.Tweets[0], &tweet)
	}

	if memoryTweetWriter.Tweets[1] != tweet2 {
		t.Errorf("A tweet in the writer was expected, it was %v and it should %v", memoryTweetWriter.Tweets[1], &tweet2)
	}
}
