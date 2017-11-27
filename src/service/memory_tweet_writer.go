package service

import (
	"github.com/go/src/domain"
)

type MemoryTweetWriter struct {
}

type ChannelTweetWriter struct {
}

func NewMemoryTweetWriter() MemoryTweetWriter {
	return MemoryTweetWriter{}
}

func NewChannelTweetWriter(memoryWriter MemoryTweetWriter) ChannelTweetWriter {
	return ChannelTweetWriter{}
}

func WriteTweet(channel chan domain.Tweeter, channel2 bool) {

}
