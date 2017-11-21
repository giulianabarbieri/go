package service

var Tweet string

func PublishTweet(newTweet string) {
	Tweet = newTweet
}
func GetTweet() string {
	return Tweet
}
