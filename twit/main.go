package main

import (
	"fmt"
	"twit/twit"
)

func main() {

	fmt.Println(twit.Win("i_kanganakaraut", "5", "5"))
	twit.ThrowWin() //called by cron
	twit.ThrowWinners()

}
