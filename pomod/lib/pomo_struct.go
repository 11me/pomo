package lib

import (
	"fmt"
	"os/exec"
	"time"
)

// Pomo struct represents a pomodoro.
type Pomo struct {
	Status   string // status running or not could be ON or OFF
	Time     time.Duration
	Notified bool
	Emoji    string    // default tomato emoji
	Blink    [2]string // array of 2 emoji to create an animation
	Cancel   chan bool // channel for cancelation a running timer goroutine
}

// Constants to represent a status of the pomo.
const (
	ON  = "on"
	OFF = "off"
)

const NOTIFICATION_MESSAGE = "🍅 Time to take a break!"

var DEFAULT_DURATION, _ = time.ParseDuration("25m")

// Start starts the pomodoto timer by running a goroutine.
func (pomo *Pomo) Start() {

	// Start the goroutine if not already running.
	if pomo.Status == OFF {

		pomo.Status = ON
		pomo.Notified = false

		// goroutine cancelation channel
		cancel := make(chan bool)

		go func() {

			for {

				select {

				case <-cancel:
					pomo.Status = OFF
					return

				default:
					pomo.Time -= time.Second

				}
				time.Sleep(time.Second)
			}
		}()

		pomo.Cancel = cancel
	}
}

// Stop stops the running timer goroutine.
func (pomo *Pomo) Stop() {

	if pomo.Status == ON {
		pomo.Cancel <- true
		pomo.Time = DEFAULT_DURATION
	}

}

// SetDuration sets the duration of pomo timer
// it accepts time.Duration parsable format.
func (pomo *Pomo) SetDuration(d time.Duration) {

	// if pomo is running
	if pomo.Status == ON {

		pomo.Stop()
		pomo.Time = d
		pomo.Start()

		return
	}

	pomo.Time = d

	// change notified status when duration is changed
	pomo.Notified = false
}

// GetDuration returns the current duration.
func (pomo *Pomo) GetDuration() string {

	// if pomo is off do not output anything
	if pomo.Status == OFF {
		return ""
	}

	// if pomo run out of time that was set
	// make a blinking animation and send ntification
	if pomo.Time < 0 {

		// if user not notified
		if !pomo.Notified {

			// notify the user
			go notifyUser(NOTIFICATION_MESSAGE)

			pomo.Notified = true
		}

		// emoji_id is a number between 0 and 1
		emoji_id := (pomo.Time.Milliseconds() / 1000 % 2) * (-1)

		return fmt.Sprintf("%s%s\n", pomo.Blink[emoji_id], pomo.Time)
	}

	return fmt.Sprintf("%s%s\n", pomo.Emoji, pomo.Time)
}

// Restart restarts the pomo timer to default value.
func (pomo *Pomo) Restart() {
	pomo.SetDuration(DEFAULT_DURATION)
}

// NotifyUser uses external notify-send command
// sends specified message.
func notifyUser(msg string) {
	cmd := exec.Command("notify-send", "-u", "critical", msg)
	cmd.Run()
	return
}
