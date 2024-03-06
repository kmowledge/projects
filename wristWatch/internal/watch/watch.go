package general

import "time"

type Mode uint8

var M Mode

const (
	ModeClock     Mode = iota
	ModeStopwatch Mode = 1
	ModeTimer     Mode = 2
)

type Display struct {
	ModeName string
	TimeVal  time.Time
}

type Watch struct {
	Clock     Display
	Stopwatch Display
	Timer     Display
}

type WatchSimple struct {
	Mode int //int vs [3]int? I want to have range hardcoded in the class attribute
	Display string
	Clock time.Time.Format(time.TimeOnly)
	Stopwatch time.Time.Format(time.TimeOnly)
	Timer time.Time.Format(time.TimeOnly)
}

var W Watch
var W2 WatchSimple

func init() {
	W.Clock.ModeName = "Clock"
	W.Clock.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	W.Stopwatch.ModeName = "Stopwatch"
	W.Stopwatch.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	W.Timer.ModeName = "Timer"
	W.Timer.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
}

