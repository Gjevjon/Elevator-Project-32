package main

import (
	"Driver-go/elevio"
	"fmt"
	"time"
)

func main() {
	numFloors := 3
	// Here we initiate the simulator!
	elevio.Init("localhost:15657", numFloors)
	fmt.Printf("Started!\n")
	// The default direction. Should this be a global variable???
	var d elevio.Dirn = elevio.D_Up

	// Different channels
	// Drive buttons.
	drv_buttons := make(chan elevio.ButtonEvent)
	// Drive floors. Does this take the drive to request?
	drv_floors := make(chan int)
	// Is this the obstruction handler? What exactly does the obstruction do?
	drv_obstr := make(chan bool)
	// Stop requests?
	drv_stop := make(chan bool)

	drv_timer := make(chan bool)

	// We have 4 go routes running.

	// How to implement a timer?
	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)
	go elevio.Timer_TimedOut(drv_timer)

	//if(input.floorSensor() == -1){
	//    fsm_onInitBetweenFloors();
	//}

	for {
		select {
		case a := <-drv_buttons:
			elevio.Fsm_onRequestButtonPress(a.Floor, a.Button)
		case a := <-drv_floors:
			elevio.FsmOnFloorArrival(a)
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
		case a := <-drv_timer:
			fmt.Println(a)
			if a {
				elevio.Timer_Stop()
				elevio.Fsm_onDoorTimeout()
				fmt.Println("THIS")
			}
		}
		time.Sleep(1)

	}
}
