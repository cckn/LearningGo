package main

import (
	log "github.com/sirupsen/logrus"
)

type myNumber int

func (n myNumber) Add(addNum int) int {
	return int(n) + addNum
}

func (n myNumber) Subtract(subNum int) int {
	return int(n) - subNum
}
func main() {

	// Template
	defer fileClose()
	log.Println("8. 나만의 타입 - 사용자 정의 타입")
	log.Println("연습문제")

	var number myNumber = 100
	log.Println(number.Add(10))
	log.Println(number.Subtract(10))
}
