package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestGothrutweets(t *testing.T) {

	anaconda.SetConsumerKey("CRYLw41CwCqaJyUZIUvYpUcLX")
	anaconda.SetConsumerSecret("quuoBQZCNtYgs5ieLb6QRJgW72Ck8vsuvCZY6V0NeFEmXD472s")

	api := anaconda.NewTwitterApi(
		"1295387221600526337-rehqa8Aj7zH7iHnSQV87o3e0hxZIk0",
		"s0B62DLqKeyUtZqLcKSMnyG1j7EpY0UGkvmBSE0k2PN6p",
	)
	tweet, _ := api.GetTweet(1295746102654713856, nil) //1295746102654713856
	var expectedReturn []anaconda.Tweet
	retweets := Gothrutweets(tweet, api)

	//fmt.Println(reflect.ValueOf(retweets).Kind())
	if reflect.TypeOf(retweets) != reflect.TypeOf(expectedReturn) {
		t.Error("Wrong type returned.")
	}
	if len(retweets) == 0 {
		t.Error("Couldn't get any retweets! :(")

	}

}

func TestWin(t *testing.T) {

	ans := Win("i_kanganaranaut", "8", "8")
	fmt.Println(len(ans))
	if len(ans) < 1 {
		t.Error("Didn't get a valid winner.")
	}

	if reflect.TypeOf(ans) != reflect.TypeOf(5) {
		t.Error("Wrong type returned.")
	}

}
