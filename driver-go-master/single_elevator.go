package main

import (
	"Driver-go/elevio"
	//"container/list"
	"fmt"
)

func GoToFloor(destination int) {

}

func ReciveOrder(button elevio.ButtonType, floor int) {
	maxLength := 20
	var _orderFloor [maxLength]int
	var _orderType [maxLength]elevio.ButtonType

	_orderFloor = _orderFloor.append(floor)
	_orderType = _orderType.append(button)

	return
}

func SortOrders(unsortedFloors []int, unsortedButtons []elevio.ButtonType) {
	firstFloor := unsortedFloors[0]
	firstButton := unsortedButtons[0]

	length := len(unsortedButtons)

	var _sortedFloor [maxLength]int
	var _sortedButton [maxLength]elevio.ButtonType

	for i := 1; i < length-1; i++ {
		if unsortedButtons[i] == firstButton || unsortedButtons[i] == elevio.BT_Cab {
			if unsortedButtons[i] == elevio.BT_HallDown && unsortedFloors[i] < firstFloor {
				// add before first element in sorted list and update new first order
				tempFloor := firstFloor
				tempButton := firstButton
				_sortedFloor[0] = unsortedFloors[i]
				_sortedButton[0] = unsortedButtons[i]
				_sortedFloor[i] = tempFloor
				_sortedButton[i] = tempButton
			} else if unsortedButtons[i] == elevio.BT_HallUp && unsortedFloors[i] > firstFloor {
				// add before first element in sorted list update new first order
			} else {
				// add to end of sorted list
			}
		}
	}
}

func main() {
	numFloors := 4

	elevio.Init("localhost:15657", numFloors)

	var d elevio.MotorDirection = elevio.MD_Up
	//elevio.SetMotorDirection(d)
	var _currentFloor int = 0
	var _destination int = 0

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	for {
		select {
		case a := <-drv_buttons:
			fmt.Printf("%+v\n", a)

			ReciveOrder(a.Button, a.Floor)
			SortOrders() // Need to add the arrays from teh function above

			elevio.SetButtonLamp(a.Button, a.Floor, true)

		case a := <-drv_floors:
			fmt.Printf("%+v\n", a)
			if a == numFloors-1 { // if the elevator is at the top, turn to go down
				d = elevio.MD_Down
			} else if a == 0 { // if the elevator is at the ground floor go up
				d = elevio.MD_Up
			}
			elevio.SetMotorDirection(d)

		case a := <-drv_obstr:
			fmt.Printf("%+v\n", a)
			if a { // if there is an obstruction stop
				elevio.SetMotorDirection(elevio.MD_Stop)
			} else { // else continue
				elevio.SetMotorDirection(d)
			}

		case a := <-drv_stop:
			fmt.Printf("%+v\n", a)
			for f := 0; f < numFloors; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}
		}
	}
}
