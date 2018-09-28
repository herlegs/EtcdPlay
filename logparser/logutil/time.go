package logutil

import (
	"fmt"
	"time"
)

func ParseTime(timeStr, layout string) time.Time {
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("error parsing time: %v\n", err)
	}
	return t
}
