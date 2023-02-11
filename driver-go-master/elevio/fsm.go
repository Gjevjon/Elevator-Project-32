package elevio

import (
	"fmt"
)

var elevator Elevator

func Fsm_onRequestButtonPress(btn_floor int, btn_type Button) {
	fmt.Printf("\n\n%s(%d, %s)\n", "fsm_onRequestButtonPress", btn_floor, elevio_button_toString(btn_type))
	elevator_print(elevator)

	switch elevator.behaviour {
	case EB_DoorOpen:
		if requests_shouldClearImmediately(elevator, btn_floor, btn_type) {
			timer_start(elevator.config.doorOpenDuration_s)
		} else {
			elevator.requests[btn_floor][btn_type] = true
		}
	case EB_Moving:
		elevator.requests[btn_floor][btn_type] = true
	case EB_Idle:
		elevator.requests[btn_floor][btn_type] = true
		var pair DirnBehaviourPair = requests_chooseDirection(elevator)
		elevator.dirn = pair.dirn
		elevator.behaviour = pair.behaviour
		switch pair.behaviour {
		case EB_DoorOpen:
			SetDoorOpenLamp(true)
			timer_start(elevator.config.doorOpenDuration_s)
			elevator = requests_clearAtCurrentFloor(elevator)
		case EB_Moving:
			SetMotorDirection(elevator.dirn)
		case EB_Idle:
		}
	}
}
