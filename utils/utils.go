package utils

import (
	"fmt"
	"os"
	"os/user"
)

func DefaultCachePath() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	defaultPath := fmt.Sprintf("%s/.daily_programmer_cache", currentUser.HomeDir)
	return defaultPath
}

func DefaultCacheExists() bool {
	_, err := os.Stat(DefaultCachePath())
	if os.IsNotExist(err) {
		return false
	}
	return true
}
