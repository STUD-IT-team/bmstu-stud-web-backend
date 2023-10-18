package times

import "time"

func Elapsed(startedAt time.Time) time.Duration {
	return time.Since(startedAt)

}
