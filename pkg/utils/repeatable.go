package utils

import (
	"os"
	"time"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}

var (
	PublicFilePath = "./../../public"
)

func CrateDir() error {
	return os.MkdirAll(PublicFilePath, os.ModePerm)
}
