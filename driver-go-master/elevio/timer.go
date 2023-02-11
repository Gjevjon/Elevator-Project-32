package elevio

import (
	"fmt"
	"time"
)

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
