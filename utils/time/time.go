package timeUtils

import (
	"time"
)

func TimestampToTime(timestamp int64) *time.Time {
	tm := time.Unix(timestamp, 0)
	return &tm
}

func TimeToTimestamp(t *time.Time) int64 {
	return t.Unix()
}
