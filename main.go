package main

import (
	log "github.com/sirupsen/logrus"
)

type NoiseMaker interface {
	MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	log.Println("BBeP BBeP")
}
func (r Robot) Walk() {
	log.Println("Powering legs")
}

func main() {
	defer fileClose()
	log.Println("Part.11")
	log.Println("타입단언")

	var noiseMaker NoiseMaker = Robot("")
	noiseMaker.MakeSound()

	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}
