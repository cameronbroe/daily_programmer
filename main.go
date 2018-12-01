package main

import (
	"fmt"
	"github.com/cameronbroe/daily_programmer/cache"
	"github.com/cameronbroe/daily_programmer/reddit"
	"github.com/cameronbroe/daily_programmer/utils"
	// "github.com/go-resty/resty"
)

const Version = "0.0.1"

func main() {
	fmt.Println("daily_programmer", Version)
	dpCache := cache.New()
	testEntry := cache.CreateEntry(
		1,
		utils.DifficultyEasy,
		"foo bar",
		"https://google.com",
	)
	testEntry2 := cache.CreateEntry(
		2,
		utils.DifficultyEasy,
		"foo bar 2",
		"https://cameronbroe.com",
	)
	dpCache.AddEntry(&testEntry)
	dpCache.AddEntry(&testEntry2)
	dpCache.Display()
	dpCache.SavePretty()

	fmt.Println()
	fmt.Println()

	dpCache2 := cache.FromFile()
	dpCache2.Display()
	dpCache2.RemoveEntry(utils.DifficultyEasy, 1)
	fmt.Println()
	dpCache2.Display()

	client := reddit.New()
	client.GetPosts()
}
