package helpers

import "time"

func Now() string {
	layout := "2006-01-02 15:04:05"
	return time.Now().Format(layout)
}
