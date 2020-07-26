package main

import (
	"github.com/cckn/go-utils/date"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer fileClose()
	log.Println("Part.10")
	log.Println("캡슐화")
	log.Println("따로 빼둬서 내부 필드 접근 불가하도록")

	day := date.Date{}

	err := day.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = day.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	err = day.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(day)

}
