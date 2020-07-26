package main

import (
	"github.com/cckn/go-utils/geo"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer fileClose()
	log.Println("Part.10")
	log.Println("캡슐화")
	log.Println("연습 문제")

	var place geo.Coordinate

	err := place.SetLatitude(30)
	if err != nil {
		log.Fatal(err)
	}
	err = place.SetLongitude(-170)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(place.Latitude(), place.Longitude())

}
