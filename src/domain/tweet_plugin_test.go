package domain_test

import (
	"testing"

	"github.com/go/src/domain"
	"github.com/go/src/service"
)

func TestAddPlugin(t *testing.T) {
	//Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)
	tweetManager := service.NewTweetManager(tweetWriter)
	var myplugin domain.TweetPlugin = domain.NewFacebookPlugin()

	//Operation
	tweetManager.AddPlugin(myplugin)

	//Validation
	listaDePlugins := tweetManager.GetPlugins()
	if len(listaDePlugins) != 1 {
		t.Errorf("no cargó el plugin")
		return
	}
	if listaDePlugins[0] != myplugin {
		t.Errorf("cargó plugin distinto")
		return
	}
}

func TestCanAddMoreThan1Plugin(t *testing.T) {
	//Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)
	tweetManager := service.NewTweetManager(tweetWriter)

	var myplugin domain.TweetPlugin = domain.NewFacebookPlugin()
	var anotherplugin domain.TweetPlugin = domain.NewGooglePlusPlugin()

	//Operation
	tweetManager.AddPlugin(myplugin)
	tweetManager.AddPlugin(anotherplugin)

	//Validation
	listaDePlugins := tweetManager.GetPlugins()
	if len(listaDePlugins) != 2 {
		t.Errorf("No cargó adecuadamente la cantidad de plugins")
		return
	}
	if listaDePlugins[0] != myplugin || listaDePlugins[1] != anotherplugin {
		t.Errorf("No cargó adecuadamente los plugins indicados")
		return
	}
}

func TestPluginWorks(t *testing.T) {
	//Initialization
	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)
	tweetManager := service.NewTweetManager(tweetWriter)

	var myplugin domain.TweetPlugin = domain.NewFacebookPlugin()
	var anotherplugin domain.TweetPlugin = domain.NewGooglePlusPlugin()

	tweetManager.AddPlugin(myplugin)
	tweetManager.AddPlugin(anotherplugin)
	var tweet domain.Tweet = domain.NewTextTweet("bs", "asd")

	//Operation
	quit := make(chan bool)
	tweetManager.PublishTweet(tweet, quit)
	<-quit

	//Validation

	if len(tweetManager.PluginMessages()) != 2 {
		t.Errorf("No se corrieron los plugins")
		return
	}
}
