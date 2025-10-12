package user

import (
	"errors"
	"strings"
)

func CheckUsername(username string) bool {
	if len(username) < 6 || strings.Contains(username, "admin") {
		return false
	}
	return true
}

func Login(username string) (error, bool) {
	if !CheckUsername(username) {
		return errors.New("invalid username"), false
	}
	return nil, true
}
