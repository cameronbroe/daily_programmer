package utils

import (
	"testing"
)

func TestEasyTitle(t *testing.T) {
	title := "[2018-09-04] Challenge #367 [Easy] Subfactorials - Another Twist on Factorials"
	_, difficulty := ParseTitle(title)
	if difficulty == DifficultyEasy {
		t.Log("Difficulty has been succesfully parsed as: ", difficulty)
	} else {
		t.Fatal("Difficulty has been non-successfully parsed as: ", difficulty)
	}
}
