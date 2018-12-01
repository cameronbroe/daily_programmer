package utils

import (
	"../cache"
	"regexp"
	"strconv"
)

func ParseTitle(title string) (number int, difficulty cache.DifficultyType) {
	// Get the number first
	numberRegex, err := regexp.Compile(`\s#(\d+)\s`)
	if err != nil {
		panic(err)
	}
	numberMatch := numberRegex.FindStringSubmatch(title)

	if len(numberMatch) > 0 {
		number, err = strconv.Atoi(numberMatch[0])
		if err != nil {
			panic(err)
		}
	} else {
		number = 0
	}

	easyRegex, err := regexp.Compile(`[Easy]`)
	if err != nil {
		panic(err)
	}
	intermediateRegex, err := regexp.Compile(`[Intermediate]`)
	if err != nil {
		panic(err)
	}
	hardRegex, err := regexp.Compile(`[Hard]`)
	if err != nil {
		panic(err)
	}

	switch {
	case easyRegex.MatchString(title):
		difficulty = DifficultyEasy
	case intermediateRegex.MatchString(title):
		difficulty = DifficultyMedium
	case hardRegex.MatchString(title):
		difficulty = DifficultyHard
	}

	return
}
