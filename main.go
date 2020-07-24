package main

import (
	log "github.com/sirupsen/logrus"
)

func printInfo(s *subscriber) {
	log.Println(s.name)
	log.Println(s.rate)
	log.Println(s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5
	s.active = true

	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
func main() {

	// Template
	defer fileClose()
	log.Println("8. 구조체")
	log.Println("구조체 변수")

	s1 := defaultSubscriber("aman")
	applyDiscount(s1)
	printInfo(s1)

	s2 := defaultSubscriber("beth")
	printInfo(s2)
}
