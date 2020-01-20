package main

import "fmt"

type car struct {
	gasPedal       uint16
	brakePedal     uint16
	steeringWheerl int16
	topSpeedKmh    float64
}

func main() {
	// both assignments are doing same for structs
	// aCar := car{1, 0, 12561, 225.0}
	aCar := car{gasPedal: 1, brakePedal: 0, steeringWheerl: 12561, topSpeedKmh: 225.0}
	fmt.Println(aCar)
}
