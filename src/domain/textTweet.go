package domain

import "time"

//TextTweet es la estructura de un tweet de solo texto
type TextTweet struct {
	UserATR string
	TextATR string
	Date    *time.Time
	IDATR   int
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
	finalText = finalText + tweet.User()
	finalText = finalText + ": "
	finalText = finalText + tweet.Text()
	return finalText
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *TextTweet) User() string {
	return tweet.UserATR
}

func (tweet *TextTweet) Text() string {
	return tweet.TextATR
}
func (tweet *TextTweet) Id() int {
	return tweet.IDATR
}
