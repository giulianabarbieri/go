package domain

type TweetPlugin interface {
	RunPlugin() string
}

type FacebookPlugin struct {
}

func NewFacebookPlugin() *FacebookPlugin {
	return &FacebookPlugin{}
}

func (fbPlugin *FacebookPlugin) RunPlugin() string {
	return "se publicó en facebook"
}

type GooglePlusPlugin struct {
}

func NewGooglePlusPlugin() *GooglePlusPlugin {
	return &GooglePlusPlugin{}
}

func (GPPlugin *GooglePlusPlugin) RunPlugin() string {
	return "se publicó en Google Plus"
}
