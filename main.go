package main

import (
	"LearningGo/geo"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Template
	defer fileClose()
	log.Println("8. 구조체")
	log.Println("연습문제")

	location := geo.Landmark{
		Name: "The GooglePlex",
		Coordinates: geo.Coordinates{
			Latitude:  37.42,
			Longitude: -122.08,
		},
	}

	log.Println(location)
}
