// Pomodoro daemon runs in the background
// and pomodoro client interacts with it.
// Author: github.com/11me/pomo
// Date: 29/07/2021
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/11me/pomo/pomod/lib"
)

// Declaring global pomo variable.
var pomo lib.Pomo

// Initializing pomo fields.
func init() {

	pomo.Status = lib.OFF

	pomo.Time, _ = time.ParseDuration("25m")

	pomo.Emoji = "🍅"

	pomo.Blink = [2]string{"🚶", "🏃"}

	fmt.Printf("%s pomo is initialized with default time %s\n", pomo.Emoji, pomo.Time)

}

func main() {

	// Create a new API.
	api := new(lib.PomoAPI)
	api.Pomo = &pomo

	rpc.Register(api)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4444")

	if err != nil {
		log.Fatal("error: ", err)
	}

	http.Serve(listener, nil)
}
