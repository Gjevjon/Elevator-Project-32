package main

import (
	"Driver-go/elevio"
	"fmt"
)

func order_handler(order elevio.ButtonEvent, old_cab_orders [4]bool, old_up_orders [4]bool, old_down_orders [4]bool) (cab_orders [4]bool, up_orders [4]bool, down_orders [4]bool) {

	// Just have list arrays equal to the number of floors
	// Are there any ups greater than.
	// up: 1 0 0 0
	// Are there any smaller than.
	// down: 1 0 0 0
	// cab: 1 1 1 1

	cab_orders = old_cab_orders
	up_orders = old_up_orders
	down_orders = old_down_orders

	// Missing capacity
	if order.Button == 1 {
		down_orders[order.Floor] = true
	} else if order.Button == 2 {
		up_orders[order.Floor] = true
	} else {
		cab_orders[order.Floor] = true
	}

	// Send this data to the floor handler
	return
}

func main() {
	numFloors := 4
	// Here we initiate the simulator!
	elevio.Init("localhost:15657", numFloors)
	fmt.Printf("Started!\n")
	// The default direction. Should this be a global variable???
	var d elevio.Dirn = elevio.D_Up

	// Function to control the elevator. Seems like we start driving straight up-
	elevio.SetMotorDirection(d)

	// Different channels
	// Drive buttons.
	drv_buttons := make(chan elevio.ButtonEvent)
	// Drive floors. Does this take the drive to request?
	drv_floors := make(chan int)
	// Is this the obstruction handler? What exactly does the obstruction do?
	drv_obstr := make(chan bool)
	// Stop requests?
	drv_stop := make(chan bool)

	// We have 4 go routes running.

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	for {
		select {
		case a := <-drv_buttons:
			elevio.Fsm_onRequestButtonPress(a.Floor, a.Button)
		case a := <-drv_floors:

			//fmt.Printf("%+v\n", a)
			// What to happen if we stop/ Run out of orders.
			// This is the last things we check?
			if a == numFloors-1 {
				d = elevio.D_Down
			} else if a == 0 {
				d = elevio.D_Up
			}

			elevio.SetMotorDirection(d)

			// Controller
			// If we are on a floor with a request
			// Stop motors
			// Set door open lamp true
			// Set bools of that floor to 0, only the ones in the same direction as the elevator is moving and cab.
			// Set door open lamp false
			// Check if there are more requests in same direction.
			// Check greater thans
			// The same goes for the last floor.
			// If empty stop.
			//
		case a := <-drv_obstr:
			fmt.Printf("%+v\n", a)
			if a {
				elevio.SetMotorDirection(elevio.D_Stop)
			} else {
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
