package util

import (
	"log"
	"time"
)

func IsTimeInRange(srcTime string, d time.Duration) bool {
	t, err := time.ParseInLocation("2006-01-02", srcTime, time.Local)
	if err != nil {
		log.Println("Failed to parse time:", err)
		return true
	}
	return t.After(time.Now().Add(-d))
}
