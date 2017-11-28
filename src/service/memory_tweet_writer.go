package service

import (
	"github.com/go/src/domain"
)

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

type ChannelTweetWriter struct {
	Memory *MemoryTweetWriter
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}

func NewChannelTweetWriter(memoryWriter *MemoryTweetWriter) *ChannelTweetWriter {
	return &ChannelTweetWriter{
		memoryWriter,
	}
}

func (channelTweet *ChannelTweetWriter) WriteTweet(channel chan domain.Tweet, channel2 chan bool) {
	tweet, open := <-channel
	for open {
		channelTweet.Memory.Tweets = append(channelTweet.Memory.Tweets, tweet)
		tweet, open = <-channel
	}
	channel2 <- true
}
