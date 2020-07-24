package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	// Template
	defer fileClose()
	log.Println("8. 구조체")
	log.Println("사용자 정의 타입")

	type part struct {
		description string
		count       int
	}

	type car struct {
		name     string
		topSpeed float64
	}

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	log.Println(porsche.name)
	log.Println(porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	log.Println(bolts.description)
	log.Println(bolts.count)

}
