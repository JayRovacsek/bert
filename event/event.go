package event

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	robot "github.com/go-vgo/robotgo"
)

const (
	left = iota
	right
)

type Event struct {
	x          int
	xDeviance  [2]int
	y          int
	yDeviance  [2]int
	click      bool
	mouseKey   click
	keyPress   bool
	keyEvent   fyne.KeyEvent
	timeBefore time.Duration
	timeAfter  time.Duration
	exitKey    rune
}

type click struct {
	button int
}

func (e Event) Run(i int) (bool, error) {
	fmt.Println(fmt.Sprintf("Running event: %v. Click: %v, Keypress: %v, Mousebounds: x: %v, y: %v", i, e.click, e.keyPress, e.xDeviance, e.yDeviance))
	pid := robot.ActivePID

	fmt.Println(pid)

	return true, nil
}
