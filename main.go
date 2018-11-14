package main

import (
    "fmt"
    "./cache"
    // "github.com/go-resty/resty"
)

const Version = "0.0.1"

func main() {
    fmt.Println("daily_programmer", Version)
    dpCache := cache.New()
    testEntry := cache.CreateEntry(
        1,
        cache.DifficultyEasy,
        "foo bar",
        "https://google.com",
    )
    testEntry2 := cache.CreateEntry(
        2,
        cache.DifficultyEasy,
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
    dpCache2.RemoveEntry(cache.DifficultyEasy, 1)
    fmt.Println()
    dpCache2.Display()
}
