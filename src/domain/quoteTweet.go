package domain

import "time"

//QuoteTweet es un tweet que posee imagenes
type QuoteTweet struct {
	User        string
	Text        string
	Date        *time.Time
	ID          int
	TweetQuoted Tweeter
}

//NewQuoteTweet crea un tweet de texto
func NewQuoteTweet(user, text string, tweetToQuote Tweeter) *QuoteTweet {
	date := time.Now()

	tweet := QuoteTweet{
		user,
		text,
		&date,
		-1,
		tweetToQuote,
	}
	return &tweet
}

//PrintableTweet transforma tweet a texto
func (tweet *QuoteTweet) PrintableTweet() string {
	finalText := `@`
	finalText = finalText + tweet.User
	finalText = finalText + `: `
	finalText = finalText + tweet.Text + " "

	quotedText := tweet.TweetQuoted.PrintableTweet()

	finalText = finalText + `"` + quotedText + `"`
	return finalText
}

func (tweet *QuoteTweet) String() string {
	return tweet.PrintableTweet()
}
