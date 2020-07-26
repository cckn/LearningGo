package main

import (
	"github.com/cckn/go-utils/geo"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer fileClose()
	log.Println("Part.10")
	log.Println("연습 문제")
	log.Println("landmark getter/setter")

	var lm geo.Landmark
	err := lm.SetName("LF outlet")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lm.Name())

}
