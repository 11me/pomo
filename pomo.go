// Pomodoro client interacts with pomo daemon.
// Author: github.com/11me/pomo
// Date: 29/07/2021
package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// Reply type to save reply from API.
type Reply string

var r Reply

// Flags
var (
	start bool

	stop bool

	set_duration time.Duration

	get_duration bool

	restart bool
)

// Initizlize flags.
func init() {

	var t time.Duration //FIXME

	flag.BoolVar(&start, "start", false, "start pomo")

	flag.BoolVar(&stop, "stop", false, "stop pomo")

	flag.DurationVar(&set_duration, "set-duration", t, "set duration")

	flag.BoolVar(&get_duration, "get-duration", false, "get duration")

	flag.BoolVar(&restart, "restart", false, "pomo reset timer")

	flag.Parse()

}

func main() {

	// Connect to pomo daemon.
	pomo, err := rpc.DialHTTP("tcp", "localhost:4444")

	if err != nil {
		log.Fatal("cannot connect to pomo daemon:", err)
	}

	// Check each given flag and make an API calls.
	flag.Visit(func(f *flag.Flag) {

		switch (*f).Name {

		case "start":
			pomo.Call("PomoAPI.Start", true, &r)

		case "stop":
			pomo.Call("PomoAPI.Stop", true, &r)

		case "set-duration":
			pomo.Call("PomoAPI.SetDuration", set_duration, &r)

		case "get-duration":
			pomo.Call("PomoAPI.GetDuration", true, &r)
			fmt.Printf("%s", r)

		case "restart":
			pomo.Call("PomoAPI.Restart", true, &r)
		}

	})

}
