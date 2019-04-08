package event

import (
	"fmt"
	"time"

	"fyne.io/fyne"
)

const (
	left = iota
	right
)

// Event struct to define components of an event
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

// Run method for executing events
func (e Event) Run(i int) (bool, error) {
	fmt.Println(fmt.Sprintf("Running event: %v. Click: %v, Keypress: %v, Mouse deviance: x: %v, y: %v", i, e.click, e.keyPress, e.xDeviance, e.yDeviance))

	return true, nil
}
