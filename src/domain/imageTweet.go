package domain

import "time"

//ImageTweet es un tweet que posee imagenes
type ImageTweet struct {
	User     string
	Text     string
	Date     *time.Time
	ID       int
	ImageURL string
}

//NewImageTweet crea un tweet de texto
func NewImageTweet(user, text, imageURL string) *ImageTweet {
	date := time.Now()

	tweet := ImageTweet{
		user,
		text,
		&date,
		-1,
		imageURL,
	}
	return &tweet
}

//PrintableTweet transforma tweet a texto
func (tweet *ImageTweet) PrintableTweet() string {
	finalText := "@"
	finalText = finalText + tweet.User
	finalText = finalText + ": "
	finalText = finalText + tweet.Text
	finalText = finalText + " " + tweet.ImageURL
	return finalText
}

func (tweet *ImageTweet) String() string {
	return tweet.PrintableTweet()
}
