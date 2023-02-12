/*
var timerActive bool
var timerEndTime int64

	func get_wall_time() int64 {
		return time.Now().UnixNano()
	}

	func timer_start(duration int64) {
		timerEndTime = get_wall_time() + duration
		fmt.Println(duration)
		fmt.Println(get_wall_time())
		fmt.Println(timerEndTime)
		timerActive = true
	}

	func Timer_Stop() {
		timerActive = false
	}

// Should have its own timer channel. We write to the channel only when timed out. :)

	func Timer_TimedOut(receiver chan<- bool) {
		receiver <- (timerActive && (get_wall_time() > timerEndTime))
	}
*/
package elevio

import (
	"fmt"
	"time"
)

var timerEndTime time.Time
var timerActive bool

func Timer_start(duration float64) {
	timerEndTime = time.Now().Add(3 * time.Second)
	fmt.Println("timerEndTime is:", timerEndTime)
	timerActive = true
}

func Timer_stop() {
	timerActive = false
}

/*
func Timer_timedOut() bool {
	if timerActive && time.Now().After(timerEndTime) {
		fmt.Println("timerEndTime is:", timerEndTime)
		return true
	}
	return false
}*/

func Timer_timedOut(receiver chan<- bool) {
	for {
		receiver <- (timerActive && time.Now().After(timerEndTime))
	}
}
