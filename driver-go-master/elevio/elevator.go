package elevio

import (
	"fmt"
)

type ElevatorBehaviour int

const (
	EB_Idle     ElevatorBehaviour = 0
	EB_DoorOpen ElevatorBehaviour = 1
	EB_Moving   ElevatorBehaviour = 2
)

type ClearRequestVariant int

const (
	CV_All    ClearRequestVariant = 0
	CV_InDirn ClearRequestVariant = 1
)

type Elevator struct {
	floor     int
	dirn      Dirn
	requests  [N_FLOORS][N_BUTTONS]bool
	behaviour ElevatorBehaviour
	config    Config
}

type Config struct {
	clearRequestVariant ClearRequestVariant
	doorOpenDuration_s  int
}

// Should be moved to elevio
func elevio_button_toString(b ButtonType) string {
	if b == B_HallUp {
		return "B_HallUp"
	} else if b == B_HallDown {
		return "B_HallDown"
	} else if b == B_Cab {
		return "B_Cab"
	} else {
		return "B_UNDEFINED"
	}
}

func eb_toString(eb ElevatorBehaviour) string {
	if eb == EB_Idle {
		return "EB_Idle"
	} else if eb == EB_DoorOpen {
		return "EB_DoorOpen"
	} else if eb == EB_Moving {
		return "EB_Moving"
	} else {
		return "EB_UNDEFINED"
	}
}

func elevio_dirn_toString(d Dirn) string {
	if d == D_Up {
		return "D_Up"
	} else if d == D_Down {
		return "D_Down"
	} else if d == D_Stop {
		return "D_Stop"
	} else {
		return "D_UNDEFINED"
	}
}

func elevator_print(es Elevator) {
	fmt.Printf("  +--------------------+\n")
	fmt.Printf("  |floor = %d           |\n", es.floor)
	fmt.Printf("  |dirn  = %-12.12s|\n", elevio_dirn_toString(es.dirn))
	fmt.Printf("  |behav = %-12.12s|\n", eb_toString(es.behaviour))
	fmt.Printf("  +--------------------+\n")
	fmt.Printf("  |  | up  | dn  | cab |\n")
	for f := N_FLOORS - 1; f >= 0; f-- {
		fmt.Printf("  | %d", f)
		for btn := 0; btn < N_BUTTONS; btn++ {
			if (f == N_FLOORS-1 && btn == int(B_HallUp)) || (f == 0 && btn == int(B_HallDown)) {
				fmt.Printf("|     ")
			} else {
				if es.requests[f][btn] {
					fmt.Printf("|  #  ")
				} else {
					fmt.Printf("|  -  ")
				}
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("  +--------------------+\n")
}

func Elevator_uninitialized() {
	elevator = Elevator{floor: -1, dirn: D_Stop, behaviour: EB_Idle, config: Config{clearRequestVariant: CV_All, doorOpenDuration_s: 3}}
}
