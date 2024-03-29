package general

import (
	"time"
	"fmt"
	"strconv"
)

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

type Watch2 struct {
	Mode int //int vs [3]int? I want to have range hardcoded in the class attribute
	Display string
	Clock time.Time.Format(time.TimeOnly)
	Stopwatch time.Time.Format(time.TimeOnly)
	Timer time.Time.Format(time.TimeOnly)
}

var W Watch
var W2 Watch2

func init() {
	W.Clock.ModeName = "Clock"
	W.Clock.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	W.Stopwatch.ModeName = "Stopwatch"
	W.Stopwatch.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	W.Timer.ModeName = "Timer"
	W.Timer.TimeVal = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
}

// bellow adding the content from other files
type Button uint8

var b Button

// func init() { b = 1 }

const (
	Button1Start Button = iota
	Button1Stop  Button = 1
	Button1Reset Button = 2
	Button2Mode  Button = 3
	Button2Set   Button = 4
	Button3Plus  Button = 5
	Button4Minus Button = 6
)

func (W general.Watch) Snap(b Button) {
	var b1 string
	fmt.Scanln(&b1)
	b2, _ := strconv.ParseUint(b1, 10, 8)
	b = Button(b2)
	switch b {
	case Button1Start:
		fmt.Println("Start")
		Start()
	case Button1Stop:
		fmt.Println("Stop")
		Stop()
	case Button1Reset:
		fmt.Println("Reset")
		Reset()
	case Button2Mode:
		//przeskakuje pomiędzy 0 1 2
		fmt.Println("Mode")
		Mode()
	case Button2Set:
		fmt.Println("Set")
		Set()
	case Button3Plus:
		fmt.Println("Plus")
		Plus()
	case Button4Minus:
		fmt.Println("Minus")
		Minus()
	}
}


func (W *general.Watch) Mode() {
	general.M++
	switch general.M {
	case general.ModeClock:
		general.Watch.Clock.ModeName = "Clock"
		general.Watch.Clock.TimeVal = ""
	case general.ModeStopwatch:
		general.Watch.Stopwatch.ModeName = "Stopwatch"
		general.Watch.Stopwatch.TimeVal = "" //zachowana wartość z fmt.Sprint lub global var
	case general.ModeTimer:
		general.Watch.Timer.ModeName = "Timer"
		general.Watch.Timer.TimeVal = ""
	}
}

func (d *display.Display) Start() {
	time.Since() //Start counting

}

func (d *display.Display) Start() {
	time.Start()
}

type Starter interface {
	Start()
}

type Stopwatch interface {
	Starter()
	Stopper()
	Resetter()
}

type Timer interface {
	Starter()
	Stopper()
	Resetter()
	Setter()
}
