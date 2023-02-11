package elevio

import "time"

var timerActive bool
var timerEndTime int64

func get_wall_time() int64 {
	return time.Now().UnixNano()
}

func timer_start(duration int64) {
	timerEndTime = get_wall_time() + duration
	timerActive = true
}

func timer_stop() {
	timerActive = false
}

func timer_timedOut() bool {
	var timedOut bool = (timerActive && get_wall_time() > timerEndTime)
	return timedOut
}
