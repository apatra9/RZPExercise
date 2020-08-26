package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"twit/twit"

	"github.com/ChimeraCoder/anaconda"
)

func TestGothrutweets(t *testing.T) {

	anaconda.SetConsumerKey(os.Getenv("CONSUMERKEY"))
	anaconda.SetConsumerSecret("CONSUMERSECRET")

	api := anaconda.NewTwitterApi(
		"APPKEY",
		"APPSECRET",
	)
	tweet, _ := api.GetTweet(1295746102654713856, nil) //1295746102654713856, valid tweet

	var expectedReturn []anaconda.Tweet

	tweet, _ = api.GetTweet(2295746102654713856, nil) //2295746102654713856, invalid tweet

	retweets := twit.Gothrutweets(tweet, api)

	//fmt.Println(reflect.ValueOf(retweets).Kind())
	if reflect.TypeOf(retweets) != reflect.TypeOf(expectedReturn) {
		t.Error("Wrong type returned.")
	}
	if len(retweets) == 0 {
		t.Error("Couldn't get any retweets! :(") //add happy test cases and sad test cases separately

	}

	tweet, _ = api.GetTweet(2295746102654713856, nil) //2295746102654713856, invalid tweet
	if reflect.TypeOf(retweets) != reflect.TypeOf(expectedReturn) {
		t.Error("Wrong type returned.")
	}
	if len(retweets) == 0 {
		t.Error("Couldn't get any retweets! :(") //add happy test cases and sad test cases separately

	}

	tweet, _ = api.GetTweet(-2, nil) //negative, invalid tweet

	if reflect.TypeOf(retweets) != reflect.TypeOf(expectedReturn) {
		t.Error("Wrong type returned.")
	}
	if len(retweets) == 0 {
		t.Error("Couldn't get any retweets! :(") //add happy test cases and sad test cases separately

	}

}

func TestWin(t *testing.T) {

	ans := twit.Win("i_kanganaranaut", "8", "8") //good id
	fmt.Println(len(ans))
	if len(ans) < 1 {
		t.Error("Didn't get a valid winner.")
	}

	if reflect.TypeOf(ans) != reflect.TypeOf(5) { //sad test case
		t.Error("Wrong type returned.")
	}

	if reflect.TypeOf(ans) != reflect.TypeOf("abc") { //happy test case
		t.Error("Wrong type returned.")
	}
	ans = twit.Win("i_kanganakaraut", "8", "8") //bad id

	if len(ans) < 1 {
		t.Error("Didn't get a valid winner.")
	}

	if reflect.TypeOf(ans) != reflect.TypeOf(5) { //sad test case
		t.Error("Wrong type returned.")
	}

	if reflect.TypeOf(ans) != reflect.TypeOf("abc") { //happy test case
		t.Error("Wrong type returned.")
	}

}
