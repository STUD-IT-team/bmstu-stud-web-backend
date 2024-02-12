package times

import "time"

var (
	TZMoscow, _ = time.LoadLocation("Europe/Moscow")
)

func Elapsed(startedAt time.Time) time.Duration {
	return time.Since(startedAt)

}
