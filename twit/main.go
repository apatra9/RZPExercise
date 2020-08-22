package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"sync"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Gothrutweets(tweet anaconda.Tweet, api *anaconda.TwitterApi) []anaconda.Tweet {

	z := url.Values{}
	var retweets []anaconda.Tweet
	z.Set("id", tweet.IdStr)
	z.Set("count", "5")
	catch, _ := api.GetRetweets(tweet.Id, z)
	//fmt.Println(catch)
	retweets = append(retweets, catch...) ///... used to concat two slice , working
	//userids = append(userids, catch.User)
	return retweets

}

var db *gorm.DB

type UserModel struct {
	ID    int    `gorm:"primary_key"; "AUTO_INCREMENT"`
	Name  string `gorm:"size:255"`
	Count string `gorm:"type:varchar(100)"`
}

func database(db *gorm.DB, max string, count string) {

	user := &UserModel{Name: max, Count: count}
	db.Debug().Create(user)
	db.Debug().Where("Count LIKE ?", "%3%").Find(&user)
	log.Println(user)
	users := []UserModel{}
	db.Find(&users)
	log.Println(users)

}

func GetUsers(c *gin.Context) {
	users := []UserModel{}
	db.Find(&users)
	c.JSON(200, &users)
}

func querybooster(c *gin.Context) { //tests localhost:8080/winner?id=i_kanganaranaut&tweetCount=15&rtCount=15
	a := c.Query("id")
	b := c.Query("tweetCount")
	d := c.Query("rtCount")

	str := Win(a, b, d)

	c.JSON(200, gin.H{
		"id is":            a,
		"tweetcount is":    b,
		"retweetcount is:": d,
		"winner is:":       str,
	})

}

func Win(username string, tweetCount string, rtCount string) string {
	var retweets []anaconda.Tweet

	db, err := gorm.Open("sqlite3", "/tmp/gorm.db") //create init database

	if err != nil {
		log.Panic(err)
	}

	//db.DropTableIfExists(&UserModel{})
	db.AutoMigrate(&UserModel{})

	anaconda.SetConsumerKey("CRYLw41CwCqaJyUZIUvYpUcLX")
	anaconda.SetConsumerSecret("quuoBQZCNtYgs5ieLb6QRJgW72Ck8vsuvCZY6V0NeFEmXD472s")

	api := anaconda.NewTwitterApi(
		"1295387221600526337-rehqa8Aj7zH7iHnSQV87o3e0hxZIk0",
		"s0B62DLqKeyUtZqLcKSMnyG1j7EpY0UGkvmBSE0k2PN6p",
	)

	v := url.Values{}

	v.Set("screen_name", username)
	v.Set("count", tweetCount)
	v.Set("include_rts", "false")

	results, _ := (api.GetUserTimeline(v)) //get 100 tweets by user

	var wg sync.WaitGroup
	wg.Add(len(results))

	for _, tweet := range results { //getting retweets from each of 100 tweets

		go func(tweet anaconda.Tweet) {
			defer wg.Done()

			z := url.Values{}
			z.Set("id", tweet.IdStr)
			z.Set("count", rtCount)
			catch, _ := api.GetRetweets(tweet.Id, z)

			retweets = append(retweets, catch...) ///... used to concat two slice , working

		}(tweet)

	}

	wg.Wait()

	retweetmap := make(map[string]int)
	temp := 1
	max := ""

	for _, t := range retweets {
		retweetmap[(t.User.ScreenName)]++
	}
	fmt.Println(retweetmap)

	for name, count := range retweetmap {
		//fmt.Println(name, age)
		if temp < count {
			max = name
			temp = count
		}
	}
	fmt.Println("Winner is:", max, "with rt:", retweetmap[max])
	//database(db, max, strconv.Itoa(retweetmap[max]))
	database(db, max, strconv.Itoa(retweetmap[max]))

	return max

}

func main() {

	r := gin.Default()
	r.GET("/winner", querybooster)

	//win("rf", 5, 5)

	r.GET("/winners", GetUsers)
	r.Run()

}
