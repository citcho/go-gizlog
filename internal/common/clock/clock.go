package clock

import (
	"time"
)

var now time.Time = time.Now()

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (fc FixedClocker) Now() time.Time {
	jst := time.FixedZone("JST", 9*60*60)
	return now.In(jst)
}
