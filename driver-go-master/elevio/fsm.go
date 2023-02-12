package elevio

import (
	"fmt"
)

var elevator Elevator

func setAllLights(es Elevator) {
	for floor := 0; floor < N_FLOORS; floor++ {
		for btn := 0; btn < N_BUTTONS; btn++ {
			// Unsure wther this works
			SetButtonLamp(ButtonType(btn), floor, es.requests[floor][btn])
		}
	}
}

func Fsm_onInitBetweenFloors() {
	SetMotorDirection(D_Down)
	elevator.dirn = D_Down
	elevator.behaviour = EB_Moving
}

func Fsm_onRequestButtonPress(btn_floor int, btn_type ButtonType) {
	fmt.Printf("\n\n%s(%d, %s)\n", "fsm_onRequestButtonPress", btn_floor, elevio_button_toString(btn_type))
	elevator_print(elevator)

	switch elevator.behaviour {
	case EB_DoorOpen:
		if requests_shouldClearImmediately(elevator, btn_floor, btn_type) {
			Timer_start(elevator.config.doorOpenDuration_s)
		} else {
			elevator.requests[btn_floor][btn_type] = true
		}
	case EB_Moving:
		elevator.requests[btn_floor][btn_type] = true
	case EB_Idle:
		elevator.requests[btn_floor][btn_type] = true
		pair := requests_chooseDirection(elevator)
		elevator.dirn = pair.dirn
		elevator.behaviour = pair.behaviour
		switch pair.behaviour {
		case EB_DoorOpen:
			SetDoorOpenLamp(true)
			Timer_start(elevator.config.doorOpenDuration_s)
			elevator = requests_clearAtCurrentFloor(elevator)
		case EB_Moving:
			SetMotorDirection(elevator.dirn)
		case EB_Idle:
		}
	}
	setAllLights(elevator)
	fmt.Printf("\nNew state:\n")
	elevator_print(elevator)

}

func Fsm_OnFloorArrival(newFloor int) {
	fmt.Println("\n\n", "fsmOnFloorArrival", "(", newFloor, ")")
	elevator_print(elevator)

	elevator.floor = newFloor

	SetFloorIndicator(elevator.floor)
	switch elevator.behaviour {
	case EB_Moving:
		if requests_shouldStop(elevator) {
			fmt.Printf("Stop on the way up")
			SetMotorDirection(D_Stop)
			SetDoorOpenLamp(true)
			elevator = requests_clearAtCurrentFloor(elevator)
			Timer_start(elevator.config.doorOpenDuration_s)
			fmt.Println("Floor arrival time input is", elevator.config.doorOpenDuration_s)
			setAllLights(elevator)
			elevator.behaviour = EB_DoorOpen
		}
	default:
	}

	fmt.Println("\nNew state:")
	elevator_print(elevator)
}

func Fsm_onDoorTimeout() {
	fmt.Printf("\n\n fsm_onDoorTimeout()\n")
	elevator_print(elevator)

	switch elevator.behaviour {
	case EB_DoorOpen:
		pair := requests_chooseDirection(elevator)
		elevator.dirn = pair.dirn
		elevator.behaviour = pair.behaviour

		switch elevator.behaviour {
		case EB_DoorOpen:
			Timer_start(elevator.config.doorOpenDuration_s)
			fmt.Println("Floor door time out input is", elevator.config.doorOpenDuration_s)
			elevator = requests_clearAtCurrentFloor(elevator)
			setAllLights(elevator)
		case EB_Moving, EB_Idle:
			SetDoorOpenLamp(false)
			SetMotorDirection(elevator.dirn)
		}
	default:
	}

	fmt.Printf("\nNew state:\n")
	elevator_print(elevator)
}
