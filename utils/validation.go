package utils

import (
	"errors"
	"regexp"
)

func CheckEmailFormat(email string) error {
	emailRegexp := `/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$/`
	re := regexp.MustCompile(emailRegexp)
	if !re.MatchString(email) {
		return errors.New("email not well formatted")
	}
	return nil
}
