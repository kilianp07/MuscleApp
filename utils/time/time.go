package timeUtils

import "time"

func TimestampToTime(timestamp int64) *time.Time {
	t := time.Unix(timestamp, 0)
	return &t
}
