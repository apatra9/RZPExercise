package twit

import (
	"fmt"
	"log"
	"net/url"
	"os"
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

func WinCron(c *gin.Context) {

}

func querybooster(c *gin.Context) { //tests localhost:8080/winner?id=i_kanganaranaut&tweetCount=15&rtCount=15
	id := c.Query("id")
	tc := c.Query("tweetCount")
	rc := c.Query("rtCount") //bind json

	if len(id) == 0 {
		c.JSON(200, gin.H{
			"error": "Invalid twitter id",
		})

		tweetc, _ := strconv.Atoi(tc)

		if (tweetc) < 1 {
			c.JSON(200, gin.H{
				"error": "tweet count too low",
			})
		}
		retweetc, _ := strconv.Atoi(rc)

		if (retweetc) < 1 {
			c.JSON(200, gin.H{
				"error": "retweet count too low",
			})
		}

	}

	str := Win(id, tc, rc)
	//add validation here str

	c.JSON(200, gin.H{
		"id":            id,
		"tweetcount":    tc,
		"retweetcount:": rc,
		"winner":        str,
	})

}

func Win(username string, tweetCount string, rtCount string) string {
	var retweets []anaconda.Tweet
	var err error

	db, err = gorm.Open("sqlite3", "/tmp/gorm.db") //create init database

	if err != nil {
		log.Panic(err)
	}

	//db.DropTableIfExists(&UserModel{})
	db.AutoMigrate(&UserModel{})

	anaconda.SetConsumerKey(os.Getenv("CONSUMERKEY"))			//get env variables for these
	anaconda.SetConsumerSecret(os.Getenv("CONSUMERSECRET"))

	api := anaconda.NewTwitterApi(
		os.Getenv("APPKEY"),
		os.Getenv("APPSECRET"),
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
			defer wg.Done() //look

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

func ThrowWin() {
	r := gin.Default()
	r.GET("/winner", querybooster)
	r.Run()

}

func ThrowWinners() {

	r := gin.Default()
	r.GET("/winners", GetUsers)
	r.Run()
}

//curl --location --request GET 'localhost:8080/winner?id=i_kanganaranaut&tweet_count=5&retweet_count=25'
//crontab paste this at certain interval
//environment variable store secrets
