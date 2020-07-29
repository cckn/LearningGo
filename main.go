package main

import (
	log "github.com/sirupsen/logrus"
)

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	defer fileClose()
	log.Println("Part.11")
	log.Println("Stringer interface")

	coffeePot := CoffeePot("LuxBrew")
	log.Println(coffeePot.String())

}
