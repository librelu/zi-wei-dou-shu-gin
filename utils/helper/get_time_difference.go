package helper

import (
	"errors"
	"fmt"
	"time"
)

func GetIndex(currentTime time.Time, targetTime time.Time) (int, error) {
	currentTimeUnix := currentTime.Unix()
	targetTimeUnix := targetTime.Unix()
	if targetTimeUnix > currentTimeUnix {
		return 0, errors.New("target time should smaller than current time")
	}
	fmt.Println((currentTimeUnix - targetTimeUnix) / 60 / 60 / 24)
	return int((currentTimeUnix - targetTimeUnix) / 60 / 60 / 24 % 12), nil
}
