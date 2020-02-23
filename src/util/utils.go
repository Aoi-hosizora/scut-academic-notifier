package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 00.00 -> 2019-01-01
func ParseSpTimeString(s string) string {
	sp := strings.Split(s, ".") // 08.30
	ret := s
	if len(sp) > 2 {
		month, err1 := strconv.Atoi(sp[0])
		day, err2 := strconv.Atoi(sp[1])
		year := time.Now().Year()
		if err1 != nil || err2 != nil {
			return ret
		}

		fakeTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		if fakeTime.After(time.Now()) {
			year -= 1
		}
		ret = fmt.Sprintf("%d-%d-%d", year, month, day)
	}
	return ret
}
