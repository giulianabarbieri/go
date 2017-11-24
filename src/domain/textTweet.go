package domain

import "time"

//TextTweet es la estructura de un tweet de solo texto
type TextTweet struct {
	User string
	Text string
	Date *time.Time
	ID   int
}

//NewTextTweet crea un tweet de texto
func NewTextTweet(user, text string) *TextTweet {
	date := time.Now()

	tweet := TextTweet{
		user,
		text,
		&date,
		-1,
	}
	return &tweet
}

//PrintableTweet transforma tweet a texto
func (tweet *TextTweet) PrintableTweet() string {
	finalText := "@"
	finalText = finalText + tweet.User
	finalText = finalText + ": "
	finalText = finalText + tweet.Text
	return finalText
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}
