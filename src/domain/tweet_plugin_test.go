package domain_test

import (
	"testing"

	"github.com/go/src/domain"
	"github.com/go/src/service"
)

func TestAddPlugin(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweeterManager()
	var myplugin1 domain.TweetPlugin = domain.NewFacebookPlugin()
	myplugin := &myplugin1
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
	tweetManager := service.NewTweeterManager()
	var myplugin1 domain.TweetPlugin = domain.NewFacebookPlugin()
	myplugin := &myplugin1
	var anotherplugin1 domain.TweetPlugin = domain.NewGooglePlusPlugin()
	anotherplugin := &anotherplugin1

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
	tweetManager := service.NewTweeterManager()
	var myplugin1 domain.TweetPlugin = domain.NewFacebookPlugin()
	myplugin := &myplugin1
	var anotherplugin1 domain.TweetPlugin = domain.NewGooglePlusPlugin()
	anotherplugin := &anotherplugin1
	tweetManager.AddPlugin(myplugin)
	tweetManager.AddPlugin(anotherplugin)
	var tweet domain.Tweeter = domain.NewTextTweet("bs", "asd")

	//Operation
	_, _, pluginMessages := tweetManager.Publish2Tweeter(&tweet)

	//Validation

	if len(pluginMessages) != 2 {
		t.Errorf("No se corrieron los plugins")
		return
	}
}
