package main

import (
	"github.com/cckn/go-utils/vehicle"
	log "github.com/sirupsen/logrus"
)

type Vehicles interface {
	Accelerate()
	Break()
	Steer(diriection string)
}

func main() {
	defer fileClose()
	log.Println("Part.11")
	log.Println("인터페이스")

	var v1 Vehicles = vehicle.Car("")
	var v2 Vehicles = vehicle.Truck("")

	v1.Accelerate()
	v1.Break()

	v2.Accelerate()
	v2.Break()
	v2.Steer("left")
}
