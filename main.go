package main

import (
	log "github.com/sirupsen/logrus"
)

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {

	// Template
	defer fileClose()
	log.Println("8. 구조체")
	log.Println("구조체 내용 변경하기")

	var s subscriber
	applyDiscount(&s)
	log.Println(s.rate)

}
