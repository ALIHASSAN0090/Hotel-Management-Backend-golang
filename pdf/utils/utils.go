package pdf

import "time"

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}