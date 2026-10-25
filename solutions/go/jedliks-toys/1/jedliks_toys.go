package jedlik

import "fmt"

// TODO: define the 'Drive()' method
//
//	type Car struct {
//		speed        int
//		batteryDrain int
//		battery      int
//		distance     int
//	}
func (car *Car) Drive() {
	postBatteryDrain := car.battery - car.batteryDrain
	if postBatteryDrain >= 0 {
		car.battery = car.battery - car.batteryDrain
		car.distance = car.distance + car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)

}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	// 	speedRequiredForCompletingTrackDistance := trackDistance / car.speed
	// 	pendinDistance := trackDistance % car.speed
	// 	batteryDrainWithSpeedPerMeter := speedRequiredForCompletingTrackDistance * car.batteryDrain
	// 	carBatteryAfterDepletion := car.battery - batteryDrainWithSpeedPerMeter
	//
	// 	return carBatteryAfterDepletion >= 0 && pendinDistance == 0

	drivesLeft := car.battery / car.batteryDrain
	maxDistance := drivesLeft * car.speed
	return maxDistance >= trackDistance

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
